package pkg

import "fmt"

type NoItemState struct {
	vendingMachine *VendingMachine
}

func (n *NoItemState) AddItem() error {
	n.vendingMachine.itemCount++
	n.vendingMachine.setState(n.vendingMachine.hasItem)
	return nil
}

func (n *NoItemState) RequestItem() error {
	return fmt.Errorf("Item out of stock")
}

func (n *NoItemState) InsertMoney(money int) error {
	return fmt.Errorf("Item out of stock")
}

func (n *NoItemState) DispenseItem() error {
	return fmt.Errorf("Item out of stock")
}
