package pkg

import (
	"errors"
	"fmt"
	"time"
)

type Bank struct {
	Name  string
	Cards []Card
}

func (b *Bank) CheckBalance(cardNumber string) error {
	fmt.Println(fmt.Sprintf("[Банк] Получения остатка по карте [%s]", cardNumber))
	time.Sleep(time.Millisecond * 300)
	for _, card := range b.Cards {
		if card.Name != cardNumber {
			continue
		}
		if card.Balance <= 0 {
			return errors.New("[Банк] Недостаточно средств")
		}
	}
	fmt.Println("[Банк] Остаток положителен")
	return nil
}
