package main

import "factory-method/pkg"

var types = []string{pkg.ServerType, pkg.PersonalComputerType, pkg.NotebookType, "monoblock", "mobilephone"}

func main() {
	for _, typeName := range types {
		computer := pkg.New(typeName)
		if computer == nil {
			continue
		}
		computer.PrintDetails()
	}
}
