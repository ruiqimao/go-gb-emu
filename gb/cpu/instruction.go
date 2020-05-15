package cpu

// InstructionIO is the set of functions an Instruction can use to interface with the CPU and MMU.
type InstructionIO struct {
	// Register access.
	Load    func(reg Register) uint8
	Load16  func(reg Register16) uint16
	Store   func(reg Register, v uint8)
	Store16 func(reg Register16, v uint16)
	GetFlag func(flag Flag) bool
	SetFlag func(flag Flag, v bool)

	// Memory access.
	Read    func(addr uint16) uint8
	Read16  func(addr uint16) uint16
	Write   func(addr uint16, v uint8)
	Write16 func(addr uint16, v uint16)

	// No-op used for cycle counting.
	Nop func()
}

// An Instruction represents a single CPU instruction.
type Instruction func(InstructionIO)

// Create the CPU instruction set.
func (c *CPU) initInstructionSet() {
	// TODO.
}
