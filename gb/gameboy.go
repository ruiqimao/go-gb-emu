package gb

type GameBoy struct {
	cpu *Cpu
	mem *Memory

	// Instruction set.
	instructions [0x200]Instruction

	// Boot ROM.
	boot [256]uint8
}

func NewGameBoy() (*GameBoy, error) {
	gb := &GameBoy{}

	// Create the components.
	gb.cpu = NewCpu(gb)
	gb.mem = NewMemory(gb)

	// Create the instruction set.
	gb.createInstructionSet()

	return gb, nil
}
