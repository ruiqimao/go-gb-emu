package cpu

import (
	"github.com/ruiqimao/go-gb-emu/utils"
)

// Increment by a machine cycle.
func (c *CPU) incrementMCycle() {
	c.clocks += 4
	c.updateTimers()
}

// Update timers for one machine cycle.
func (c *CPU) updateTimers() {
	// Update the internal counter.
	c.ic += 4

	// Check if timers are enabled.
	if !utils.GetBit(c.tac, 2) {
		return
	}

	// Determine the divisor bit.
	var divisor int
	switch c.tac & 0x3 {
	case 0x0:
		// 4096 Hz, 1024 clocks per tick.
		divisor = 9
	case 0x1:
		// 262144 Hz, 16 clocks per tick.
		divisor = 3
	case 0x2:
		// 65536 Hz, 64 clocks per tick.
		divisor = 5
	case 0x3:
		// 16386 Hz, 256 clocks per tick.
		divisor = 7
	}

	// If there was an overflow last cycle, reset TIMA.
	// This ignores any writes to TIMA that may have occurred during this cycle.
	if c.of {
		// Reset TIMA to TMA.
		c.tima = c.tma
		c.RequestInterrupt(InterruptTimer)
		c.of = false
	}

	// Look for a falling edge.
	falling := utils.GetBit16(c.ic-4, divisor) && !utils.GetBit16(c.ic, divisor)

	// If there is a falling edge, increment TIMA.
	if falling {
		c.tima++

		// If tima is now 0, there was an overflow.
		if c.tima == 0x0 {
			c.of = true
		}
	}
}

// Get the DIV register.
func (c *CPU) DIV() uint8 {
	// DIV register is upper 8 bits of internal counter.
	return uint8(c.ic >> 8)
}

// Set the DIV register.
func (c *CPU) SetDIV(v uint8) {
	// Set the upper 8 bits of internal counter to DIV.
	c.ic = (uint16(v) << 8) | (c.ic & 0x0f)
}

// Get the TIMA register.
func (c *CPU) TIMA() uint8 {
	return c.tima
}

// Set the TIMA register.
func (c *CPU) SetTIMA(v uint8) {
	c.tima = v
}

// Get the TMA register.
func (c *CPU) TMA() uint8 {
	return c.tma
}

// Set the TMA register.
func (c *CPU) SetTMA(v uint8) {
	c.tma = v
}

// Get the TAC register.
func (c *CPU) TAC() uint8 {
	return c.tac
}

// Set the TAC register.
func (c *CPU) SetTAC(v uint8) {
	// Only lower 3 bits are writable.
	c.tac = v & 0x07
}
