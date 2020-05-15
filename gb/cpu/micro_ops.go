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

// Generate a micro-op that reads the value of a flag.
func opFlag(flag Flag) OpFlagSrc {
	return func(io InstructionIO) bool {
		return io.GetFlag(flag)
	}
}

// Generate a micro-op that inverts the value of a flag.
func opNotFlag(flag Flag) OpFlagSrc {
	return func(io InstructionIO) bool {
		return !io.GetFlag(flag)
	}
}

// Generate a micro-op that sets the value of a flag.
func opSetFlag(flag Flag) OpFlagDst {
	return func(io InstructionIO, v bool) {
		io.SetFlag(flag, v)
	}
}

// Generate a micro-op that always resolves to true.
func opTrue() OpFlagSrc {
	return func(io InstructionIO) bool {
		return true
	}
}

// Generate a micro-op that reads from memory.
func opRead(addr OpSrc16) OpSrc {
	return func(io InstructionIO) uint8 {
		return io.Read(addr(io))
	}
}

// Generate a micro-op that reads 16 bits from memory.
func opRead16(addr OpSrc16) OpSrc16 {
	return func(io InstructionIO) uint16 {
		a := addr(io)
		hi := io.Read(a)
		lo := io.Read(a + 1)
		return utils.CombineBytes(hi, lo)
	}
}

// Generate a micro-op that stores into memory.
func opWrite(addr OpSrc16) OpDst {
	return func(io InstructionIO, v uint8) {
		io.Write(addr(io), v)
	}
}

// Generate a micro-op that stores 16 bits into memory.
func opWrite16(addr OpSrc16) OpDst16 {
	return func(io InstructionIO, v uint16) {
		a := addr(io)
		hi, lo := utils.SplitShort(v)
		io.Write(a, hi)
		io.Write(a+1, lo)
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
		lo := io.Read(pc + 1)
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

// Generate a micro-op that reads from memory at the value in HL, then increments HL.
func opReadHLI() OpSrc {
	return func(io InstructionIO) uint8 {
		hl := io.Load16(RegisterHL)
		v := io.Read(hl)
		io.Store16(RegisterHL, hl+1)
		return v
	}
}

// Generate a micro-op that reads from memory at the value in HL, then decrements HL.
func opReadHLD() OpSrc {
	return func(io InstructionIO) uint8 {
		hl := io.Load16(RegisterHL)
		v := io.Read(hl)
		io.Store16(RegisterHL, hl-1)
		return v
	}
}

// Generate a micro-op that writes to memory at the value in HL, then increments HL.
func opWriteHLI() OpDst {
	return func(io InstructionIO, v uint8) {
		hl := io.Load16(RegisterHL)
		io.Write(hl, v)
		io.Store16(RegisterHL, hl+1)
	}
}

// Generate a micro-op that writes to memory at the value in HL, then decrements HL.
func opWriteHLD() OpDst {
	return func(io InstructionIO, v uint8) {
		hl := io.Load16(RegisterHL)
		io.Write(hl, v)
		io.Store16(RegisterHL, hl-1)
	}
}

// Generate a micro-op that adds 0xFF00 to the value.
func opHigh(src OpSrc) OpSrc16 {
	return func(io InstructionIO) uint16 {
		return 0xff00 + uint16(src(io))
	}
}

// Generate a micro-op that performs a signed add.
func opSAdd(srcA OpSrc16, srcB OpSrc) OpSrc16 {
	return func(io InstructionIO) uint16 {
		a := srcA(io)
		b := srcB(io)
		r := int32(a) + int32(int8(b))

		io.SetFlag(FlagZ, false)
		io.SetFlag(FlagN, false)
		io.SetFlag(FlagH, (uint8(a)&0xf)-(b&0xf) > 0xf)
		io.SetFlag(FlagC, r > 0xff)

		return uint16(r)
	}
}
