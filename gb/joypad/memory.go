package joypad

// Joypad interrupt.
const (
	InterruptJoypad = 4
)

// MMU interface.
type MMU interface {
	RequestInterrupt(int)
}

func (j *Joypad) interruptJOYP() {
	if j.mmu != nil {
		j.mmu.RequestInterrupt(InterruptJoypad)
	}
}
