package pkg

// ToggleOffCommand implements the Command interface
type ToggleOffCommand struct {
	Receiver *Receiver
}

func (c *ToggleOffCommand) Execute() string {
	return c.Receiver.ToggleOff()
}
