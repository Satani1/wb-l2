package main

import (
	"os"
	"task3/strings-sort"
)

/*
Example of run:
go run main.go -n true "./test/nums.txt"
go run main.go  -r true -u=true "./test/text.txt"
*/
func main() {
	os.Exit(strings_sort.CLI(os.Args[1:]))
}
