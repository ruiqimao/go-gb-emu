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
	Read    func(addr uint16) uint8
	Read16  func(addr uint16) uint16
	Write   func(addr uint16, v uint8)
	Write16 func(addr uint16, v uint16)

	// Immediate value access.
	PC      func() uint16
	SetPC   func(v uint16)
	PopPC   func() uint8
	PopPC16 func() uint16

	// Stack access.
	SP     func() uint16
	SetSP  func(v uint16)
	PopSP  func() uint16
	PushSP func(uint16)

	// Interrupt access.
	SetIME func(v bool)
	Halt   func()

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

		0xf9: opPAD(opLD16(opSetSP(), opLoad16(RegisterHL))),
		0xf8: opADDSP(),
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

		0xe8: opPAD(opPAD(opLD16(opSetSP(), opSAdd(opSP(), opImmediate())))),

		0x03: opINC16(opStore16(RegisterBC), opLoad16(RegisterBC)),
		0x13: opINC16(opStore16(RegisterDE), opLoad16(RegisterDE)),
		0x23: opINC16(opStore16(RegisterHL), opLoad16(RegisterHL)),
		0x33: opINC16(opSetSP(), opSP()),

		0x0b: opDEC16(opStore16(RegisterBC), opLoad16(RegisterBC)),
		0x1b: opDEC16(opStore16(RegisterDE), opLoad16(RegisterDE)),
		0x2b: opDEC16(opStore16(RegisterHL), opLoad16(RegisterHL)),
		0x3b: opDEC16(opSetSP(), opSP()),

		// Jumps.
		0xc3: opJP(opImmediate16(), opTrue()),
		0xc2: opJP(opImmediate16(), opNotFlag(FlagZ)),
		0xca: opJP(opImmediate16(), opFlag(FlagZ)),
		0xd2: opJP(opImmediate16(), opNotFlag(FlagC)),
		0xda: opJP(opImmediate16(), opFlag(FlagC)),

		0xe9: opJPHL(),

		0x18: opJR(opTrue()),
		0x20: opJR(opNotFlag(FlagZ)),
		0x28: opJR(opFlag(FlagZ)),
		0x30: opJR(opNotFlag(FlagC)),
		0x38: opJR(opFlag(FlagC)),

		0xcd: opCALL(opTrue()),
		0xc4: opCALL(opNotFlag(FlagZ)),
		0xcc: opCALL(opFlag(FlagZ)),
		0xd4: opCALL(opNotFlag(FlagC)),
		0xdc: opCALL(opFlag(FlagC)),

		0xc7: opRST(0x0000),
		0xcf: opRST(0x0008),
		0xd7: opRST(0x0010),
		0xdf: opRST(0x0018),
		0xe7: opRST(0x0020),
		0xef: opRST(0x0028),
		0xf7: opRST(0x0030),
		0xff: opRST(0x0038),

		0xc9: opRET(opTrue()),
		0xc0: opPAD(opRET(opNotFlag(FlagZ))),
		0xc8: opPAD(opRET(opFlag(FlagZ))),
		0xd0: opPAD(opRET(opNotFlag(FlagC))),
		0xd8: opPAD(opRET(opFlag(FlagC))),

		0xd9: opRETI(),

		// Miscellaneous.
		0x27: opDAA(),
		0x2f: opCPL(),
		0x3f: opCCF(),
		0x37: opSCF(),
		0x00: opNOP(),
		0x76: opHALT(),
		0x10: opSTOP(),
		0xf3: opDI(),
		0xfb: opEI(),

		// Bit operations.
		0x07: opRLCA(),
		0x17: opRLA(),
		0x0f: opRRCA(),
		0x1f: opRRA(),

		// CB extensions.
		0x100: opRL(opLoad(RegisterB), opStore(RegisterB), true),
		0x101: opRL(opLoad(RegisterC), opStore(RegisterC), true),
		0x102: opRL(opLoad(RegisterD), opStore(RegisterD), true),
		0x103: opRL(opLoad(RegisterE), opStore(RegisterE), true),
		0x104: opRL(opLoad(RegisterH), opStore(RegisterH), true),
		0x105: opRL(opLoad(RegisterL), opStore(RegisterL), true),
		0x106: opRL(opRead(opLoad16(RegisterHL)), opWrite(opLoad16(RegisterHL)), true),
		0x107: opRL(opLoad(RegisterA), opStore(RegisterA), true),

		0x108: opRR(opLoad(RegisterB), opStore(RegisterB), true),
		0x109: opRR(opLoad(RegisterC), opStore(RegisterC), true),
		0x10a: opRR(opLoad(RegisterD), opStore(RegisterD), true),
		0x10b: opRR(opLoad(RegisterE), opStore(RegisterE), true),
		0x10c: opRR(opLoad(RegisterH), opStore(RegisterH), true),
		0x10d: opRR(opLoad(RegisterL), opStore(RegisterL), true),
		0x10e: opRR(opRead(opLoad16(RegisterHL)), opWrite(opLoad16(RegisterHL)), true),
		0x10f: opRR(opLoad(RegisterA), opStore(RegisterA), true),

		0x110: opRL(opLoad(RegisterB), opStore(RegisterB), false),
		0x111: opRL(opLoad(RegisterC), opStore(RegisterC), false),
		0x112: opRL(opLoad(RegisterD), opStore(RegisterD), false),
		0x113: opRL(opLoad(RegisterE), opStore(RegisterE), false),
		0x114: opRL(opLoad(RegisterH), opStore(RegisterH), false),
		0x115: opRL(opLoad(RegisterL), opStore(RegisterL), false),
		0x116: opRL(opRead(opLoad16(RegisterHL)), opWrite(opLoad16(RegisterHL)), false),
		0x117: opRL(opLoad(RegisterA), opStore(RegisterA), false),

		0x118: opRR(opLoad(RegisterB), opStore(RegisterB), false),
		0x119: opRR(opLoad(RegisterC), opStore(RegisterC), false),
		0x11a: opRR(opLoad(RegisterD), opStore(RegisterD), false),
		0x11b: opRR(opLoad(RegisterE), opStore(RegisterE), false),
		0x11c: opRR(opLoad(RegisterH), opStore(RegisterH), false),
		0x11d: opRR(opLoad(RegisterL), opStore(RegisterL), false),
		0x11e: opRR(opRead(opLoad16(RegisterHL)), opWrite(opLoad16(RegisterHL)), false),
		0x11f: opRR(opLoad(RegisterA), opStore(RegisterA), false),

		0x120: opSLA(opLoad(RegisterB), opStore(RegisterB)),
		0x121: opSLA(opLoad(RegisterC), opStore(RegisterC)),
		0x122: opSLA(opLoad(RegisterD), opStore(RegisterD)),
		0x123: opSLA(opLoad(RegisterE), opStore(RegisterE)),
		0x124: opSLA(opLoad(RegisterH), opStore(RegisterH)),
		0x125: opSLA(opLoad(RegisterL), opStore(RegisterL)),
		0x126: opSLA(opRead(opLoad16(RegisterHL)), opWrite(opLoad16(RegisterHL))),
		0x127: opSLA(opLoad(RegisterA), opStore(RegisterA)),

		0x128: opSR(opLoad(RegisterB), opStore(RegisterB), true),
		0x129: opSR(opLoad(RegisterC), opStore(RegisterC), true),
		0x12a: opSR(opLoad(RegisterD), opStore(RegisterD), true),
		0x12b: opSR(opLoad(RegisterE), opStore(RegisterE), true),
		0x12c: opSR(opLoad(RegisterH), opStore(RegisterH), true),
		0x12d: opSR(opLoad(RegisterL), opStore(RegisterL), true),
		0x12e: opSR(opRead(opLoad16(RegisterHL)), opWrite(opLoad16(RegisterHL)), true),
		0x12f: opSR(opLoad(RegisterA), opStore(RegisterA), true),

		0x130: opSWAP(opLoad(RegisterB), opStore(RegisterB)),
		0x131: opSWAP(opLoad(RegisterC), opStore(RegisterC)),
		0x132: opSWAP(opLoad(RegisterD), opStore(RegisterD)),
		0x133: opSWAP(opLoad(RegisterE), opStore(RegisterE)),
		0x134: opSWAP(opLoad(RegisterH), opStore(RegisterH)),
		0x135: opSWAP(opLoad(RegisterL), opStore(RegisterL)),
		0x136: opSWAP(opRead(opLoad16(RegisterHL)), opWrite(opLoad16(RegisterHL))),
		0x137: opSWAP(opLoad(RegisterA), opStore(RegisterA)),

		0x138: opSR(opLoad(RegisterB), opStore(RegisterB), false),
		0x139: opSR(opLoad(RegisterC), opStore(RegisterC), false),
		0x13a: opSR(opLoad(RegisterD), opStore(RegisterD), false),
		0x13b: opSR(opLoad(RegisterE), opStore(RegisterE), false),
		0x13c: opSR(opLoad(RegisterH), opStore(RegisterH), false),
		0x13d: opSR(opLoad(RegisterL), opStore(RegisterL), false),
		0x13e: opSR(opRead(opLoad16(RegisterHL)), opWrite(opLoad16(RegisterHL)), false),
		0x13f: opSR(opLoad(RegisterA), opStore(RegisterA), false),

		0x140: opBIT(opLoad(RegisterB), 0),
		0x141: opBIT(opLoad(RegisterC), 0),
		0x142: opBIT(opLoad(RegisterD), 0),
		0x143: opBIT(opLoad(RegisterE), 0),
		0x144: opBIT(opLoad(RegisterH), 0),
		0x145: opBIT(opLoad(RegisterL), 0),
		0x146: opBIT(opRead(opLoad16(RegisterHL)), 0),
		0x147: opBIT(opLoad(RegisterA), 0),

		0x148: opBIT(opLoad(RegisterB), 1),
		0x149: opBIT(opLoad(RegisterC), 1),
		0x14a: opBIT(opLoad(RegisterD), 1),
		0x14b: opBIT(opLoad(RegisterE), 1),
		0x14c: opBIT(opLoad(RegisterH), 1),
		0x14d: opBIT(opLoad(RegisterL), 1),
		0x14e: opBIT(opRead(opLoad16(RegisterHL)), 1),
		0x14f: opBIT(opLoad(RegisterA), 1),

		0x150: opBIT(opLoad(RegisterB), 2),
		0x151: opBIT(opLoad(RegisterC), 2),
		0x152: opBIT(opLoad(RegisterD), 2),
		0x153: opBIT(opLoad(RegisterE), 2),
		0x154: opBIT(opLoad(RegisterH), 2),
		0x155: opBIT(opLoad(RegisterL), 2),
		0x156: opBIT(opRead(opLoad16(RegisterHL)), 2),
		0x157: opBIT(opLoad(RegisterA), 2),

		0x158: opBIT(opLoad(RegisterB), 3),
		0x159: opBIT(opLoad(RegisterC), 3),
		0x15a: opBIT(opLoad(RegisterD), 3),
		0x15b: opBIT(opLoad(RegisterE), 3),
		0x15c: opBIT(opLoad(RegisterH), 3),
		0x15d: opBIT(opLoad(RegisterL), 3),
		0x15e: opBIT(opRead(opLoad16(RegisterHL)), 3),
		0x15f: opBIT(opLoad(RegisterA), 3),

		0x160: opBIT(opLoad(RegisterB), 4),
		0x161: opBIT(opLoad(RegisterC), 4),
		0x162: opBIT(opLoad(RegisterD), 4),
		0x163: opBIT(opLoad(RegisterE), 4),
		0x164: opBIT(opLoad(RegisterH), 4),
		0x165: opBIT(opLoad(RegisterL), 4),
		0x166: opBIT(opRead(opLoad16(RegisterHL)), 4),
		0x167: opBIT(opLoad(RegisterA), 4),

		0x168: opBIT(opLoad(RegisterB), 5),
		0x169: opBIT(opLoad(RegisterC), 5),
		0x16a: opBIT(opLoad(RegisterD), 5),
		0x16b: opBIT(opLoad(RegisterE), 5),
		0x16c: opBIT(opLoad(RegisterH), 5),
		0x16d: opBIT(opLoad(RegisterL), 5),
		0x16e: opBIT(opRead(opLoad16(RegisterHL)), 5),
		0x16f: opBIT(opLoad(RegisterA), 5),

		0x170: opBIT(opLoad(RegisterB), 6),
		0x171: opBIT(opLoad(RegisterC), 6),
		0x172: opBIT(opLoad(RegisterD), 6),
		0x173: opBIT(opLoad(RegisterE), 6),
		0x174: opBIT(opLoad(RegisterH), 6),
		0x175: opBIT(opLoad(RegisterL), 6),
		0x176: opBIT(opRead(opLoad16(RegisterHL)), 6),
		0x177: opBIT(opLoad(RegisterA), 6),

		0x178: opBIT(opLoad(RegisterB), 7),
		0x179: opBIT(opLoad(RegisterC), 7),
		0x17a: opBIT(opLoad(RegisterD), 7),
		0x17b: opBIT(opLoad(RegisterE), 7),
		0x17c: opBIT(opLoad(RegisterH), 7),
		0x17d: opBIT(opLoad(RegisterL), 7),
		0x17e: opBIT(opRead(opLoad16(RegisterHL)), 7),
		0x17f: opBIT(opLoad(RegisterA), 7),

		0x180: opSET(opStore(RegisterB), opLoad(RegisterB), 0, false),
		0x181: opSET(opStore(RegisterC), opLoad(RegisterC), 0, false),
		0x182: opSET(opStore(RegisterD), opLoad(RegisterD), 0, false),
		0x183: opSET(opStore(RegisterE), opLoad(RegisterE), 0, false),
		0x184: opSET(opStore(RegisterH), opLoad(RegisterH), 0, false),
		0x185: opSET(opStore(RegisterL), opLoad(RegisterL), 0, false),
		0x186: opSET(opWrite(opLoad16(RegisterHL)), opRead(opLoad16(RegisterHL)), 0, false),
		0x187: opSET(opStore(RegisterA), opLoad(RegisterA), 0, false),

		0x188: opSET(opStore(RegisterB), opLoad(RegisterB), 1, false),
		0x189: opSET(opStore(RegisterC), opLoad(RegisterC), 1, false),
		0x18a: opSET(opStore(RegisterD), opLoad(RegisterD), 1, false),
		0x18b: opSET(opStore(RegisterE), opLoad(RegisterE), 1, false),
		0x18c: opSET(opStore(RegisterH), opLoad(RegisterH), 1, false),
		0x18d: opSET(opStore(RegisterL), opLoad(RegisterL), 1, false),
		0x18e: opSET(opWrite(opLoad16(RegisterHL)), opRead(opLoad16(RegisterHL)), 1, false),
		0x18f: opSET(opStore(RegisterA), opLoad(RegisterA), 1, false),

		0x190: opSET(opStore(RegisterB), opLoad(RegisterB), 2, false),
		0x191: opSET(opStore(RegisterC), opLoad(RegisterC), 2, false),
		0x192: opSET(opStore(RegisterD), opLoad(RegisterD), 2, false),
		0x193: opSET(opStore(RegisterE), opLoad(RegisterE), 2, false),
		0x194: opSET(opStore(RegisterH), opLoad(RegisterH), 2, false),
		0x195: opSET(opStore(RegisterL), opLoad(RegisterL), 2, false),
		0x196: opSET(opWrite(opLoad16(RegisterHL)), opRead(opLoad16(RegisterHL)), 2, false),
		0x197: opSET(opStore(RegisterA), opLoad(RegisterA), 2, false),

		0x198: opSET(opStore(RegisterB), opLoad(RegisterB), 3, false),
		0x199: opSET(opStore(RegisterC), opLoad(RegisterC), 3, false),
		0x19a: opSET(opStore(RegisterD), opLoad(RegisterD), 3, false),
		0x19b: opSET(opStore(RegisterE), opLoad(RegisterE), 3, false),
		0x19c: opSET(opStore(RegisterH), opLoad(RegisterH), 3, false),
		0x19d: opSET(opStore(RegisterL), opLoad(RegisterL), 3, false),
		0x19e: opSET(opWrite(opLoad16(RegisterHL)), opRead(opLoad16(RegisterHL)), 3, false),
		0x19f: opSET(opStore(RegisterA), opLoad(RegisterA), 3, false),

		0x1a0: opSET(opStore(RegisterB), opLoad(RegisterB), 4, false),
		0x1a1: opSET(opStore(RegisterC), opLoad(RegisterC), 4, false),
		0x1a2: opSET(opStore(RegisterD), opLoad(RegisterD), 4, false),
		0x1a3: opSET(opStore(RegisterE), opLoad(RegisterE), 4, false),
		0x1a4: opSET(opStore(RegisterH), opLoad(RegisterH), 4, false),
		0x1a5: opSET(opStore(RegisterL), opLoad(RegisterL), 4, false),
		0x1a6: opSET(opWrite(opLoad16(RegisterHL)), opRead(opLoad16(RegisterHL)), 4, false),
		0x1a7: opSET(opStore(RegisterA), opLoad(RegisterA), 4, false),

		0x1a8: opSET(opStore(RegisterB), opLoad(RegisterB), 5, false),
		0x1a9: opSET(opStore(RegisterC), opLoad(RegisterC), 5, false),
		0x1aa: opSET(opStore(RegisterD), opLoad(RegisterD), 5, false),
		0x1ab: opSET(opStore(RegisterE), opLoad(RegisterE), 5, false),
		0x1ac: opSET(opStore(RegisterH), opLoad(RegisterH), 5, false),
		0x1ad: opSET(opStore(RegisterL), opLoad(RegisterL), 5, false),
		0x1ae: opSET(opWrite(opLoad16(RegisterHL)), opRead(opLoad16(RegisterHL)), 5, false),
		0x1af: opSET(opStore(RegisterA), opLoad(RegisterA), 5, false),

		0x1b0: opSET(opStore(RegisterB), opLoad(RegisterB), 6, false),
		0x1b1: opSET(opStore(RegisterC), opLoad(RegisterC), 6, false),
		0x1b2: opSET(opStore(RegisterD), opLoad(RegisterD), 6, false),
		0x1b3: opSET(opStore(RegisterE), opLoad(RegisterE), 6, false),
		0x1b4: opSET(opStore(RegisterH), opLoad(RegisterH), 6, false),
		0x1b5: opSET(opStore(RegisterL), opLoad(RegisterL), 6, false),
		0x1b6: opSET(opWrite(opLoad16(RegisterHL)), opRead(opLoad16(RegisterHL)), 6, false),
		0x1b7: opSET(opStore(RegisterA), opLoad(RegisterA), 6, false),

		0x1b8: opSET(opStore(RegisterB), opLoad(RegisterB), 7, false),
		0x1b9: opSET(opStore(RegisterC), opLoad(RegisterC), 7, false),
		0x1ba: opSET(opStore(RegisterD), opLoad(RegisterD), 7, false),
		0x1bb: opSET(opStore(RegisterE), opLoad(RegisterE), 7, false),
		0x1bc: opSET(opStore(RegisterH), opLoad(RegisterH), 7, false),
		0x1bd: opSET(opStore(RegisterL), opLoad(RegisterL), 7, false),
		0x1be: opSET(opWrite(opLoad16(RegisterHL)), opRead(opLoad16(RegisterHL)), 7, false),
		0x1bf: opSET(opStore(RegisterA), opLoad(RegisterA), 7, false),

		0x1c0: opSET(opStore(RegisterB), opLoad(RegisterB), 0, true),
		0x1c1: opSET(opStore(RegisterC), opLoad(RegisterC), 0, true),
		0x1c2: opSET(opStore(RegisterD), opLoad(RegisterD), 0, true),
		0x1c3: opSET(opStore(RegisterE), opLoad(RegisterE), 0, true),
		0x1c4: opSET(opStore(RegisterH), opLoad(RegisterH), 0, true),
		0x1c5: opSET(opStore(RegisterL), opLoad(RegisterL), 0, true),
		0x1c6: opSET(opWrite(opLoad16(RegisterHL)), opRead(opLoad16(RegisterHL)), 0, true),
		0x1c7: opSET(opStore(RegisterA), opLoad(RegisterA), 0, true),

		0x1c8: opSET(opStore(RegisterB), opLoad(RegisterB), 1, true),
		0x1c9: opSET(opStore(RegisterC), opLoad(RegisterC), 1, true),
		0x1ca: opSET(opStore(RegisterD), opLoad(RegisterD), 1, true),
		0x1cb: opSET(opStore(RegisterE), opLoad(RegisterE), 1, true),
		0x1cc: opSET(opStore(RegisterH), opLoad(RegisterH), 1, true),
		0x1cd: opSET(opStore(RegisterL), opLoad(RegisterL), 1, true),
		0x1ce: opSET(opWrite(opLoad16(RegisterHL)), opRead(opLoad16(RegisterHL)), 1, true),
		0x1cf: opSET(opStore(RegisterA), opLoad(RegisterA), 1, true),

		0x1d0: opSET(opStore(RegisterB), opLoad(RegisterB), 2, true),
		0x1d1: opSET(opStore(RegisterC), opLoad(RegisterC), 2, true),
		0x1d2: opSET(opStore(RegisterD), opLoad(RegisterD), 2, true),
		0x1d3: opSET(opStore(RegisterE), opLoad(RegisterE), 2, true),
		0x1d4: opSET(opStore(RegisterH), opLoad(RegisterH), 2, true),
		0x1d5: opSET(opStore(RegisterL), opLoad(RegisterL), 2, true),
		0x1d6: opSET(opWrite(opLoad16(RegisterHL)), opRead(opLoad16(RegisterHL)), 2, true),
		0x1d7: opSET(opStore(RegisterA), opLoad(RegisterA), 2, true),

		0x1d8: opSET(opStore(RegisterB), opLoad(RegisterB), 3, true),
		0x1d9: opSET(opStore(RegisterC), opLoad(RegisterC), 3, true),
		0x1da: opSET(opStore(RegisterD), opLoad(RegisterD), 3, true),
		0x1db: opSET(opStore(RegisterE), opLoad(RegisterE), 3, true),
		0x1dc: opSET(opStore(RegisterH), opLoad(RegisterH), 3, true),
		0x1dd: opSET(opStore(RegisterL), opLoad(RegisterL), 3, true),
		0x1de: opSET(opWrite(opLoad16(RegisterHL)), opRead(opLoad16(RegisterHL)), 3, true),
		0x1df: opSET(opStore(RegisterA), opLoad(RegisterA), 3, true),

		0x1e0: opSET(opStore(RegisterB), opLoad(RegisterB), 4, true),
		0x1e1: opSET(opStore(RegisterC), opLoad(RegisterC), 4, true),
		0x1e2: opSET(opStore(RegisterD), opLoad(RegisterD), 4, true),
		0x1e3: opSET(opStore(RegisterE), opLoad(RegisterE), 4, true),
		0x1e4: opSET(opStore(RegisterH), opLoad(RegisterH), 4, true),
		0x1e5: opSET(opStore(RegisterL), opLoad(RegisterL), 4, true),
		0x1e6: opSET(opWrite(opLoad16(RegisterHL)), opRead(opLoad16(RegisterHL)), 4, true),
		0x1e7: opSET(opStore(RegisterA), opLoad(RegisterA), 4, true),

		0x1e8: opSET(opStore(RegisterB), opLoad(RegisterB), 5, true),
		0x1e9: opSET(opStore(RegisterC), opLoad(RegisterC), 5, true),
		0x1ea: opSET(opStore(RegisterD), opLoad(RegisterD), 5, true),
		0x1eb: opSET(opStore(RegisterE), opLoad(RegisterE), 5, true),
		0x1ec: opSET(opStore(RegisterH), opLoad(RegisterH), 5, true),
		0x1ed: opSET(opStore(RegisterL), opLoad(RegisterL), 5, true),
		0x1ee: opSET(opWrite(opLoad16(RegisterHL)), opRead(opLoad16(RegisterHL)), 5, true),
		0x1ef: opSET(opStore(RegisterA), opLoad(RegisterA), 5, true),

		0x1f0: opSET(opStore(RegisterB), opLoad(RegisterB), 6, true),
		0x1f1: opSET(opStore(RegisterC), opLoad(RegisterC), 6, true),
		0x1f2: opSET(opStore(RegisterD), opLoad(RegisterD), 6, true),
		0x1f3: opSET(opStore(RegisterE), opLoad(RegisterE), 6, true),
		0x1f4: opSET(opStore(RegisterH), opLoad(RegisterH), 6, true),
		0x1f5: opSET(opStore(RegisterL), opLoad(RegisterL), 6, true),
		0x1f6: opSET(opWrite(opLoad16(RegisterHL)), opRead(opLoad16(RegisterHL)), 6, true),
		0x1f7: opSET(opStore(RegisterA), opLoad(RegisterA), 6, true),

		0x1f8: opSET(opStore(RegisterB), opLoad(RegisterB), 7, true),
		0x1f9: opSET(opStore(RegisterC), opLoad(RegisterC), 7, true),
		0x1fa: opSET(opStore(RegisterD), opLoad(RegisterD), 7, true),
		0x1fb: opSET(opStore(RegisterE), opLoad(RegisterE), 7, true),
		0x1fc: opSET(opStore(RegisterH), opLoad(RegisterH), 7, true),
		0x1fd: opSET(opStore(RegisterL), opLoad(RegisterL), 7, true),
		0x1fe: opSET(opWrite(opLoad16(RegisterHL)), opRead(opLoad16(RegisterHL)), 7, true),
		0x1ff: opSET(opStore(RegisterA), opLoad(RegisterA), 7, true),
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
		dst(io, io.PopSP())
	}
}

