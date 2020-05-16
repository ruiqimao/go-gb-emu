package gb

import (
	"github.com/ruiqimao/go-gb-emu/utils"
)

// Joypad flags.
const (
	JoypadSelectDPad   = 4
	JoypadSelectButton = 5

	JoypadA      = 0
	JoypadB      = 1
	JoypadSelect = 2
	JoypadStart  = 3
	JoypadRight  = 4
	JoypadLeft   = 5
	JoypadUp     = 6
	JoypadDown   = 7
)

// Joypad interrupt.
const (
	InterruptJoypad = 4
)

// Input is an input event for the joypad.
type Input struct {
	button int
	state  bool
}

type Joypad struct {
	gb *GameBoy

	// Input lines.
	// Lower 4 bits are for buttons. Higher 4 bits are for the D-Pad.
	input uint8

	// JOYP register.
	joyp uint8
}

func NewInput(button int, state bool) Input {
	return Input{
		button: button,
		state:  state,
	}
}

func NewJoypad(gb *GameBoy) *Joypad {
	j := &Joypad{
		gb:    gb,
		input: 0xff, // Input is all pulled high by default.
	}
	return j
}

// Handle an input event.
func (j *Joypad) Handle(input Input) {
	j.input = utils.SetBit(j.input, input.button, !input.state)
	j.update()
}

// Get the JOYP register.
func (j *Joypad) JOYP() uint8 {
	return j.joyp
}

// Set the JOYP register.
func (j *Joypad) SetJOYP(v uint8) {
	// Set the select bits.
	j.joyp = utils.CopyBit(v, j.joyp, JoypadSelectButton)
	j.joyp = utils.CopyBit(v, j.joyp, JoypadSelectDPad)
	j.update()
}

// Update the JOYP register.
func (j *Joypad) update() {
	// Save the old JOYP input line bits.
	oldBits := j.joyp & 0x0f

	// Reset the lower 4 bits of JOYP.
	j.joyp |= 0x0f

	// Turn off bits depending on the select lines.
	if !utils.GetBit(j.joyp, JoypadSelectDPad) {
		// D-Pad lines need to be shifted down 4 bits.
		j.joyp &= j.input >> 4
	}
	if !utils.GetBit(j.joyp, JoypadSelectButton) {
		j.joyp &= j.input & 0x0f
	}

	// Look for a falling edge in the input line bits.
	newBits := j.joyp & 0x0f
	if newBits^(oldBits|newBits) != 0x0 {
		// Trigger a joypad interrupt.
		j.gb.cpu.RequestInterrupt(InterruptJoypad)
	}
}
