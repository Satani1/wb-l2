package main

import (
	"os"
	"task3/strings-sort"
)

/*
Example of run:
go run main.go "./test/nums.txt" -n true
go run main.go "./test/text.txt" -r true -u true
*/
func main() {
	os.Exit(strings_sort.CLI(os.Args[1:]))
}
