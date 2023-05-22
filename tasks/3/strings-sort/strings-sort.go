package strings_sort

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strings"
)

type Application struct {
	column    int
	numeric   bool
	reverse   bool
	duplicate bool
	reader    io.ReadCloser
}

func CLI(args []string) int {
	var app Application
	err := app.fromArgs(args)
	if err != nil {
		return 2
	}
	if err = app.run(); err != nil {
		fmt.Fprintf(os.Stderr, "Runtime error: %v\n", err)
		return 1
	}
	return 0
}

func (app *Application) fromArgs(args []string) error {
	flagSet := flag.NewFlagSet("sort-file", flag.ContinueOnError)
	//read keys from terminal
	flagSet.IntVar(&app.column, "k", 1, "sort based on column")
	flagSet.BoolVar(&app.numeric, "n", false, "sort numbers")
	flagSet.BoolVar(&app.reverse, "r", false, "sort in reverse")
	flagSet.BoolVar(&app.duplicate, "u", false, "delete duplicates")

	if err := flagSet.Parse(args); err != nil {
		flagSet.Usage()
		return err
	}
	log.Println(flagSet.Args())

	//check file info
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		app.reader = os.Stdin
		return nil
	}
	//open a file
	file, err := os.Open(flagSet.Arg(0))
	if err != nil {
		return err
	}

	app.reader = file
	return nil

}

func (app *Application) run() error {
	defer app.reader.Close()
	//text from file
	data := make([]string, 0)

	//scan file
	scanner := bufio.NewScanner(app.reader)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	//sort
	if app.column == 1 && !app.numeric {
		data = app.sort(data)
		//stdout sort data
		OutTerminal(data)
		return nil
	}

	data = app.columnSort(data)
	//stdout sort data
	OutTerminal(data)
	return nil
}

func (app *Application) sort(data []string) []string {
	if app.reverse {
		sort.Sort(sort.Reverse(sort.StringSlice(data)))
	} else {
		sort.Strings(data)
	}

	if app.duplicate {
		data = DeleteDuplicate(data)
	}

	return data
}

func (app *Application) columnSort(data []string) []string {
	ts := tableStrings{
		data:    make([][]string, 0, len(data)),
		column:  app.column - 1,
		numeric: app.numeric,
	}

	//fill table with strings
	for _, value := range data {
		ts.data = append(ts.data, strings.Fields(value))
	}

	//sorting
	if app.reverse {
		sort.Sort(sort.Reverse(ts))
	} else {
		sort.Sort(ts)
	}

	//back sort values
	for index, value := range ts.data {
		data[index] = strings.Join(value, " ")
	}
	//delete duplicate data
	if app.duplicate {
		data = DeleteDuplicate(data)
	}
	return data
}
