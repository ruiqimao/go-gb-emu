package cpu

import (
	"github.com/ruiqimao/go-gb-emu/utils"
)

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
	Read  func(addr uint16) uint8
	Write func(addr uint16, v uint8)

	// Immediate value access.
	PC    func() uint16
	SetPC func(v uint16)

	// Stack access.
	SP    func() uint16
	SetSP func(v uint16)

	// No-op used for cycle counting.
	Nop func()
}

// An Instruction represents a single CPU instruction.
type Instruction func(InstructionIO)

// Create the CPU instruction set.
func (c *CPU) initInstructionSet() {
	c.instructions = [0x200]Instruction{
		// 8-bit loads.
		0x06: opLD(opStore(RegisterB), opImmediate()),
		0x0e: opLD(opStore(RegisterC), opImmediate()),
		0x16: opLD(opStore(RegisterD), opImmediate()),
		0x1e: opLD(opStore(RegisterE), opImmediate()),
		0x26: opLD(opStore(RegisterH), opImmediate()),
		0x2e: opLD(opStore(RegisterL), opImmediate()),

		0x40: opLD(opStore(RegisterB), opLoad(RegisterB)),
		0x41: opLD(opStore(RegisterB), opLoad(RegisterC)),
		0x42: opLD(opStore(RegisterB), opLoad(RegisterD)),
		0x43: opLD(opStore(RegisterB), opLoad(RegisterE)),
		0x44: opLD(opStore(RegisterB), opLoad(RegisterH)),
		0x45: opLD(opStore(RegisterB), opLoad(RegisterL)),
		0x46: opLD(opStore(RegisterB), opRead(opLoad16(RegisterHL))),
		0x47: opLD(opStore(RegisterB), opLoad(RegisterA)),

		0x48: opLD(opStore(RegisterC), opLoad(RegisterB)),
		0x49: opLD(opStore(RegisterC), opLoad(RegisterC)),
		0x4a: opLD(opStore(RegisterC), opLoad(RegisterD)),
		0x4b: opLD(opStore(RegisterC), opLoad(RegisterE)),
		0x4c: opLD(opStore(RegisterC), opLoad(RegisterH)),
		0x4d: opLD(opStore(RegisterC), opLoad(RegisterL)),
		0x4e: opLD(opStore(RegisterC), opRead(opLoad16(RegisterHL))),
		0x4f: opLD(opStore(RegisterC), opLoad(RegisterA)),

		0x50: opLD(opStore(RegisterD), opLoad(RegisterB)),
		0x51: opLD(opStore(RegisterD), opLoad(RegisterC)),
		0x52: opLD(opStore(RegisterD), opLoad(RegisterD)),
		0x53: opLD(opStore(RegisterD), opLoad(RegisterE)),
		0x54: opLD(opStore(RegisterD), opLoad(RegisterH)),
		0x55: opLD(opStore(RegisterD), opLoad(RegisterL)),
		0x56: opLD(opStore(RegisterD), opRead(opLoad16(RegisterHL))),
		0x57: opLD(opStore(RegisterD), opLoad(RegisterA)),

		0x58: opLD(opStore(RegisterE), opLoad(RegisterB)),
		0x59: opLD(opStore(RegisterE), opLoad(RegisterC)),
		0x5a: opLD(opStore(RegisterE), opLoad(RegisterD)),
		0x5b: opLD(opStore(RegisterE), opLoad(RegisterE)),
		0x5c: opLD(opStore(RegisterE), opLoad(RegisterH)),
		0x5d: opLD(opStore(RegisterE), opLoad(RegisterL)),
		0x5e: opLD(opStore(RegisterE), opRead(opLoad16(RegisterHL))),
		0x5f: opLD(opStore(RegisterE), opLoad(RegisterA)),

		0x60: opLD(opStore(RegisterH), opLoad(RegisterB)),
		0x61: opLD(opStore(RegisterH), opLoad(RegisterC)),
		0x62: opLD(opStore(RegisterH), opLoad(RegisterD)),
		0x63: opLD(opStore(RegisterH), opLoad(RegisterE)),
		0x64: opLD(opStore(RegisterH), opLoad(RegisterH)),
		0x65: opLD(opStore(RegisterH), opLoad(RegisterL)),
		0x66: opLD(opStore(RegisterH), opRead(opLoad16(RegisterHL))),
		0x67: opLD(opStore(RegisterH), opLoad(RegisterA)),

		0x68: opLD(opStore(RegisterL), opLoad(RegisterB)),
		0x69: opLD(opStore(RegisterL), opLoad(RegisterC)),
		0x6a: opLD(opStore(RegisterL), opLoad(RegisterD)),
		0x6b: opLD(opStore(RegisterL), opLoad(RegisterE)),
		0x6c: opLD(opStore(RegisterL), opLoad(RegisterH)),
		0x6d: opLD(opStore(RegisterL), opLoad(RegisterL)),
		0x6e: opLD(opStore(RegisterL), opRead(opLoad16(RegisterHL))),
		0x6f: opLD(opStore(RegisterL), opLoad(RegisterA)),

		0x70: opLD(opWrite(opLoad16(RegisterHL)), opLoad(RegisterB)),
		0x71: opLD(opWrite(opLoad16(RegisterHL)), opLoad(RegisterC)),
		0x72: opLD(opWrite(opLoad16(RegisterHL)), opLoad(RegisterD)),
		0x73: opLD(opWrite(opLoad16(RegisterHL)), opLoad(RegisterE)),
		0x74: opLD(opWrite(opLoad16(RegisterHL)), opLoad(RegisterH)),
		0x75: opLD(opWrite(opLoad16(RegisterHL)), opLoad(RegisterL)),
		0x77: opLD(opWrite(opLoad16(RegisterHL)), opLoad(RegisterA)),
		0x36: opLD(opWrite(opLoad16(RegisterHL)), opImmediate()),

		0x78: opLD(opStore(RegisterA), opLoad(RegisterB)),
		0x79: opLD(opStore(RegisterA), opLoad(RegisterC)),
		0x7a: opLD(opStore(RegisterA), opLoad(RegisterD)),
		0x7b: opLD(opStore(RegisterA), opLoad(RegisterE)),
		0x7c: opLD(opStore(RegisterA), opLoad(RegisterH)),
		0x7d: opLD(opStore(RegisterA), opLoad(RegisterL)),
		0x7e: opLD(opStore(RegisterA), opRead(opLoad16(RegisterHL))),
		0x7f: opLD(opStore(RegisterA), opLoad(RegisterA)),
		0x0a: opLD(opStore(RegisterA), opRead(opLoad16(RegisterBC))),
		0x1a: opLD(opStore(RegisterA), opRead(opLoad16(RegisterDE))),
		0xfa: opLD(opStore(RegisterA), opRead(opImmediate16())),
		0x3e: opLD(opStore(RegisterA), opImmediate()),

		0x2a: opLD(opStore(RegisterA), opReadHLI()),
		0x3a: opLD(opStore(RegisterA), opReadHLD()),
		0x22: opLD(opWriteHLI(), opLoad(RegisterA)),
		0x32: opLD(opWriteHLD(), opLoad(RegisterA)),

		0x02: opLD(opWrite(opLoad16(RegisterBC)), opLoad(RegisterA)),
		0x12: opLD(opWrite(opLoad16(RegisterDE)), opLoad(RegisterA)),
		0xea: opLD(opWrite(opImmediate16()), opLoad(RegisterA)),

		0xe0: opLD(opWrite(opHigh(opImmediate())), opLoad(RegisterA)),
		0xe2: opLD(opWrite(opHigh(opLoad(RegisterC))), opLoad(RegisterA)),
		0xf0: opLD(opStore(RegisterA), opRead(opHigh(opImmediate()))),
		0xf2: opLD(opStore(RegisterA), opRead(opHigh(opLoad(RegisterC)))),

		// 16-bit loads.
		0x01: opLD16(opStore16(RegisterBC), opImmediate16()),
		0x11: opLD16(opStore16(RegisterDE), opImmediate16()),
		0x21: opLD16(opStore16(RegisterHL), opImmediate16()),
		0x31: opLD16(opSetSP(), opImmediate16()),

		0xf9: opLD16(opSetSP(), opLoad16(RegisterHL)),
		0xf8: opLD16(opStore16(RegisterHL), opSAdd(opSP(), opLoad(RegisterHL), false)),
		0x08: opLD16(opWrite16(opImmediate16()), opSP()),

		0xc5: opPUSH(opLoad16(RegisterBC)),
		0xd5: opPUSH(opLoad16(RegisterDE)),
		0xe5: opPUSH(opLoad16(RegisterHL)),
		0xf5: opPUSH(opLoad16(RegisterAF)),

		0xc1: opPOP(opStore16(RegisterBC)),
		0xd1: opPOP(opStore16(RegisterDE)),
		0xe1: opPOP(opStore16(RegisterHL)),
		0xf1: opPOP(opStore16(RegisterAF)),

		// 8-bit arithmetic.
		0x80: opADD(opLoad(RegisterB), false),
		0x81: opADD(opLoad(RegisterC), false),
		0x82: opADD(opLoad(RegisterD), false),
		0x83: opADD(opLoad(RegisterE), false),
		0x84: opADD(opLoad(RegisterH), false),
		0x85: opADD(opLoad(RegisterL), false),
		0x86: opADD(opRead(opLoad16(RegisterHL)), false),
		0x87: opADD(opLoad(RegisterA), false),
		0xc6: opADD(opImmediate(), false),

		0x88: opADD(opLoad(RegisterB), true),
		0x89: opADD(opLoad(RegisterC), true),
		0x8a: opADD(opLoad(RegisterD), true),
		0x8b: opADD(opLoad(RegisterE), true),
		0x8c: opADD(opLoad(RegisterH), true),
		0x8d: opADD(opLoad(RegisterL), true),
		0x8e: opADD(opRead(opLoad16(RegisterHL)), true),
		0x8f: opADD(opLoad(RegisterA), true),
		0xce: opADD(opImmediate(), true),

		0x90: opSUB(opLoad(RegisterB), false),
		0x91: opSUB(opLoad(RegisterC), false),
		0x92: opSUB(opLoad(RegisterD), false),
		0x93: opSUB(opLoad(RegisterE), false),
		0x94: opSUB(opLoad(RegisterH), false),
		0x95: opSUB(opLoad(RegisterL), false),
		0x96: opSUB(opRead(opLoad16(RegisterHL)), false),
		0x97: opSUB(opLoad(RegisterA), false),
		0xd6: opSUB(opImmediate(), false),

		0x98: opSUB(opLoad(RegisterB), true),
		0x99: opSUB(opLoad(RegisterC), true),
		0x9a: opSUB(opLoad(RegisterD), true),
		0x9b: opSUB(opLoad(RegisterE), true),
		0x9c: opSUB(opLoad(RegisterH), true),
		0x9d: opSUB(opLoad(RegisterL), true),
		0x9e: opSUB(opRead(opLoad16(RegisterHL)), true),
		0x9f: opSUB(opLoad(RegisterA), true),
		0xde: opSUB(opImmediate(), true),

		0xa0: opAND(opLoad(RegisterB)),
		0xa1: opAND(opLoad(RegisterC)),
		0xa2: opAND(opLoad(RegisterD)),
		0xa3: opAND(opLoad(RegisterE)),
		0xa4: opAND(opLoad(RegisterH)),
		0xa5: opAND(opLoad(RegisterL)),
		0xa6: opAND(opRead(opLoad16(RegisterHL))),
		0xa7: opAND(opLoad(RegisterA)),
		0xe6: opAND(opImmediate()),

		0xa8: opXOR(opLoad(RegisterB)),
		0xa9: opXOR(opLoad(RegisterC)),
		0xaa: opXOR(opLoad(RegisterD)),
		0xab: opXOR(opLoad(RegisterE)),
		0xac: opXOR(opLoad(RegisterH)),
		0xad: opXOR(opLoad(RegisterL)),
		0xae: opXOR(opRead(opLoad16(RegisterHL))),
		0xaf: opXOR(opLoad(RegisterA)),
		0xee: opXOR(opImmediate()),

		0xb0: opOR(opLoad(RegisterB)),
		0xb1: opOR(opLoad(RegisterC)),
		0xb2: opOR(opLoad(RegisterD)),
		0xb3: opOR(opLoad(RegisterE)),
		0xb4: opOR(opLoad(RegisterH)),
		0xb5: opOR(opLoad(RegisterL)),
		0xb6: opOR(opRead(opLoad16(RegisterHL))),
		0xb7: opOR(opLoad(RegisterA)),
		0xf6: opOR(opImmediate()),

		0xb8: opCP(opLoad(RegisterB)),
		0xb9: opCP(opLoad(RegisterC)),
		0xba: opCP(opLoad(RegisterD)),
		0xbb: opCP(opLoad(RegisterE)),
		0xbc: opCP(opLoad(RegisterH)),
		0xbd: opCP(opLoad(RegisterL)),
		0xbe: opCP(opRead(opLoad16(RegisterHL))),
		0xbf: opCP(opLoad(RegisterA)),
		0xfe: opCP(opImmediate()),

		0x04: opINC(opStore(RegisterB), opLoad(RegisterB)),
		0x0c: opINC(opStore(RegisterC), opLoad(RegisterC)),
		0x14: opINC(opStore(RegisterD), opLoad(RegisterD)),
		0x1c: opINC(opStore(RegisterE), opLoad(RegisterE)),
		0x24: opINC(opStore(RegisterH), opLoad(RegisterH)),
		0x2c: opINC(opStore(RegisterL), opLoad(RegisterL)),
		0x34: opINC(opWrite(opLoad16(RegisterHL)), opRead(opLoad16(RegisterHL))),
		0x3c: opINC(opStore(RegisterA), opLoad(RegisterA)),

		0x05: opDEC(opStore(RegisterB), opLoad(RegisterB)),
		0x0d: opDEC(opStore(RegisterC), opLoad(RegisterC)),
		0x15: opDEC(opStore(RegisterD), opLoad(RegisterD)),
		0x1d: opDEC(opStore(RegisterE), opLoad(RegisterE)),
		0x25: opDEC(opStore(RegisterH), opLoad(RegisterH)),
		0x2d: opDEC(opStore(RegisterL), opLoad(RegisterL)),
		0x35: opDEC(opWrite(opLoad16(RegisterHL)), opRead(opLoad16(RegisterHL))),
		0x3d: opDEC(opStore(RegisterA), opLoad(RegisterA)),

		// 16-bit arithmetic.
		0x09: opADD16(opLoad16(RegisterBC)),
		0x19: opADD16(opLoad16(RegisterDE)),
		0x29: opADD16(opLoad16(RegisterHL)),
		0x39: opADD16(opSP()),

		0xe8: opLD16(opSetSP(), opSAdd(opSP(), opImmediate(), true)),

		0x03: opINC16(opStore16(RegisterBC), opLoad16(RegisterBC)),
		0x13: opINC16(opStore16(RegisterDE), opLoad16(RegisterDE)),
		0x23: opINC16(opStore16(RegisterHL), opLoad16(RegisterHL)),
		0x33: opINC16(opSetSP(), opSP()),

		0x0b: opDEC16(opStore16(RegisterBC), opLoad16(RegisterBC)),
		0x1b: opDEC16(opStore16(RegisterDE), opLoad16(RegisterDE)),
		0x2b: opDEC16(opStore16(RegisterHL), opLoad16(RegisterHL)),
		0x3b: opDEC16(opSetSP(), opSP()),
	}
}