// Generate an ADD SP,r8 instruction.
func opADDSP() Instruction {
	return func(io InstructionIO) {
		src := opSAdd(opSP(), opImmediate())
		dst := opStore16(RegisterHL)
		r := src(io)
		io.Nop()
		dst(io, r)
	}
}

// Generate a PUSH instruction.
func opPUSH(src OpSrc16) Instruction {
	return func(io InstructionIO) {
		io.Nop()
		io.PushSP(src(io))
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

// Generate a JP instruction.
func opJP(src OpSrc16, flag OpFlagSrc) Instruction {
	return func(io InstructionIO) {
		f := flag(io)
		a := src(io)
		if f {
			io.Nop()
			io.SetPC(a)
		}
	}
}

// Generate a JP,HL instruction.
func opJPHL() Instruction {
	return func(io InstructionIO) {
		io.SetPC(io.Load16(RegisterHL))
	}
}

// Generate a JR instruction.
func opJR(flag OpFlagSrc) Instruction {
	return func(io InstructionIO) {
		f := flag(io)
		r := uint16(int8(io.PopPC()))
		target := io.PC() + r
		if f {
			io.Nop()
			io.SetPC(target)
		}
	}
}

// Generate a CALL instruction.
func opCALL(flag OpFlagSrc) Instruction {
	return func(io InstructionIO) {
		a := opImmediate16()(io)
		if flag(io) {
			io.Nop()
			io.PushSP(io.PC())
			io.SetPC(a)
		}
	}
}

// Generate a RST instruction.
func opRST(addr uint16) Instruction {
	return func(io InstructionIO) {
		io.Nop()
		io.PushSP(io.PC())
		io.SetPC(addr)
	}
}

// Generate a RET instruction.
func opRET(flag OpFlagSrc) Instruction {
	return func(io InstructionIO) {
		if flag(io) {
			io.Nop()
			io.SetPC(io.PopSP())
		}
	}
}

// Generate a RETI instruction.
func opRETI() Instruction {
	return func(io InstructionIO) {
		io.Nop()
		io.SetPC(io.PopSP())
		io.SetIME(true)
	}
}

// Generate a DAA instruction.
func opDAA() Instruction {
	return func(io InstructionIO) {
		a := io.Load(RegisterA)

		// Stolen from https://forums.nesdev.com/viewtopic.php?t=15944#p196282.
		if !io.GetFlag(FlagN) {
			if io.GetFlag(FlagC) || a > 0x99 {
				a += 0x60
				io.SetFlag(FlagC, true)
			}
			if io.GetFlag(FlagH) || (a&0xf) > 0x09 {
				a += 0x06
			}
		} else {
			if io.GetFlag(FlagC) {
				a -= 0x60
				io.SetFlag(FlagC, true)
			}
			if io.GetFlag(FlagH) {
				a -= 0x06
			}
		}

		io.SetFlag(FlagZ, a == 0)
		io.SetFlag(FlagH, false)

		io.Store(RegisterA, a)
	}
}

// Generate a CPL instruction.
func opCPL() Instruction {
	return func(io InstructionIO) {
		io.Store(RegisterA, ^io.Load(RegisterA))
		io.SetFlag(FlagN, true)
		io.SetFlag(FlagH, true)
	}
}

// Generate a CCF instruction.
func opCCF() Instruction {
	return func(io InstructionIO) {
		io.SetFlag(FlagN, false)
		io.SetFlag(FlagH, false)
		io.SetFlag(FlagC, !io.GetFlag(FlagC))
	}
}

// Generate a SCF instruction.
func opSCF() Instruction {
	return func(io InstructionIO) {
		io.SetFlag(FlagN, false)
		io.SetFlag(FlagH, false)
		io.SetFlag(FlagC, true)
	}
}

// Generate a NOP instruction.
func opNOP() Instruction {
	return func(io InstructionIO) {
		// Do nothing.
	}
}

// Generate a HALT instruction.
func opHALT() Instruction {
	return func(io InstructionIO) {
		io.Halt()
	}
}

// Generate a STOP instruction.
func opSTOP() Instruction {
	return func(io InstructionIO) {
		// TODO: Properly implement.
		io.Halt()
		io.SetPC(io.PC() + 1)
	}
}

// Generate a DI instruction.
func opDI() Instruction {
	return func(io InstructionIO) {
		io.SetIME(false)
	}
}

// Generate a EI instruction.
func opEI() Instruction {
	return func(io InstructionIO) {
		io.SetIME(true)
	}
}

// Generate an RLCA instruction.
func opRLCA() Instruction {
	return func(io InstructionIO) {
		a := io.Load(RegisterA)
		r := a<<1 | a>>7

		io.SetFlag(FlagZ, false)
		io.SetFlag(FlagN, false)
		io.SetFlag(FlagH, false)
		io.SetFlag(FlagC, a>>7 == 0x1)

		io.Store(RegisterA, r)
	}
}

// Generate an RLA instruction.
func opRLA() Instruction {
	return func(io InstructionIO) {
		a := io.Load(RegisterA)
		r := a << 1
		if io.GetFlag(FlagC) {
			r |= 0x1
		}

		io.SetFlag(FlagZ, false)
		io.SetFlag(FlagN, false)
		io.SetFlag(FlagH, false)
		io.SetFlag(FlagC, a>>7 == 0x1)

		io.Store(RegisterA, r)
	}
}

// Generate an RRCA instruction.
func opRRCA() Instruction {
	return func(io InstructionIO) {
		a := io.Load(RegisterA)
		r := a>>1 | (a&0x1)<<7

		io.SetFlag(FlagZ, false)
		io.SetFlag(FlagN, false)
		io.SetFlag(FlagH, false)
		io.SetFlag(FlagC, a&0x1 == 0x1)

		io.Store(RegisterA, r)
	}
}

// Generate an RRA instruction.
func opRRA() Instruction {
	return func(io InstructionIO) {
		a := io.Load(RegisterA)
		r := a >> 1
		if io.GetFlag(FlagC) {
			r |= 0x80
		}

		io.SetFlag(FlagZ, false)
		io.SetFlag(FlagN, false)
		io.SetFlag(FlagH, false)
		io.SetFlag(FlagC, a&0x1 == 0x1)

		io.Store(RegisterA, r)
	}
}

// Generate an RL/RLC instruction.
func opRL(src OpSrc, dst OpDst, c bool) Instruction {
	return func(io InstructionIO) {
		a := src(io)
		r := a << 1
		if c {
			r |= a >> 7
		} else if io.GetFlag(FlagC) {
			r |= 0x1
		}

		io.SetFlag(FlagZ, r == 0)
		io.SetFlag(FlagN, false)
		io.SetFlag(FlagH, false)
		io.SetFlag(FlagC, a>>7 == 0x1)

		dst(io, r)
	}
}

// Generate an RR/RRC instruction.
func opRR(src OpSrc, dst OpDst, c bool) Instruction {
	return func(io InstructionIO) {
		a := src(io)
		r := a >> 1
		if c {
			r |= (a & 0x1) << 7
		} else if io.GetFlag(FlagC) {
			r |= 0x80
		}

		io.SetFlag(FlagZ, r == 0)
		io.SetFlag(FlagN, false)
		io.SetFlag(FlagH, false)
		io.SetFlag(FlagC, a&0x1 == 0x1)

		dst(io, r)
	}
}

// Generate an SLA instruction.
func opSLA(src OpSrc, dst OpDst) Instruction {
	return func(io InstructionIO) {
		a := src(io)
		r := a << 1

		io.SetFlag(FlagZ, r == 0)
		io.SetFlag(FlagN, false)
		io.SetFlag(FlagH, false)
		io.SetFlag(FlagC, a>>7 == 0x1)

		dst(io, r)
	}
}

// Generate an SRA/SRL instruction.
func opSR(src OpSrc, dst OpDst, keepMSB bool) Instruction {
	return func(io InstructionIO) {
		a := src(io)
		r := a >> 1
		if keepMSB {
			r = utils.SetBit(r, 7, utils.GetBit(a, 7))
		}

		io.SetFlag(FlagZ, r == 0)
		io.SetFlag(FlagN, false)
		io.SetFlag(FlagH, false)
		io.SetFlag(FlagC, a&0x1 == 0x1)

		dst(io, r)
	}
}

// Generate a SWAP instruction.
func opSWAP(src OpSrc, dst OpDst) Instruction {
	return func(io InstructionIO) {
		a := src(io)
		r := (a&0x0f)<<4 | (a&0xf0)>>4

		io.SetFlag(FlagZ, r == 0)
		io.SetFlag(FlagN, false)
		io.SetFlag(FlagH, false)
		io.SetFlag(FlagC, false)

		dst(io, r)
	}
}

// Generate a BIT instruction.
func opBIT(src OpSrc, bit int) Instruction {
	return func(io InstructionIO) {
		a := src(io)

		io.SetFlag(FlagZ, !utils.GetBit(a, bit))
		io.SetFlag(FlagN, false)
		io.SetFlag(FlagH, true)
	}
}

// Generate a SET/RST instruction.
func opSET(dst OpDst, src OpSrc, bit int, set bool) Instruction {
	return func(io InstructionIO) {
		a := src(io)
		r := utils.SetBit(a, bit, set)
		dst(io, r)
	}
}

// Pad an instruction with an extra no-op at the beginning.
func opPAD(inst Instruction) Instruction {
	return func(io InstructionIO) {
		io.Nop()
		inst(io)
	}
}
