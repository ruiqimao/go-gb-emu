package cpu

import (
	"github.com/ruiqimao/go-gb-emu/utils"
)

// Memory interface for the CPU.
type MemoryIO interface {
	Read(uint16) uint8
	Write(uint16, uint8)
}

// Read a byte from memory.
func (c *CPU) readMemory(addr uint16) uint8 {
	c.incrementMCycle()
	return c.mmu.Read(addr)
}

// Write a byte to memory.
func (c *CPU) writeMemory(addr uint16, v uint8) {
	c.incrementMCycle()
	c.mmu.Write(addr, v)
}

// Read a short from memory.
func (c *CPU) readMemory16(addr uint16) uint16 {
	hi := c.readMemory(addr + 1)
	lo := c.readMemory(addr)
	return utils.CombineBytes(hi, lo)
}

// Write a short to memory.
func (c *CPU) writeMemory16(addr uint16, v uint16) {
	hi, lo := utils.SplitShort(v)
	c.writeMemory(addr, lo)
	c.writeMemory(addr+1, hi)
}
