package gb

import (
	"github.com/ruiqimao/go-gb-emu/utils"
)

// The bits corresponding to each type of interrupt.
const (
	IntVBlank = 0
	IntStat   = 1
	IntTimer  = 2
	IntSerial = 3
	IntJoypad = 4
)

// Handle interrupts. Returns how many cycles were used.
func (c *CPU) handleInterrupts(ime bool) int {
	iE := c.gb.mem.Read(AddrIE)
	iF := c.gb.mem.Read(AddrIF)
	ief := iE & iF

	// If IE & IF == 0, then there are no interrupts to be handled.
	if ief == 0 {
		return 0
	}

	// There are enabled interrupts being requested, so exit halt mode.
	c.SetHalt(false)

	// If IME is disabled, do not dispatch.
	if !ime {
		return 0
	}

	// Dispatch the highest priority interrupt.
	for i := 0; i < 5; i++ {
		if utils.GetBit(ief, i) {
			// Disable interrupt master enable.
			c.SetIME(false)

			// Unassert interrupt.
			c.gb.mem.Write(AddrIF, utils.SetBit(iF, i, false))

			// Push the program counter onto the stack.
			c.PushSP(c.PC())

			// Jump to the interrupt handler.
			c.SetPC(0x0040 + uint16(i)*0x08)

			return 20
		}
	}

	return 0
}

// Request an interrupt.
func (c *CPU) RequestInterrupt(interrupt int) {
	iF := c.gb.mem.Read(AddrIF)
	iF = utils.SetBit(iF, interrupt, true)
	c.gb.mem.Write(AddrIF, iF)
}

// Get the IF register.
func (c *CPU) IF() uint8 {
	// Upper 3 bits are always high.
	return c.gb.mem.IO[AddrIF-AddrIO] | 0xe0
}
