package shell

import (
	"bufio"
	"fmt"
	"github.com/mitchellh/go-ps"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func CLI() int {
	var app Application

	app.Run()
	return 0
}

type Application struct {
	out io.Writer
}

func (app *Application) Run() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		command := scanner.Text()

		if strings.Contains(command, "|") {
			if err := app.ExecCommand(command); err != nil {
				fmt.Fprintf(os.Stderr, "%v\n", err)
			}
		} else {
			app.out = os.Stdout
			if err := app.ExecCommand(command); err != nil {
				fmt.Fprintf(os.Stderr, "%v\n", err)
			}
		}

		path, _ := filepath.Abs(".")
		fmt.Printf("%s\n$:", path)
		scanner.Scan()
	}
}

func (app *Application) ExecCommand(command string) error {
	c := strings.Split(command, " ")

	switch c[0] {
	case "cd":
		if len(c) < 2 {
			dir, err := os.UserHomeDir()
			if err != nil {
				return err
			}
			return os.Chdir(dir)
		}
		return os.Chdir(c[1])
	case "pwd":
		pwd, err := os.Getwd()
		if err != nil {
			return err
		}
		fmt.Fprintln(app.out, pwd)
		return nil
	case "echo":
		for i := 1; i < len(c); i++ {
			fmt.Fprint(app.out, c[i], " ")
		}
		fmt.Fprintln(app.out)
		return nil
	case "kill":
		if len(c) < 2 {
			return fmt.Errorf("[kill] not enough arguments")
		}
		pid, err := strconv.Atoi(c[1])
		if err != nil {
			return err
		}

		p, err := os.FindProcess(pid)
		if err != nil {
			return err
		}
		return p.Kill()
	case "ps":
		p, err := ps.Processes()
		if err != nil {
			return err
		}

		for _, proc := range p {
			fmt.Fprintf(app.out, "%d\t%s\n", proc.Pid(), proc.Executable())
		}
		return nil
	case "quit":
		fmt.Fprint(app.out, "[shell] exiting\n")
		os.Exit(0)
	default:
		return fmt.Errorf("[shell] command not found: %s\n", c[0])
	}
	return nil
}
