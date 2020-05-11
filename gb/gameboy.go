package gb

type GameBoy struct {
	cpu *Cpu
	mem *Memory

	// Boot ROM.
	boot [256]uint8
}

func NewGameBoy() (*GameBoy, error) {
	gb := &GameBoy{}

	// Create the components.
	gb.cpu = NewCpu(gb)
	gb.mem = NewMemory(gb)

	return gb, nil
}
