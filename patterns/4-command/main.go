package main

import (
	"command/pkg"
	"fmt"
)

func main() {
	invoker := &pkg.Invoker{}
	receiver := &pkg.Receiver{}

	invoker.StoreCommand(&pkg.ToggleOnCommand{
		Receiver: receiver,
	})
	invoker.StoreCommand(&pkg.ToggleOffCommand{
		Receiver: receiver,
	})

	result := invoker.Execute()
	fmt.Println(result)
}
