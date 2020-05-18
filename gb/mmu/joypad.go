package mmu

// Joypad interface.
type Joypad interface {
	JOYP() uint8

	SetJOYP(uint8)
}

type JoypadBus struct {
	mmu *MMU
}

func (b *JoypadBus) RequestInterrupt(interrupt int) {
	b.mmu.requestInterrupt(interrupt)
}
