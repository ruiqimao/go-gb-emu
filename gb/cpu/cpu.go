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

	println("  0 1 2 3 4 5 6 7 8 9 A B C D E F")
	for k := 0; k < 2; k++ {
		for i := 0; i < 16; i++ {
			switch i {
			case 10:
				print("A ")
			case 11:
				print("B ")
			case 12:
				print("C ")
			case 13:
				print("D ")
			case 14:
				print("E ")
			case 15:
				print("F ")
			default:
				print(i)
				print(" ")
			}
			for j := 0; j < 16; j++ {
				if c.instructions[i*16+j+k*256] != nil {
					print("X ")
				} else {
					print("- ")
				}
			}
			println()
		}
	}

	return c
}
