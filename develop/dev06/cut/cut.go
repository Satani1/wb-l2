package cut

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type Application struct {
	fields    []int
	delimiter string
	separated bool
	reader    io.ReadCloser
	end       bool
}

func CLI(args []string) int {
	var app Application

	err := app.fromArgs(args)
	if err != nil {
		log.Fatalln(err)
		return 2
	}

	if err := app.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Runtime error: %v\n", err)
		return 1
	}
	return 0
}

func (app *Application) fromArgs(args []string) error {
	flagSet := flag.NewFlagSet("cut", flag.ContinueOnError)

	flagSet.Func("f", "select fields", app.Fields)
	flagSet.StringVar(&app.delimiter, "d", "\t", "select separator")
	flagSet.BoolVar(&app.separated, "s", false, "do/doesnt print lines without separator")

	if err := flagSet.Parse(args); err != nil {
		return err
	}
	//check file
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		app.reader = os.Stdin
		return nil
	}

	//open file
	file, err := os.Open(flagSet.Arg(0))
	if err != nil {
		return err
	}

	app.reader = file
	fmt.Println(app)
	return nil
}

func (app *Application) Run() error {
	defer app.reader.Close()

	scanner := bufio.NewScanner(app.reader)

	for scanner.Scan() {
		result := strings.Split(scanner.Text(), app.delimiter)
		l := len(result)

		if l == 1 && app.separated {
			continue
		}
		if l == 1 && !app.separated {
			fmt.Println(result[0])
			continue
		}

		//fields
		if app.end {
			for i := app.fields[0]; i < l; i++ {
				fmt.Printf("%s%s", result[i], app.delimiter)
			}
			fmt.Println()
			continue
		}

		for _, v := range app.fields {
			if v < l {
				fmt.Printf("%s%s", result[v], app.delimiter)
			}
		}
		fmt.Println()
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

func (app *Application) Fields(s string) error {
	if strings.Contains(s, ",") {
		strNums := strings.Split(s, ",")
		for _, v := range strNums {
			field, err := strconv.Atoi(v)
			if err != nil {
				return err
			}
			app.fields = append(app.fields, field-1)
		}
		return nil
	}

	if strings.Contains(s, "-") {
		var from, to int
		var err error
		if len(s) == 3 {
			from, err = strconv.Atoi(string(s[0]))
			if err != nil {
				return err
			}
			to, err = strconv.Atoi(string(s[2]))
			if err != nil {
				return err
			}
		}
		if len(s) == 2 {
			if s[0] == '-' {
				from = 1
				to, err = strconv.Atoi(string(s[1]))
				if err != nil {
					return err
				}
			}
			if s[1] == '-' {
				app.end = true
				from, err = strconv.Atoi(string(s[0]))
				if err != nil {
					return err
				}
				to = from

			}
		}
		for ; from <= to; from++ {
			app.fields = append(app.fields, from-1)
		}

		return nil
	}

	n, err := strconv.Atoi(s)
	if err != nil {
		return err
	}
	app.fields = append(app.fields, n-1)

	return nil
}
