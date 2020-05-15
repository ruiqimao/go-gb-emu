package cpu

type CPU struct {
	// Registers.
	rg [0x8]uint8

	// Stack pointer and program counter.
	sp uint16
	pc uint16

	// Halt flags.
	halt    bool
	haltBug bool

	// Interrupt enable flags.
	ime bool
	ie  bool

	// Timer counter and overflow flag.
	ic uint16
	of bool

	iio InstructionIO
}

// Create a new CPU.
func NewCPU() (*CPU, error) {
	c := &CPU{}

	// Create the InstructionIO.
	c.iio = InstructionIO{
		Load:    c.getRegister,
		Load16:  c.getRegister16,
		Store:   c.setRegister,
		Store16: c.setRegister16,
		GetFlag: c.getFlag,
		SetFlag: c.setFlag,

		Read:    c.readMemory,
		Read16:  c.readMemory16,
		Write:   c.writeMemory,
		Write16: c.writeMemory16,

		Nop: c.incrementMCycle,
	}

	// Initialize the instruction set.
	c.initInstructionSet()

	return c, nil
}
