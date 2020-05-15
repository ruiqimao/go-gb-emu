package cpu

import (
	"github.com/ruiqimao/go-gb-emu/utils"
)

type Interrupt int
const (
	InterruptVBlank = 0
	InterruptStat   = 1
	InterruptTimer  = 2
	InterruptSerial = 3
	InterruptJoypad = 4
)

// Handle interrupts.
func (c *CPU) handleInterrupts(ime bool) {
	ief := c.iE & c.iF

	// If IE & IF == 0, then there are no interrupts to be handled.
	if ief == 0 {
		return
	}

	// There are enabled interrupts being requested, so exit halt mode.
	c.halt = false

	// If IME is disabled, do not dispatch.
	if !c.ime {
		return
	}

	// Dispatch the highest priority interrupt.
	for i := 0; i < 5; i++ {
		if utils.GetBit(ief, i) {
			// Execute two no-ops.
			c.incrementMCycle()
			c.incrementMCycle()

			// Disable interrupt master enable.
			c.ime = false

			// Unassert interrupt.
			c.iF = utils.SetBit(c.iF, i, false)

			// Push the program counter onto the stack.
			c.pushSP(c.pc)

			// Jump to the interrupt handler.
			c.incrementMCycle()
			c.pc = 0x0040 + uint16(i)*0x08
		}
	}
}

// Halt the CPU.
// Perform a halt. Handles HALT bug, documented in section 4.10 of
// https://github.com/AntonioND/giibiiadvance/blob/master/docs/TCAGBD.pdf.
func (c *CPU) triggerHalt() {
	if c.ime {
		// HALT is executed normally.
		c.halt = true
	} else {
		if c.iE&c.iF&0x1f == 0x0 {
			// HALT is executed normally.
			c.halt = true
		} else {
			// HALT is not executed. Instead, the halt bug is triggered, and the CPU will fail to
			// increment the program counter on the next instruction.
			c.haltBug = true
		}
	}
}

// Set the interrupt master enable.
func (c *CPU) setIME(v bool) {
	c.ime = v
}

// Request an interrupt.
func (c *CPU) RequestInterrupt(interrupt Interrupt) {
	c.iF = utils.SetBit(c.iF, int(interrupt), true)
}

// Get the interrupt enable register.
func (c *CPU) IE() uint8 {
	return c.iE
}

// Set the interrupt enable register.
func (c *CPU) SetIE(v uint8) {
	c.iE = v
}

// Get the interrupt assert register.
func (c *CPU) IF() uint8 {
	return c.iF | 0xe0 // Upper 3 bits are always 1.
}

// Set the interrupt assert register.
func (c *CPU) SetIF(v uint8) {
	c.iF = v
}
