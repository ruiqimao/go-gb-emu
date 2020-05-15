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

	// Instruction set.
	instructions [0x200]Instruction
	iio InstructionIO
}

// Create a new CPU.
func NewCPU() *CPU {
	c := &CPU{}

	// Create the InstructionIO.
	c.iio = InstructionIO{
		Load:    c.getRegister,
		Load16:  c.getRegister16,
		Store:   c.setRegister,
		Store16: c.setRegister16,
		GetFlag: c.getFlag,
		SetFlag: c.setFlag,

		Read:  c.readMemory,
		Write: c.writeMemory,

		PC:    c.getPC,
		SetPC: c.setPC,

		SP:    c.getSP,
		SetSP: c.setSP,

		Nop: c.incrementMCycle,
	}

	// Initialize the instruction set.
	c.initInstructionSet()

	return c
}
