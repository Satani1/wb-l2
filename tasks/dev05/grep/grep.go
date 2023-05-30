package grep

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"
)

type Application struct {
	reader     io.ReadCloser
	after      int
	before     int
	context    int
	count      int
	ignoreCase bool
	invert     bool
	fixed      bool
	lineNum    bool
	input      []string
	pattern    string
}

func CLI(args []string) int {
	var app Application

	err := app.fromArgs(args)
	if err != nil {
		fmt.Println(err)
		return 2
	}

	if err := app.Run(); err != nil {
		return 1
	}

	return 0
}

func (app *Application) fromArgs(args []string) error {
	//parsing keys from terminal
	flagSet := flag.NewFlagSet("gogrep", flag.ContinueOnError)

	flagSet.IntVar(&app.after, "A", 0, "print +N lines after matching lines")
	flagSet.IntVar(&app.before, "B", 0, "print +N lines before matching lines")
	flagSet.IntVar(&app.context, "C", 0, "print +-N (A + B) lines around matching lines")
	flagSet.IntVar(&app.count, "c", -1, "amount of lines")
	flagSet.BoolVar(&app.ignoreCase, "i", false, "ignore registry")
	flagSet.BoolVar(&app.invert, "v", false, "invert result; print non matching lines")
	flagSet.BoolVar(&app.fixed, "F", false, "exactly matching line")
	flagSet.BoolVar(&app.lineNum, "n", true, "print numbers of lines")

	if err := flagSet.Parse(args); err != nil {
		flagSet.Usage()
		return err
	}
	fmt.Println(app)
	if app.after == 0 {
		app.after = app.context
	}
	if app.before == 0 {
		app.before = app.context
	}

	fmt.Println(flagSet.Args())
	app.pattern = flagSet.Arg(0)

	if app.fixed {
		fmt.Sprintf("^%s$", app.pattern)
	}
	if app.ignoreCase {
		app.pattern = "(?i)" + app.pattern
	}

	//check info about file
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		app.reader = os.Stdin
		return nil
	}
	//open file
	file, err := os.Open(flagSet.Arg(1))
	if err != nil {
		return err
	}

	app.reader = file
	fmt.Println(app)
	return nil
}

func (app *Application) Run() error {
	defer app.reader.Close()

	//scan file and append lines to app.input[]
	scanner := bufio.NewScanner(app.reader)

	for scanner.Scan() {
		app.input = append(app.input, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	//create regexp type
	r, err := regexp.Compile(app.pattern)
	if err != nil {
		return err
	}

	//print result
	app.PrintResult(app.FindMatching(r))

	return nil
}

func (app *Application) FindMatching(r *regexp.Regexp) []int {
	matchLines := make([]int, 0)

	for index, line := range app.input {
		if app.count == 0 {
			break
		}

		if (r.MatchString(line) && !app.invert) || (!r.MatchString(line) && app.invert) {
			matchLines = append(matchLines, index)
			app.count--
		}
	}

	return matchLines
}

func (app *Application) PrintResult(data []int) {
	printedLines := make(map[int]struct{})
	matchedLines := make(map[int]struct{})

	for _, numberLine := range data {
		matchedLines[numberLine] = struct{}{}
	}

	var lastPrintedLine int
	for _, numberLine := range data {

		if app.before > 0 || app.after > 0 {
			if numberLine-lastPrintedLine > 2 {
				fmt.Println("----")
			}
			if _, found := printedLines[numberLine]; found {
				continue
			}

			start := numberLine - app.before
			if numberLine-app.before < 0 {
				start = 0
			}
			finish := numberLine + app.after
			if numberLine+app.after > len(app.input)-1 {
				finish = len(app.input) - 1
			}

			for ; start <= finish; start++ {
				if _, found := printedLines[start]; found {
					continue
				}
				if _, found := matchedLines[start]; found && start != 0 {
					break
				}

				if app.lineNum {
					if _, found := matchedLines[start]; found {
						fmt.Printf("%v:%v\n", start+1, app.input[start])
					} else {
						fmt.Printf("%v-%v\n", start+1, app.input[start])
					}
					printedLines[start] = struct{}{}
					lastPrintedLine = start
					continue
				}
				fmt.Println(app.input[start])
				printedLines[start] = struct{}{}
				lastPrintedLine = start
			}
		}
		if app.lineNum {
			fmt.Printf("%v:%v\n", numberLine+1, app.input[numberLine])
			continue
		}

		fmt.Println(app.input[numberLine])
	}

}
