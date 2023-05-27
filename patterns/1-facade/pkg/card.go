package pkg

import (
	"fmt"
	"time"
)

type Card struct {
	Name    string
	Balance float64
	Bank    *Bank
}

func (c *Card) CheckBalance() error {
	fmt.Println("[Карта] Запрос в банк для проверки остатка")
	time.Sleep(time.Millisecond * 800)

	return c.Bank.CheckBalance(c.Name)
}
