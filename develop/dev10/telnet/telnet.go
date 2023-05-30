package telnet

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"golang.org/x/sync/errgroup"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func CLI(args []string) int {
	var app Application
	err := app.fromArgs(args)
	if err != nil {
		return 2
	}
	if err = app.Run(); err != nil {
		log.Fatalln(err)
		return 1
	}
	return 0
}

type Application struct {
	timeout time.Duration
	addr    string
}

func (app *Application) fromArgs(args []string) error {
	flagSet := flag.NewFlagSet("telnet", flag.ContinueOnError)

	flagSet.DurationVar(&app.timeout, "timeout", time.Second*10, "timeout to connect to the server")

	if err := flagSet.Parse(args); err != nil {
		flagSet.Usage()
		return err
	}

	app.addr = net.JoinHostPort(flagSet.Arg(0), flagSet.Arg(1))

	return nil
}
func (app *Application) Run() error {
	dialer := net.Dialer{
		Timeout: app.timeout,
	}

	c, err := dialer.Dial("tcp", app.addr)
	if err != nil {
		return err
	}
	defer c.Close()

	ctx, cancel := context.WithCancel(context.Background())

	eGroup := new(errgroup.Group)

	eGroup.Go(func() error {
		reader := bufio.NewReader(os.Stdin)
		for {
			select {
			case <-ctx.Done():
				fmt.Println("first group stop ch")
				return nil

			default:
				fmt.Print("$: ")
				text, err := reader.ReadString('\n')
				if err != nil {
					return err
				}

				_, err = fmt.Fprint(c, text)
				if err != nil {
					return err
				}
			}
		}
	})
	eGroup.Go(func() error {
		reader := bufio.NewReader(os.Stdin)
		for {
			select {
			case <-ctx.Done():
				fmt.Println("second group stop ch")
				return nil

			default:
				text, err := reader.ReadString('\n')
				if err != nil {
					return err
				}
				fmt.Printf("got from server: %s\n", text)
			}
		}
	})

	go func() {
		stop := make(chan os.Signal, 1)
		signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
		<-stop
		fmt.Printf("\ninterrupt signal\n")
		cancel()
	}()

	err = eGroup.Wait()
	if err != io.EOF {
		return err
	}
	fmt.Printf("server close\nexiting...")

	return nil
}
