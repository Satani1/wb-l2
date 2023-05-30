package main

import (
	"facade/pkg"
	"fmt"
)

var (
	bank = pkg.Bank{
		Name:  "SberBank",
		Cards: []pkg.Card{},
	}
	card1 = pkg.Card{
		Name:    "card-dev01",
		Balance: 200,
		Bank:    &bank,
	}
	card2 = pkg.Card{
		Name:    "card-dev02",
		Balance: 10,
		Bank:    &bank,
	}

	user1 = pkg.User{
		Name: "user-dev01",
		Card: &card1,
	}
	user2 = pkg.User{
		Name: "user-dev02",
		Card: &card2,
	}
	product1 = pkg.Product{
		Name:  "Milk",
		Price: 150,
	}
	shop = pkg.Shop{
		Name: "shop-dev01",
		Products: []pkg.Product{
			product1,
		},
	}
)

func main() {
	fmt.Println("[Банк] Выпуск карт")
	bank.Cards = append(bank.Cards, card1, card2)

	fmt.Printf("[%s]\n", user1.Name)
	err := shop.Sell(user1, product1.Name)
	if err != nil {
		println(err.Error())
		return
	}

	fmt.Printf("[%s]\n", user2.Name)
	err = shop.Sell(user2, product1.Name)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
