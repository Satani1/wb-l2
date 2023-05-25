package pkg

type State interface {
	AddItem() error
	RequestItem() error
	InsertMoney(money int) error
	DispenseItem() error
}
