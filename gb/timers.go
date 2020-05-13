package gb

import (
	"github.com/ruiqimao/go-gb-emu/utils"
)

// Update timers for the given number of cycles.
func (c *CPU) updateTimers(cycles int) {
	// Increment the internal counter. We keep track of the initial counter for per cycle processing
	// later.
	ic := c.IC()
	c.SetIC(uint16((int(ic) + cycles)))

	// Check if timers are enabled.
	tac := c.gb.mem.Read(AddrTAC)
	if !utils.GetBit(tac, 2) {
		return
	}

	// Determine the divisor bit.
	var divisor int
	switch tac & 0x3 {
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

	// We know cycles is always a multiple of 4, so we can process per 4 cycles.
	for cycles > 0 {
		// Get the current TIMA value.
		tima := c.gb.mem.Read(AddrTIMA)

		// If there was an overflow last cycle, reset TIMA.
		// This ignores any writes to TIMA that may have occurred during this cycle.
		if c.of {
			// Reset TIMA to TMA.
			tima = c.gb.mem.Read(AddrTMA)
			c.RequestInterrupt(IntTimer)
			c.of = false
		}

		// Look for a falling edge.
		falling := utils.GetBit16(ic, divisor) && !utils.GetBit16(ic+4, divisor)

		// If there is a falling edge, increment TIMA.
		if falling {
			tima++

			// If tima is now 0, there was an overflow.
			if tima == 0x0 {
				c.of = true
			}
		}

		// Update TIMA.
		c.gb.mem.Write(AddrTIMA, tima)

		ic += 4
		cycles -= 4
	}
}

// Get the DIV register.
func (c *CPU) DIV() uint8 {
	// DIV register is upper 8 bits of internal counter.
	return uint8(c.IC() >> 8)
}

// Set the DIV register.
func (c *CPU) SetDIV(v uint8) {
	// DIV maps directly to internal counter.
	lo := uint8(c.IC())
	c.SetIC(utils.CombineBytes(v, lo))
}

// Set the TAC register.
func (c *CPU) SetTAC(v uint8) {
	// Only lower 3 bits are writable.
	c.gb.mem.IO[AddrTAC-AddrIO] = v & 0x07
}