// Generate an LD instruction.
func opLD(dst OpDst, src OpSrc) Instruction {
	return func(io InstructionIO) {
		// Store the source in the destination.
		dst(io, src(io))
	}
}

// Generate a 16-bit LD instruction.
func opLD16(dst OpDst16, src OpSrc16) Instruction {
	return func(io InstructionIO) {
		// Store the source in the destination.
		dst(io, src(io))
	}
}

// Generate a POP instruction.
func opPOP(dst OpDst16) Instruction {
	return func(io InstructionIO) {
		sp := io.SP()
		hi := io.Read(sp)
		lo := io.Read(sp+1)
		io.SetSP(sp + 2)
		dst(io, utils.CombineBytes(hi, lo))
	}
}

// Generate a PUSH instruction.
func opPUSH(src OpSrc16) Instruction {
	return func(io InstructionIO) {
		sp := io.SP()
		io.SetSP(sp - 2)
		hi, lo := utils.SplitShort(src(io))
		io.Nop()
		io.Write(sp - 2, hi)
		io.Write(sp - 1, lo)
	}
}

// Generate an ADD/ADC instruction.
func opADD(src OpSrc, useCarry bool) Instruction {
	return func(io InstructionIO) {
		a := io.Load(RegisterA)
		b := src(io)
		c := uint8(0)
		if useCarry && io.GetFlag(FlagC) {
			c = 1
		}
		r16 := uint16(a) + uint16(b) + uint16(c)
		r := uint8(r16)

		io.SetFlag(FlagZ, r == 0)
		io.SetFlag(FlagN, false)
		io.SetFlag(FlagH, (a&0xf)+(b&0xf)+(c&0xf) > 0xf)
		io.SetFlag(FlagC, r16 > 0xff)

		io.Store(RegisterA, r)
	}
}

