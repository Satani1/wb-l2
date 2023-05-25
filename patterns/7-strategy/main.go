package main

import "strategy/pkg"

var (
	start      = 10
	end        = 100
	strategies = []pkg.Strategy{
		&pkg.PublicTransportStrategy{},
		&pkg.RoadStrategy{},
		&pkg.WalkStrategy{},
	}
)

func main() {
	navigator := pkg.Navigator{}
	for _, strategy := range strategies {
		navigator.SetStrategy(strategy)
		navigator.Route(start, end)
	}
}
