package pkg

type VendingMachine struct {
	hasItem       State
	itemRequested State
	hasMoney      State
	noItem        State
	currentState  State
	itemCount     int
	itemPrice     int
}

func (v *VendingMachine) AddItem() error {
	return v.currentState.AddItem()
}

func (v *VendingMachine) RequestItem() error {
	return v.currentState.RequestItem()
}

func (v *VendingMachine) InsertMoney(money int) error {
	return v.currentState.InsertMoney(money)
}

func (v *VendingMachine) DispenseItem() error {
	return v.currentState.DispenseItem()
}

func (v *VendingMachine) setState(s State) {
	v.currentState = s
}

func NewVendingMachine(itemCount, itemPrice int) *VendingMachine {
	v := &VendingMachine{itemCount: itemCount, itemPrice: itemPrice}
	hasItemState := &HasItemState{vendingMachine: v}
	hasMoneyState := &HasMoneyState{vendingMachine: v}
	itemRequestedState := &ItemRequestedState{vendingMachine: v}
	noItemState := &NoItemState{vendingMachine: v}

	v.setState(hasItemState)
	v.hasItem = hasItemState
	v.itemRequested = itemRequestedState
	v.hasMoney = hasMoneyState
	v.noItem = noItemState
	return v
}
