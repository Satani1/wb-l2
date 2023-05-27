package pkg

type User struct {
	Name string
	Card *Card
}

func (u *User) GetBalance() float64 {
	return u.Card.Balance
}
