package mmu

// Joypad interface.
type Joypad interface {
	JOYP() uint8

	SetJOYP(uint8)
}

type JoypadBus struct {
	RequestInterrupt (int)
}
