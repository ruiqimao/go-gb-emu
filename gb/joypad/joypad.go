package joypad

import (
	"github.com/ruiqimao/go-gb-emu/utils"
)

// Joypad flags.
const (
	FlagSelectDPad   = 4
	FlagSelectButton = 5
)

type Joypad struct {
	mmu MMU

	// Input lines.
	// Buttons are stored in the lower 4 bits, and the d-pad is stored in the upper 4 bits.
	input uint8

	// JOYP register.
	joyp uint8
}

func NewJoypad() *Joypad {
	j := &Joypad{
		input: 0xff, // Pull all lines high by default.
	}
	return j
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
		j.interruptJoypad()
	}
}

func (j *Joypad) AttachMMU(mmu MMU) {
	j.mmu = mmu
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
