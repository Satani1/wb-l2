package main

import (
	"os"
	"task5/grep"
)

func main() {
	os.Exit(grep.CLI(os.Args[1:]))
}
