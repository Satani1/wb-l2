package pkg

// Receiver implementation
type Receiver struct{}

// ToggleOff implementation
func (r *Receiver) ToggleOff() string {
	return "Toggle Off"
}

// ToggleOn implementation
func (r *Receiver) ToggleOn() string {
	return "Toggle On"
}
