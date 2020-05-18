package main

type Button int

const (
	ButtonA      Button = 0
	ButtonB             = 1
	ButtonSelect        = 2
	ButtonStart         = 3
	ButtonRight         = 4
	ButtonLeft          = 5
	ButtonUp            = 6
	ButtonDown          = 7
)

type Input struct {
	button Button
	state  bool
}

func (i Input) Button() int {
	return int(i.button)
}

func (i Input) State() bool {
	return i.state
}
