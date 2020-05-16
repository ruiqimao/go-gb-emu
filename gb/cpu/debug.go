package cpu

// The functions in this file should be used only for debugging purposes.

func (c *CPU) PC() uint16 {
	return c.pc
}

func (c *CPU) SetPC(v uint16) {
	c.setPC(v)
}

func (c *CPU) SP() uint16 {
	return c.sp
}

func (c *CPU) SetSP(v uint16) {
	c.setSP(v)
}

func (c *CPU) GetRegister(reg Register) uint8 {
	return c.getRegister(reg)
}

func (c *CPU) SetRegister(reg Register, v uint8) {
	c.setRegister(reg, v)
}

func (c *CPU) GetRegister16(reg Register16) uint16 {
	return c.getRegister16(reg)
}

func (c *CPU) SetRegister16(reg Register16, v uint16) {
	c.setRegister16(reg, v)
}

func (c *CPU) GetFlag(flag Flag) bool {
	return c.getFlag(flag)
}

func (c *CPU) SetFlag(flag Flag, v bool) {
	c.setFlag(flag, v)
}
