package cpu

import (
	"fmt"
)

type CPU struct {
	// Memory bus.
	mmu MemoryIO

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

	// Timer registers.
	ic   uint16
	tima uint8
	tma  uint8
	tac  uint8
	of   bool

	// Number of clocks the current instruction is using.
	clocks int

	// Instruction set.
	instructions [0x200]Instruction
	iio          InstructionIO
}

// Create a new CPU.
func NewCPU(mmu MemoryIO) *CPU {
	c := &CPU{
		mmu: mmu,
	}

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

// Do a CPU step. Returns how many clocks were used.
func (c *CPU) Step() (int, error) {
	c.clocks = 0

	// Save the value of IME to use after the instruction has been executed.
	ime := c.ime

	// Execute an instruction.
	if !c.halt {
		op := uint16(c.popPC())
		if op == 0xcb {
			// CB prefixed operations are offset by 256 in the instruction set.
			op = uint16(c.popPC()) + 0x100
		}
		instruction := c.instructions[op]
		if instruction == nil {
			return c.clocks, fmt.Errorf("Invalid op code %02x", op)
		}
		c.instructions[op](c.iio)
	} else {
		c.incrementMCycle()
	}

	// Handle interrupts.
	c.handleInterrupts(ime)

	return c.clocks, nil
}
