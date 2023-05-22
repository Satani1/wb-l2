package main

import (
	"os"
	"task3/strings-sort"
)

func main() {
	os.Exit(strings_sort.CLI(os.Args[1:]))
}
