package joypad

import (
	"github.com/ruiqimao/go-gb-emu/utils"
)

// Input interface.
// Button() returns which button the event is for:
//   A:      0
//   B:      1
//   Select: 2
//   Start:  3
//   Right:  4
//   Left:   5
//   Up:     6
//   Down:   7
// State() is whether the button was pressed (true) or released (false).
type Input interface {
	Button() int
	State() bool
}

// Handle an input event.
func (j *Joypad) Handle(input Input) {
	j.input = utils.SetBit(j.input, input.Button(), !input.State())
	j.updateJOYP()
}
