package cpu

import (
	"github.com/ruiqimao/go-gb-emu/utils"
)

// Read a byte from memory.
func (c *CPU) readMemory(addr uint16) uint8 {
	// TODO.
	return 0x00
}

// Write a byte to memory.
func (c *CPU) writeMemory(addr uint16, v uint8) {
	// TODO.
}

// Read a short from memory.
func (c *CPU) readMemory16(addr uint16) uint16 {
	hi := c.readMemory(addr)
	lo := c.readMemory(addr + 1)
	return utils.CombineBytes(hi, lo)
}

// Write a short to memory.
func (c *CPU) writeMemory16(addr uint16, v uint16) {
	hi, lo := utils.SplitShort(v)
	c.writeMemory(addr, hi)
	c.writeMemory(addr+1, lo)
}
