package pkg

import "fmt"

type HasItemState struct {
	vendingMachine *VendingMachine
}

func (h *HasItemState) AddItem() error {
	fmt.Println("Item added")
	h.vendingMachine.itemCount++
	return nil
}

func (h *HasItemState) RequestItem() error {
	if h.vendingMachine.itemCount == 0 {
		h.vendingMachine.setState(h.vendingMachine.noItem)
		fmt.Errorf("no item")
	}
	fmt.Println("Item requested")
	h.vendingMachine.setState(h.vendingMachine.itemRequested)
	return nil
}

func (h *HasItemState) InsertMoney(money int) error {
	return fmt.Errorf("Please select item first")
}

func (h *HasItemState) DispenseItem() error {
	return fmt.Errorf("Please select item first")
}
