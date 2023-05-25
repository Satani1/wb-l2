package pkg

type Navigator struct {
	Strategy
}

func (n *Navigator) SetStrategy(strategy Strategy) {
	n.Strategy = strategy
}
