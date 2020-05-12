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
func (c *Cpu) handleInterrupts(ime bool) int {
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
func (c *Cpu) RequestInterrupt(interrupt int) {
	iF := c.gb.mem.Read(AddrIF)
	iF = utils.SetBit(iF, interrupt, true)
	c.gb.mem.Write(AddrIF, iF)
}

// Update timers for the given number of cycles.
func (c *Cpu) updateTimers(cycles int) {
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
