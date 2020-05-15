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

	// Interrupt flags.
	ime bool
	iE  uint8
	iF  uint8

	// Timer counter and overflow flag.
	ic uint16
	of bool

	// Instruction set.
	instructions [0x200]Instruction
	iio          InstructionIO
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

		Read:    c.readMemory,
		Read16:  c.readMemory16,
		Write:   c.writeMemory,
		Write16: c.writeMemory16,

		PC:      c.getPC,
		SetPC:   c.setPC,
		PopPC:   c.popPC,
		PopPC16: c.popPC16,

		SP:     c.getSP,
		SetSP:  c.setSP,
		PopSP:  c.popSP,
		PushSP: c.pushSP,

		SetIME: c.setIME,
		Halt:   c.triggerHalt,

		Nop: c.incrementMCycle,
	}

	// Initialize the instruction set.
	c.initInstructionSet()

	return c
}