// Generate a SUB/SBC instruction.
func opSUB(src OpSrc, useBorrow bool) Instruction {
	return func(io InstructionIO) {
		a := io.Load(RegisterA)
		b := src(io)
		c := uint8(0)
		if useBorrow && io.GetFlag(FlagC) {
			c = 1
		}
		r16 := uint16(a) - uint16(b) - uint16(c)
		r := uint8(r16)

		io.SetFlag(FlagZ, r == 0)
		io.SetFlag(FlagN, true)
		io.SetFlag(FlagH, (a&0xf)-(b&0xf)-(c&0xf) > 0xf)
		io.SetFlag(FlagC, r16 > 0xff)

		io.Store(RegisterA, r)
	}
}

// Generate an AND instruction.
func opAND(src OpSrc) Instruction {
	return func(io InstructionIO) {
		a := io.Load(RegisterA)
		b := src(io)
		r := a & b

		io.SetFlag(FlagZ, r == 0)
		io.SetFlag(FlagN, false)
		io.SetFlag(FlagH, true)
		io.SetFlag(FlagC, false)

		io.Store(RegisterA, r)
	}
}

// Generate an XOR instruction.
func opXOR(src OpSrc) Instruction {
	return func(io InstructionIO) {
		a := io.Load(RegisterA)
		b := src(io)
		r := a ^ b

		io.SetFlag(FlagZ, r == 0)
		io.SetFlag(FlagN, false)
		io.SetFlag(FlagH, false)
		io.SetFlag(FlagC, false)

		io.Store(RegisterA, r)
	}
}

