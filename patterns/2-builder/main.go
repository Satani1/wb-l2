package main

import "builder/pkg"

func main() {
	asusCollector := pkg.GetCollector("asus")
	hpCollector := pkg.GetCollector("hp")

	//create 'asus' base configuration
	factory := pkg.NewFactory(asusCollector)
	asusComputer := factory.CreateComputer()
	asusComputer.Print()

	//switch base configuration from 'asus' to 'hp'
	factory.SetCollector(hpCollector)
	hpComputer := factory.CreateComputer()
	hpComputer.Print()

	//back switch to 'asus'
	factory.SetCollector(asusCollector)
	pc := factory.CreateComputer()
	pc.Print()
}
