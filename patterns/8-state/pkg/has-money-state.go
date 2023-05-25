package pkg

import "fmt"

type HasMoneyState struct {
	vendingMachine *VendingMachine
}

func (h *HasMoneyState) AddItem() error {
	return fmt.Errorf("Item dispense in progress")
}

func (h *HasMoneyState) RequestItem() error {
	return fmt.Errorf("Item dispense in progress")
}

func (h *HasMoneyState) InsertMoney(money int) error {
	return fmt.Errorf("Item out of stock")
}

func (h *HasMoneyState) DispenseItem() error {
	fmt.Println("Dispensing item")
	h.vendingMachine.itemCount--
	if h.vendingMachine.itemCount == 0 {
		h.vendingMachine.setState(h.vendingMachine.noItem)
	} else {
		h.vendingMachine.setState(h.vendingMachine.hasItem)
	}
	return nil
}