// Generate an OR instruction.
func opOR(src OpSrc) Instruction {
	return func(io InstructionIO) {
		a := io.Load(RegisterA)
		b := src(io)
		r := a | b

		io.SetFlag(FlagZ, r == 0)
		io.SetFlag(FlagN, false)
		io.SetFlag(FlagH, false)
		io.SetFlag(FlagC, false)

		io.Store(RegisterA, r)
	}
}

// Generate a CP instruction.
func opCP(src OpSrc) Instruction {
	return func(io InstructionIO) {
		a := io.Load(RegisterA)
		b := src(io)
		r := a - b

		io.SetFlag(FlagZ, r == 0)
		io.SetFlag(FlagN, true)
		io.SetFlag(FlagH, a&0xf < b&0xf)
		io.SetFlag(FlagC, a < b)
	}
}

// Generate an INC instruction.
func opINC(dst OpDst, src OpSrc) Instruction {
	return func(io InstructionIO) {
		a := src(io)
		r := a + 1

		io.SetFlag(FlagZ, r == 0)
		io.SetFlag(FlagN, false)
		io.SetFlag(FlagH, a&0xf > 0xe)

		dst(io, r)
	}
}

// Generate a DEC instruction.
func opDEC(dst OpDst, src OpSrc) Instruction {
	return func(io InstructionIO) {
		a := src(io)
		r := a - 1

		io.SetFlag(FlagZ, r == 0)
		io.SetFlag(FlagN, true)
		io.SetFlag(FlagH, a&0xf == 0x0)

		dst(io, r)
	}
}

// Generate a 16-bit ADD instruction.
func opADD16(src OpSrc16) Instruction {
	return func(io InstructionIO) {
		a := io.Load16(RegisterHL)
		b := src(io)
		r32 := uint32(a) + uint32(b)

		io.SetFlag(FlagN, false)
		io.SetFlag(FlagH, uint32(a&0x0fff) > r32&0xfff)
		io.SetFlag(FlagC, r32 > 0xffff)

		io.Nop()
		io.Store16(RegisterHL, uint16(r32))
	}
}

// Generate a 16-bit INC instruction.
func opINC16(dst OpDst16, src OpSrc16) Instruction {
	return func(io InstructionIO) {
		a := src(io)
		io.Nop()
		dst(io, a+1)
	}
}

// Generate a 16-bit DEC instruction.
func opDEC16(dst OpDst16, src OpSrc16) Instruction {
	return func(io InstructionIO) {
		a := src(io)
		io.Nop()
		dst(io, a-1)
	}
}
