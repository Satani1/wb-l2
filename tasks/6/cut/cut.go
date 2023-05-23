package cut

import (
	"flag"
	"io"
	"os"
	"strconv"
)

type Application struct {
	fields    []int
	delimiter string
	separated bool
	reader    io.ReadCloser
}

func CLI(args []string) int {
	var app Application

	err := app.fromArgs(args)
	if err != nil {
		return 2
	}

	if err := app.Run(); err != nil {
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

	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		app.reader = os.Stdin
		return nil
	}
	return nil
}

func (app *Application) Run() error {
	return nil
}

func (app *Application) Fields(s string) error {
	n, err := strconv.Atoi(s)
	if err != nil {
		return err
	}
	app.fields = append(app.fields, n-1)

	return nil
}