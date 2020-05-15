package cpu

// Set the interrupt master enable.
func (c *CPU) setIME(v bool) {
	c.ime = v
}

// Halt the CPU.
func (c *CPU) triggerHalt() {
	// TODO.
}
