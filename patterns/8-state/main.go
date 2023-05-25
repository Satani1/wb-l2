package main

import (
	"fmt"
	"log"
	"state/pkg"
)

func main() {
	vendingMachine := pkg.NewVendingMachine(1, 10)

	err := vendingMachine.RequestItem()
	if err != nil {
		log.Fatalln(err.Error())
	}

	err = vendingMachine.InsertMoney(10)
	if err != nil {
		log.Fatalln(err.Error())
	}
	err = vendingMachine.DispenseItem()
	if err != nil {
		log.Fatalln(err.Error())
	}

	fmt.Println()

	err = vendingMachine.AddItem()
	if err != nil {
		log.Fatalln(err.Error())
	}

	fmt.Println()

	err = vendingMachine.RequestItem()
	if err != nil {
		log.Fatalln(err.Error())
	}

	err = vendingMachine.InsertMoney(10)
	if err != nil {
		log.Fatalln(err.Error())
	}

	err = vendingMachine.DispenseItem()
	if err != nil {
		log.Fatalln(err.Error())
	}

}
