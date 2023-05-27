package pkg

// ToggleOnCommand implements the Command interface
type ToggleOnCommand struct {
	Receiver *Receiver
}

func (c *ToggleOnCommand) Execute() string {
	return c.Receiver.ToggleOn()
}
