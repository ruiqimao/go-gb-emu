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
		0xf8: opLD16(opStore16(RegisterHL), opSADD(opSP(), opLoad(RegisterHL))),
		0x08: opLD16(opWrite16(opImmediate16()), opSP()),

		0xc5: opPUSH(opLoad16(RegisterBC)),
		0xd5: opPUSH(opLoad16(RegisterDE)),
		0xe5: opPUSH(opLoad16(RegisterHL)),
		0xf5: opPUSH(opLoad16(RegisterAF)),

		0xc1: opPOP(opStore16(RegisterBC)),
		0xd1: opPOP(opStore16(RegisterDE)),
		0xe1: opPOP(opStore16(RegisterHL)),
		0xf1: opPOP(opStore16(RegisterAF)),
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

// Generate a signed add instruction.
func opSADD(srcA OpSrc16, srcB OpSrc) OpSrc16 {
	return func(io InstructionIO) uint16 {
		a := srcA(io)
		b := srcB(io)
		r := int32(a) + int32(int8(b))

		io.SetFlag(FlagZ, false)
		io.SetFlag(FlagN, false)
		io.SetFlag(FlagH, (uint8(a)&0xf)-(b&0xf) > 0xf)
		io.SetFlag(FlagC, r > 0xff)

		io.Nop()
		return uint16(r)
	}
}
