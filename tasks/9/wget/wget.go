package wget

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
)

type Application struct {
	link           *url.URL
	outputFile     string
	depth          int
	recursive      bool
	pageRequisites bool
}

func CLI(args []string) int {
	var app Application

	err := app.fromArgs(args)
	if err != nil {
		return 2
	}
	if err := app.run(); err != nil {
		fmt.Fprintf(os.Stderr, "Runtime error: %v\n", err)
		return 1
	}
	return 0
}

func (app *Application) fromArgs(args []string) error {
	flagSet := flag.NewFlagSet("wget", flag.ContinueOnError)

	flagSet.StringVar(&app.outputFile, "o", "", "Path to output file;")
	flagSet.IntVar(&app.depth, "l", -1, "Max number of recursive download. Default is not set;")
	flagSet.BoolVar(&app.recursive, "r", false, "Turn on/off recursive download;")
	flagSet.BoolVar(&app.pageRequisites, "p", false, "Download all files that are necessary to display a HTML page;")

	if err := flagSet.Parse(args); err != nil {
		return err
	}

	u, err := url.Parse(flag.Arg(0))
	if err != nil {
		return err
	}
	app.link = u

	if app.outputFile == "" {
		app.outputFile = path.Base(app.link.Path)
	}
	return nil
}

func (app *Application) run() error {
	if app.recursive {
		queue := []string{app.link.String()}
		if err := os.Mkdir(app.link.Host, os.ModePerm); err != nil {
			return err
		}

		sm := NewSite(app.link.String(), app.link.Host)
		err := sm.DownloadSite(queue, app.depth)
		if err != nil {
			return err
		}
		return nil
	}
	return download(app.link.String(), app.outputFile)
}

func download(url, filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer file.Close()

	client := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}

	resp, err := client.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	size, err := io.Copy(io.Writer(file), resp.Body)
	if err != nil {
		return err
	}
	fmt.Printf("Downloaded a file [%s] with size [%s] bytes\n", filePath, size)
	return nil
}
