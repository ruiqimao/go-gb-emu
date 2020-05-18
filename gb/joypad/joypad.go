package joypad

import (
	"github.com/ruiqimao/go-gb-emu/utils"
)

// Joypad flags.
const (
	FlagSelectDPad   = 4
	FlagSelectButton = 5

	FlagA      = 0
	FlagB      = 1
	FlagSelect = 2
	FlagStart  = 3
	FlagRight  = 4
	FlagLeft   = 5
	FlagUp     = 6
	FlagDown   = 7
)

type Joypad struct {
	mmu MMU

	// JOYP register.
	joyp uint8
}

func NewJoypad() *Joypad {
	j := &Joypad{}
	return j
}

// Get the JOYP register.
func (j *Joypad) JOYP() uint8 {
	return j.joyp
}

// Set the JOYP register.
func (j *Joypad) SetJOYP(v uint8) {
	// Set the select bits.
	j.joyp = utils.CopyBit(v, j.joyp, FlagSelectButton)
	j.joyp = utils.CopyBit(v, j.joyp, FlagSelectDPad)
	j.updateJOYP()
}

// Update the JOYP register.
func (j *Joypad) updateJOYP() {
	// Save the old JOYP input line bits.
	oldBits := j.joyp & 0x0f

	// Reset the lower 4 bits of JOYP.
	j.joyp |= 0x0f

	// Turn off bits depending on the select lines.
	if !utils.GetBit(j.joyp, FlagSelectDPad) {
		// D-Pad lines need to be shifted down 4 bits.
		j.joyp &= j.input >> 4
	}
	if !utils.GetBit(j.joyp, FlagSelectButton) {
		j.joyp &= j.input & 0x0f
	}

	// Look for a falling edge in the input line bits.
	newBits := j.joyp & 0x0f
	if newBits^(oldBits|newBits) != 0x0 {
		// Trigger a joypad interrupt.
		j.interruptJOYP()
	}
}
