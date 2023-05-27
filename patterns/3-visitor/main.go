package main

import (
	"visitor/pkg"
)

func main() {
	square := &pkg.Square{Side: 2}
	circle := &pkg.Circle{Radius: 3}
	rectangle := &pkg.Rectangle{}

	areaCalc := &pkg.AreaCalculator{}

	square.Accept(areaCalc)
	circle.Accept(areaCalc)
	rectangle.Accept(areaCalc)

}