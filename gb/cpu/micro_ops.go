package cpu

import (
	"github.com/ruiqimao/go-gb-emu/utils"
)

// Fundamental micro-ops that retrieve and store data.
type OpSrc func(InstructionIO) uint8
type OpSrc16 func(InstructionIO) uint16
type OpDst func(InstructionIO, uint8)
type OpDst16 func(InstructionIO, uint16)
type OpFlagSrc func(InstructionIO) bool
type OpFlagDst func(InstructionIO, bool)

// Generate a micro-op that loads the value in a register.
func opLoad(reg Register) OpSrc {
	return func(io InstructionIO) uint8 {
		return io.Load(reg)
	}
}

// Generate a micro-op that loads the value in a 16-bit register.
func opLoad16(reg Register16) OpSrc16 {
	return func(io InstructionIO) uint16 {
		return io.Load16(reg)
	}
}

// Generate a micro-op that writes to a register.
func opStore(reg Register) OpDst {
	return func(io InstructionIO, v uint8) {
		io.Store(reg, v)
	}
}

// Generate a micro-op that writes to a 16-bit register.
func opStore16(reg Register16) OpDst16 {
	return func(io InstructionIO, v uint16) {
		io.Store16(reg, v)
	}
}

func opFlag(flag Flag) OpFlagSrc {
	return func(io InstructionIO) bool {
		return io.GetFlag(flag)
	}
}

func opSetFlag(flag Flag) OpFlagDst {
	return func(io InstructionIO, v bool) {
		io.SetFlag(flag, v)
	}
}

// Generate a micro-op that reads from memory at the value in the register.
func opRead(reg Register16) OpSrc {
	return func(io InstructionIO) uint8 {
		return io.Read(io.Load16(reg))
	}
}

// Generate a micro-op that stores into memory at the value in the register.
func opWrite(reg Register16) OpDst {
	return func(io InstructionIO, v uint8) {
		io.Write(io.Load16(reg), v)
	}
}

// Generate a micro-op that gets the program counter.
func opPC() OpSrc16 {
	return func(io InstructionIO) uint16 {
		return io.PC()
	}
}

// Generate a micro-op that sets the program counter.
func opSetPC() OpDst16 {
	return func(io InstructionIO, v uint16) {
		io.SetPC(v)
	}
}

// Generate a micro-op that reads an immediate value.
func opImmediate() OpSrc {
	return func(io InstructionIO) uint8 {
		pc := io.PC()
		v := io.Read(pc)
		io.SetPC(pc + 1)
		return v
	}
}

// Generate a micro-op that reads an immediate 16-bit value.
func opImmediate16() OpSrc16 {
	return func(io InstructionIO) uint16 {
		pc := io.PC()
		hi := io.Read(pc)
		lo := io.Read(pc+1)
		io.SetPC(pc + 2)
		return utils.CombineBytes(hi, lo)
	}
}

// Generate a micro-op that gets the stack pointer.
func opSP() OpSrc16 {
	return func(io InstructionIO) uint16 {
		return io.SP()
	}
}

// Generate a micro-op that sets the stack pointer.
func opSetSP() OpDst16 {
	return func(io InstructionIO, v uint16) {
		io.SetSP(v)
	}
}

// Generate a micro-op that pops a value from the stack.
func opPop() OpSrc16 {
	return func(io InstructionIO) uint16 {
		sp := io.SP()
		hi := io.Read(sp)
		lo := io.Read(sp+1)
		io.SetSP(sp + 2)
		return utils.CombineBytes(hi, lo)
	}
}

// Generate a micro-op that pushes a value into the stack.
func opPush() OpDst16 {
	return func(io InstructionIO, v uint16) {
		sp := io.SP()
		io.SetSP(sp - 2)
		hi, lo := utils.SplitShort(v)
		io.Write(sp - 2, hi)
		io.Write(sp - 1, lo)
	}
}
