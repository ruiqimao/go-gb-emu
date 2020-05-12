package gb

import (
	"github.com/ruiqimao/go-gb-emu/utils"
)

// An Instruction returns how many cycles it takes to execute.
type Instruction func() int

func (c *Cpu) CreateInstructionSet() {
	cpu := c
	mem := c.gb.mem

	c.instructions = [0x200]Instruction{
		// 8 bit loads.
		0x02: func() int { // LD (BC),A.
			mem.Write(cpu.BC(), cpu.A())
			return 8
		},
		0x06: func() int { // LD B,d8.
			cpu.SetB(cpu.IncPC())
			return 8
		},
		0x0a: func() int { // LD A,(BC).
			cpu.SetA(mem.Read(cpu.BC()))
			return 8
		},
		0x0e: func() int { // LD C,d8.
			cpu.SetC(cpu.IncPC())
			return 8
		},
		0x12: func() int { // LD (DE),A.
			mem.Write(cpu.DE(), cpu.A())
			return 8
		},
		0x16: func() int { // LD D,d8.
			cpu.SetD(cpu.IncPC())
			return 8
		},
		0x1a: func() int { // LD A,(DE).
			cpu.SetA(mem.Read(cpu.DE()))
			return 8
		},
		0x1e: func() int { // LD E,d8.
			cpu.SetE(cpu.IncPC())
			return 8
		},
		0x22: func() int { // LD (HL+),A.
			hl := cpu.HL()
			mem.Write(hl, cpu.A())
			cpu.SetHL(hl + 1)
			return 8
		},
		0x26: func() int { // LD H,d8.
			cpu.SetH(cpu.IncPC())
			return 8
		},
		0x2a: func() int { // LD A,(HL+).
			hl := cpu.HL()
			cpu.SetA(mem.Read(hl))
			cpu.SetHL(hl + 1)
			return 8
		},
		0x2e: func() int { // LD L,d8.
			cpu.SetL(cpu.IncPC())
			return 8
		},
		0x32: func() int { // LD (HL-),A.
			hl := cpu.HL()
			mem.Write(hl, cpu.A())
			cpu.SetHL(hl - 1)
			return 8
		},
		0x36: func() int { // LD (HL),d8.
			mem.Write(cpu.HL(), cpu.IncPC())
			return 12
		},
		0x3a: func() int { // LD A,(HL-).
			hl := cpu.HL()
			cpu.SetA(mem.Read(hl))
			cpu.SetHL(hl - 1)
			return 8
		},
		0x3e: func() int { // LD A,d8.
			cpu.SetA(cpu.IncPC())
			return 8
		},
		0x40: func() int { // LD B,B.
			cpu.SetB(cpu.B())
			return 4
		},
		0x41: func() int { // LD B,C.
			cpu.SetB(cpu.C())
			return 4
		},
		0x42: func() int { // LD B,D.
			cpu.SetB(cpu.D())
			return 4
		},
		0x43: func() int { // LD B,E.
			cpu.SetB(cpu.E())
			return 4
		},
		0x44: func() int { // LD B,H.
			cpu.SetB(cpu.H())
			return 4
		},
		0x45: func() int { // LD B,L.
			cpu.SetB(cpu.L())
			return 4
		},
		0x46: func() int { // LD B,(HL).
			cpu.SetB(mem.Read(cpu.HL()))
			return 8
		},
		0x47: func() int { // LD B,A.
			cpu.SetB(cpu.A())
			return 4
		},
		0x48: func() int { // LD C,B.
			cpu.SetC(cpu.B())
			return 4
		},
		0x49: func() int { // LD C,C.
			cpu.SetC(cpu.C())
			return 4
		},
		0x4a: func() int { // LD C,D.
			cpu.SetC(cpu.D())
			return 4
		},
		0x4b: func() int { // LD C,E.
			cpu.SetC(cpu.E())
			return 4
		},
		0x4c: func() int { // LD C,H.
			cpu.SetC(cpu.H())
			return 4
		},
		0x4d: func() int { // LD C,L.
			cpu.SetC(cpu.L())
			return 4
		},
		0x4e: func() int { // LD C,(HL).
			cpu.SetC(mem.Read(cpu.HL()))
			return 8
		},
		0x4f: func() int { // LD C,A.
			cpu.SetC(cpu.A())
			return 4
		},
		0x50: func() int { // LD D,B.
			cpu.SetD(cpu.B())
			return 4
		},
		0x51: func() int { // LD D,C.
			cpu.SetD(cpu.C())
			return 4
		},
		0x52: func() int { // LD D,D.
			cpu.SetD(cpu.D())
			return 4
		},
		0x53: func() int { // LD D,E.
			cpu.SetD(cpu.E())
			return 4
		},
		0x54: func() int { // LD D,H.
			cpu.SetD(cpu.H())
			return 4
		},
		0x55: func() int { // LD D,L.
			cpu.SetD(cpu.L())
			return 4
		},
		0x56: func() int { // LD D,(HL).
			cpu.SetD(mem.Read(cpu.HL()))
			return 8
		},
		0x57: func() int { // LD D,A.
			cpu.SetD(cpu.A())
			return 4
		},
		0x58: func() int { // LD E,B.
			cpu.SetE(cpu.B())
			return 4
		},
		0x59: func() int { // LD E,C.
			cpu.SetE(cpu.C())
			return 4
		},
		0x5a: func() int { // LD E,D.
			cpu.SetE(cpu.D())
			return 4
		},
		0x5b: func() int { // LD E,E.
			cpu.SetE(cpu.E())
			return 4
		},
		0x5c: func() int { // LD E,H.
			cpu.SetE(cpu.H())
			return 4
		},
		0x5d: func() int { // LD E,L.
			cpu.SetE(cpu.L())
			return 4
		},
		0x5e: func() int { // LD E,(HL).
			cpu.SetE(mem.Read(cpu.HL()))
			return 8
		},
		0x5f: func() int { // LD E,A.
			cpu.SetE(cpu.A())
			return 4
		},
		0x60: func() int { // LD H,B.
			cpu.SetH(cpu.B())
			return 4
		},
		0x61: func() int { // LD H,C.
			cpu.SetH(cpu.C())
			return 4
		},
		0x62: func() int { // LD H,D.
			cpu.SetH(cpu.D())
			return 4
		},
		0x63: func() int { // LD H,E.
			cpu.SetH(cpu.E())
			return 4
		},
		0x64: func() int { // LD H,H.
			cpu.SetH(cpu.H())
			return 4
		},
		0x65: func() int { // LD H,L.
			cpu.SetH(cpu.L())
			return 4
		},
		0x66: func() int { // LD H,(HL).
			cpu.SetH(mem.Read(cpu.HL()))
			return 8
		},
		0x67: func() int { // LD H,A.
			cpu.SetH(cpu.A())
			return 4
		},
		0x68: func() int { // LD L,B.
			cpu.SetL(cpu.B())
			return 4
		},
		0x69: func() int { // LD L,C.
			cpu.SetL(cpu.C())
			return 4
		},
		0x6a: func() int { // LD L,D.
			cpu.SetL(cpu.D())
			return 4
		},
		0x6b: func() int { // LD L,E.
			cpu.SetL(cpu.E())
			return 4
		},
		0x6c: func() int { // LD L,H.
			cpu.SetL(cpu.H())
			return 4
		},
		0x6d: func() int { // LD L,L.
			cpu.SetL(cpu.L())
			return 4
		},
		0x6e: func() int { // LD L,(HL).
			cpu.SetL(mem.Read(cpu.HL()))
			return 8
		},
		0x6f: func() int { // LD L,A.
			cpu.SetL(cpu.A())
			return 4
		},
		0x70: func() int { // LD (HL),B.
			mem.Write(cpu.HL(), cpu.B())
			return 8
		},
		0x71: func() int { // LD (HL),C.
			mem.Write(cpu.HL(), cpu.C())
			return 8
		},
		0x72: func() int { // LD (HL),D.
			mem.Write(cpu.HL(), cpu.D())
			return 8
		},
		0x73: func() int { // LD (HL),E.
			mem.Write(cpu.HL(), cpu.E())
			return 8
		},
		0x74: func() int { // LD (HL),H.
			mem.Write(cpu.HL(), cpu.H())
			return 8
		},
		0x75: func() int { // LD (HL),L.
			mem.Write(cpu.HL(), cpu.L())
			return 8
		},
		0x77: func() int { // LD (HL),A.
			mem.Write(cpu.HL(), cpu.A())
			return 8
		},
		0x78: func() int { // LD A,B.
			cpu.SetA(cpu.B())
			return 4
		},
		0x79: func() int { // LD A,C.
			cpu.SetA(cpu.C())
			return 4
		},
		0x7a: func() int { // LD A,D.
			cpu.SetA(cpu.D())
			return 4
		},
		0x7b: func() int { // LD A,E.
			cpu.SetA(cpu.E())
			return 4
		},
		0x7c: func() int { // LD A,H.
			cpu.SetA(cpu.H())
			return 4
		},
		0x7d: func() int { // LD A,L.
			cpu.SetA(cpu.L())
			return 4
		},
		0x7e: func() int { // LD A,(HL).
			cpu.SetA(mem.Read(cpu.HL()))
			return 8
		},
		0x7f: func() int { // LD A,A.
			cpu.SetA(cpu.A())
			return 4
		},
		0xe0: func() int { // LD ($FF00+a8),A.
			mem.Write(0xff00+uint16(cpu.IncPC()), cpu.A())
			return 12
		},
		0xe2: func() int { // LD ($FF00+C),A.
			mem.Write(0xff00+uint16(cpu.C()), cpu.A())
			return 8
		},
		0xea: func() int { // LD (a16),A.
			mem.Write(cpu.IncPC16(), cpu.A())
			return 16
		},
		0xf0: func() int { // LD A,($FF00+a8).
			cpu.SetA(mem.Read(0xff00 + uint16(cpu.IncPC())))
			return 12
		},
		0xf2: func() int { // LD A,(C).
			cpu.SetA(mem.Read(0xff00 + uint16(cpu.C())))
			return 8
		},
		0xfa: func() int { // LD A,(a16).
			cpu.SetA(mem.Read(cpu.IncPC16()))
			return 16
		},

		// 16 bit loads.
		0x01: func() int { // LD BC,d16.
			cpu.SetBC(cpu.IncPC16())
			return 12
		},
		0x08: func() int { // LD (a16),SP.
			mem.Write16(cpu.IncPC16(), cpu.SP())
			return 20
		},
		0x11: func() int { // LD DE,d16.
			cpu.SetDE(cpu.IncPC16())
			return 12
		},
		0x21: func() int { // LD HL,d16.
			cpu.SetHL(cpu.IncPC16())
			return 12
		},
		0x31: func() int { // LD SP,d16.
			cpu.SetSP(cpu.IncPC16())
			return 12
		},
		0xc1: func() int { // POP BC.
			cpu.SetBC(cpu.PopSP())
			return 12
		},
		0xc5: func() int { // PUSH BC.
			cpu.PushSP(cpu.BC())
			return 16
		},
		0xd1: func() int { // POP DE.
			cpu.SetDE(cpu.PopSP())
			return 12
		},
		0xd5: func() int { // PUSH DE.
			cpu.PushSP(cpu.DE())
			return 16
		},
		0xe1: func() int { // POP HL.
			cpu.SetHL(cpu.PopSP())
			return 12
		},
		0xe5: func() int { // PUSH HL.
			cpu.PushSP(cpu.HL())
			return 16
		},
		0xf1: func() int { // POP AF.
			cpu.SetAF(cpu.PopSP())
			return 12
		},
		0xf5: func() int { // PUSH AF.
			cpu.PushSP(cpu.AF())
			return 16
		},
		0xf8: func() int { // LD HL,SP+r8.
			cpu.SetHL(cpu.opSignedAdd(cpu.SP(), cpu.IncPC()))
			return 12
		},
		0xf9: func() int { // LD SP,HL.
			cpu.SetSP(cpu.HL())
			return 8
		},

		// 8 bit arithmetic.
		0x04: func() int { // INC B.
			cpu.SetB(cpu.opInc(cpu.B()))
			return 4
		},
		0x05: func() int { // DEC B.
			cpu.SetB(cpu.opDec(cpu.B()))
			return 4
		},
		0x0c: func() int { // INC C.
			cpu.SetC(cpu.opInc(cpu.C()))
			return 4
		},
		0x0d: func() int { // DEC C.
			cpu.SetC(cpu.opDec(cpu.C()))
			return 4
		},
		0x14: func() int { // INC D.
			cpu.SetD(cpu.opInc(cpu.D()))
			return 4
		},
		0x15: func() int { // DEC D.
			cpu.SetD(cpu.opDec(cpu.D()))
			return 4
		},
		0x1c: func() int { // INC E.
			cpu.SetE(cpu.opInc(cpu.E()))
			return 4
		},
		0x1d: func() int { // DEC E.
			cpu.SetE(cpu.opDec(cpu.E()))
			return 4
		},
		0x24: func() int { // INC H.
			cpu.SetH(cpu.opInc(cpu.H()))
			return 4
		},
		0x25: func() int { // DEC H.
			cpu.SetH(cpu.opDec(cpu.H()))
			return 4
		},
		0x27: func() int { // DAA.
			cpu.SetA(cpu.opDaa())
			return 4
		},
		0x2c: func() int { // INC L.
			cpu.SetL(cpu.opInc(cpu.L()))
			return 4
		},
		0x2d: func() int { // DEC L.
			cpu.SetL(cpu.opDec(cpu.L()))
			return 4
		},
		0x2f: func() int { // CPL.
			cpu.SetA(^cpu.A())
			cpu.SetFlagN(true)
			cpu.SetFlagH(true)
			return 4
		},
		0x34: func() int { // INC (HL).
			hl := cpu.HL()
			mem.Write(hl, cpu.opInc(mem.Read(hl)))
			return 12
		},
		0x35: func() int { // DEC (HL).
			hl := cpu.HL()
			mem.Write(hl, cpu.opDec(mem.Read(hl)))
			return 12
		},
		0x37: func() int { // SCF.
			cpu.SetFlagN(false)
			cpu.SetFlagH(false)
			cpu.SetFlagC(true)
			return 4
		},
		0x3c: func() int { // INC A.
			cpu.SetA(cpu.opInc(cpu.A()))
			return 4
		},
		0x3d: func() int { // DEC A.
			cpu.SetA(cpu.opDec(cpu.A()))
			return 4
		},
		0x3f: func() int { // CCF.
			cpu.SetFlagN(false)
			cpu.SetFlagH(false)
			cpu.SetFlagC(!cpu.FlagC())
			return 4
		},
		0x80: func() int { // ADD A,B.
			cpu.SetA(cpu.opAdd(cpu.A(), cpu.B(), false))
			return 4
		},
		0x81: func() int { // ADD A,C.
			cpu.SetA(cpu.opAdd(cpu.A(), cpu.C(), false))
			return 4
		},
		0x82: func() int { // ADD A,D.
			cpu.SetA(cpu.opAdd(cpu.A(), cpu.D(), false))
			return 4
		},
		0x83: func() int { // ADD A,E.
			cpu.SetA(cpu.opAdd(cpu.A(), cpu.E(), false))
			return 4
		},
		0x84: func() int { // ADD A,H.
			cpu.SetA(cpu.opAdd(cpu.A(), cpu.H(), false))
			return 4
		},
		0x85: func() int { // ADD A,L.
			cpu.SetA(cpu.opAdd(cpu.A(), cpu.L(), false))
			return 4
		},
		0x86: func() int { // ADD A,(HL).
			cpu.SetA(cpu.opAdd(cpu.A(), mem.Read(cpu.HL()), false))
			return 8
		},
		0x87: func() int { // ADD A,A.
			cpu.SetA(cpu.opAdd(cpu.A(), cpu.A(), false))
			return 4
		},
		0x88: func() int { // ADC A,B.
			cpu.SetA(cpu.opAdd(cpu.A(), cpu.B(), cpu.FlagC()))
			return 4
		},
		0x89: func() int { // ADC A,C.
			cpu.SetA(cpu.opAdd(cpu.A(), cpu.C(), cpu.FlagC()))
			return 4
		},
		0x8a: func() int { // ADC A,D.
			cpu.SetA(cpu.opAdd(cpu.A(), cpu.D(), cpu.FlagC()))
			return 4
		},
		0x8b: func() int { // ADC A,E.
			cpu.SetA(cpu.opAdd(cpu.A(), cpu.E(), cpu.FlagC()))
			return 4
		},
		0x8c: func() int { // ADC A,H.
			cpu.SetA(cpu.opAdd(cpu.A(), cpu.H(), cpu.FlagC()))
			return 4
		},
		0x8d: func() int { // ADC A,L.
			cpu.SetA(cpu.opAdd(cpu.A(), cpu.L(), cpu.FlagC()))
			return 4
		},
		0x8e: func() int { // ADC A,(HL).
			cpu.SetA(cpu.opAdd(cpu.A(), mem.Read(cpu.HL()), cpu.FlagC()))
			return 8
		},
		0x8f: func() int { // ADC A,A.
			cpu.SetA(cpu.opAdd(cpu.A(), cpu.A(), cpu.FlagC()))
			return 4
		},
		0x90: func() int { // SUB B.
			cpu.SetA(cpu.opSub(cpu.A(), cpu.B(), false))
			return 4
		},
		0x91: func() int { // SUB C.
			cpu.SetA(cpu.opSub(cpu.A(), cpu.C(), false))
			return 4
		},
		0x92: func() int { // SUB D.
			cpu.SetA(cpu.opSub(cpu.A(), cpu.D(), false))
			return 4
		},
		0x93: func() int { // SUB E.
			cpu.SetA(cpu.opSub(cpu.A(), cpu.E(), false))
			return 4
		},
		0x94: func() int { // SUB H.
			cpu.SetA(cpu.opSub(cpu.A(), cpu.H(), false))
			return 4
		},
		0x95: func() int { // SUB L.
			cpu.SetA(cpu.opSub(cpu.A(), cpu.L(), false))
			return 4
		},
		0x96: func() int { // SUB (HL).
			cpu.SetA(cpu.opSub(cpu.A(), mem.Read(cpu.HL()), false))
			return 8
		},
		0x97: func() int { // SUB A.
			cpu.SetA(cpu.opSub(cpu.A(), cpu.A(), false))
			return 4
		},
		0x98: func() int { // SBC A,B.
			cpu.SetA(cpu.opSub(cpu.A(), cpu.B(), cpu.FlagC()))
			return 4
		},
		0x99: func() int { // SBC A,C.
			cpu.SetA(cpu.opSub(cpu.A(), cpu.C(), cpu.FlagC()))
			return 4
		},
		0x9a: func() int { // SBC A,D.
			cpu.SetA(cpu.opSub(cpu.A(), cpu.D(), cpu.FlagC()))
			return 4
		},
		0x9b: func() int { // SBC A,E.
			cpu.SetA(cpu.opSub(cpu.A(), cpu.E(), cpu.FlagC()))
			return 4
		},
		0x9c: func() int { // SBC A,H.
			cpu.SetA(cpu.opSub(cpu.A(), cpu.H(), cpu.FlagC()))
			return 4
		},
		0x9d: func() int { // SBC A,L.
			cpu.SetA(cpu.opSub(cpu.A(), cpu.L(), cpu.FlagC()))
			return 4
		},
		0x9e: func() int { // SBC A,(HL).
			cpu.SetA(cpu.opSub(cpu.A(), mem.Read(cpu.HL()), cpu.FlagC()))
			return 8
		},
		0x9f: func() int { // SBC A,A.
			cpu.SetA(cpu.opSub(cpu.A(), cpu.A(), cpu.FlagC()))
			return 4
		},
		0xa0: func() int { // AND B.
			cpu.SetA(cpu.opAnd(cpu.A(), cpu.B()))
			return 4
		},
		0xa1: func() int { // AND C.
			cpu.SetA(cpu.opAnd(cpu.A(), cpu.C()))
			return 4
		},
		0xa2: func() int { // AND D.
			cpu.SetA(cpu.opAnd(cpu.A(), cpu.D()))
			return 4
		},
		0xa3: func() int { // AND E.
			cpu.SetA(cpu.opAnd(cpu.A(), cpu.E()))
			return 4
		},
		0xa4: func() int { // AND H.
			cpu.SetA(cpu.opAnd(cpu.A(), cpu.H()))
			return 4
		},
		0xa5: func() int { // AND L.
			cpu.SetA(cpu.opAnd(cpu.A(), cpu.L()))
			return 4
		},
		0xa6: func() int { // AND (HL).
			cpu.SetA(cpu.opAnd(cpu.A(), mem.Read(cpu.HL())))
			return 8
		},
		0xa7: func() int { // AND A.
			cpu.SetA(cpu.opAnd(cpu.A(), cpu.A()))
			return 4
		},
		0xa8: func() int { // XOR B.
			cpu.SetA(cpu.opXor(cpu.A(), cpu.B()))
			return 4
		},
		0xa9: func() int { // XOR C.
			cpu.SetA(cpu.opXor(cpu.A(), cpu.C()))
			return 4
		},
		0xaa: func() int { // XOR D.
			cpu.SetA(cpu.opXor(cpu.A(), cpu.D()))
			return 4
		},
		0xab: func() int { // XOR E.
			cpu.SetA(cpu.opXor(cpu.A(), cpu.E()))
			return 4
		},
		0xac: func() int { // XOR H.
			cpu.SetA(cpu.opXor(cpu.A(), cpu.H()))
			return 4
		},
		0xad: func() int { // XOR L.
			cpu.SetA(cpu.opXor(cpu.A(), cpu.L()))
			return 4
		},
		0xae: func() int { // XOR (HL).
			cpu.SetA(cpu.opXor(cpu.A(), mem.Read(cpu.HL())))
			return 8
		},
		0xaf: func() int { // XOR A.
			cpu.SetA(cpu.opXor(cpu.A(), cpu.A()))
			return 4
		},
		0xb0: func() int { // OR B.
			cpu.SetA(cpu.opOr(cpu.A(), cpu.B()))
			return 4
		},
		0xb1: func() int { // OR C.
			cpu.SetA(cpu.opOr(cpu.A(), cpu.C()))
			return 4
		},
		0xb2: func() int { // OR D.
			cpu.SetA(cpu.opOr(cpu.A(), cpu.D()))
			return 4
		},
		0xb3: func() int { // OR E.
			cpu.SetA(cpu.opOr(cpu.A(), cpu.E()))
			return 4
		},
		0xb4: func() int { // OR H.
			cpu.SetA(cpu.opOr(cpu.A(), cpu.H()))
			return 4
		},
		0xb5: func() int { // OR L.
			cpu.SetA(cpu.opOr(cpu.A(), cpu.L()))
			return 4
		},
		0xb6: func() int { // OR (HL).
			cpu.SetA(cpu.opOr(cpu.A(), mem.Read(cpu.HL())))
			return 8
		},
		0xb7: func() int { // OR A.
			cpu.SetA(cpu.opOr(cpu.A(), cpu.A()))
			return 4
		},
		0xb8: func() int { // CP B.
			cpu.opCp(cpu.A(), cpu.B())
			return 4
		},
		0xb9: func() int { // CP C.
			cpu.opCp(cpu.A(), cpu.C())
			return 4
		},
		0xba: func() int { // CP D.
			cpu.opCp(cpu.A(), cpu.D())
			return 4
		},
		0xbb: func() int { // CP E.
			cpu.opCp(cpu.A(), cpu.E())
			return 4
		},
		0xbc: func() int { // CP H.
			cpu.opCp(cpu.A(), cpu.H())
			return 4
		},
		0xbd: func() int { // CP L.
			cpu.opCp(cpu.A(), cpu.L())
			return 4
		},
		0xbe: func() int { // CP (HL).
			cpu.opCp(cpu.A(), mem.Read(cpu.HL()))
			return 8
		},
		0xbf: func() int { // CP A.
			cpu.opCp(cpu.A(), cpu.A())
			return 4
		},
		0xc6: func() int { // ADD A,d8.
			cpu.SetA(cpu.opAdd(cpu.A(), cpu.IncPC(), false))
			return 8
		},
		0xce: func() int { // ADC A,d8.
			cpu.SetA(cpu.opAdd(cpu.A(), cpu.IncPC(), cpu.FlagC()))
			return 8
		},
		0xd6: func() int { // SUB d8.
			cpu.SetA(cpu.opSub(cpu.A(), cpu.IncPC(), false))
			return 8
		},
		0xde: func() int { // SBC A,d8.
			cpu.SetA(cpu.opSub(cpu.A(), cpu.IncPC(), cpu.FlagC()))
			return 8
		},
		0xe6: func() int { // AND d8.
			cpu.SetA(cpu.opAnd(cpu.A(), cpu.IncPC()))
			return 8
		},
		0xee: func() int { // XOR d8.
			cpu.SetA(cpu.opXor(cpu.A(), cpu.IncPC()))
			return 8
		},
		0xf6: func() int { // OR d8.
			cpu.SetA(cpu.opOr(cpu.A(), cpu.IncPC()))
			return 8
		},
		0xfe: func() int { // CP d8.
			cpu.opCp(cpu.A(), cpu.IncPC())
			return 8
		},

		// 16 bit arithmetic.
		0x03: func() int { // INC BC.
			cpu.SetBC(cpu.BC() + 1)
			return 8
		},
		0x09: func() int { // ADD HL,BC.
			cpu.SetHL(cpu.opAdd16(cpu.HL(), cpu.BC()))
			return 8
		},
		0x0b: func() int { // DEC BC.
			cpu.SetBC(cpu.BC() - 1)
			return 8
		},
		0x13: func() int { // INC DE.
			cpu.SetDE(cpu.DE() + 1)
			return 8
		},
		0x19: func() int { // ADD HL,DE.
			cpu.SetHL(cpu.opAdd16(cpu.HL(), cpu.DE()))
			return 8
		},
		0x1b: func() int { // DEC DE.
			cpu.SetDE(cpu.DE() - 1)
			return 8
		},
		0x23: func() int { // INC HL.
			cpu.SetHL(cpu.HL() + 1)
			return 8
		},
		0x29: func() int { // ADD HL,HL.
			cpu.SetHL(cpu.opAdd16(cpu.HL(), cpu.HL()))
			return 8
		},
		0x2b: func() int { // DEC HL.
			cpu.SetHL(cpu.HL() - 1)
			return 8
		},
		0x33: func() int { // INC SP.
			cpu.SetSP(cpu.SP() + 1)
			return 8
		},
		0x39: func() int { // ADD HL,SP.
			cpu.SetHL(cpu.opAdd16(cpu.HL(), cpu.SP()))
			return 8
		},
		0x3b: func() int { // DEC SP.
			cpu.SetSP(cpu.SP() - 1)
			return 8
		},
		0xe8: func() int { // ADD SP,r8.
			cpu.SetSP(cpu.opSignedAdd(cpu.SP(), cpu.IncPC()))
			return 16
		},

		// Standard 8 bit rotations and shifts.
		0x07: func() int { // RLCA.
			cpu.SetA(cpu.opRl(cpu.A(), false))
			cpu.SetFlagZ(false)
			return 4
		},
		0x0f: func() int { // RRCA.
			cpu.SetA(cpu.opRr(cpu.A(), false))
			cpu.SetFlagZ(false)
			return 4
		},
		0x17: func() int { // RLA.
			cpu.SetA(cpu.opRl(cpu.A(), true))
			cpu.SetFlagZ(false)
			return 4
		},
		0x1f: func() int { // RRA.
			cpu.SetA(cpu.opRr(cpu.A(), true))
			cpu.SetFlagZ(false)
			return 4
		},

		// Jumps and calls.
		0x18: func() int { // JR r8.
			cpu.opJr(true, cpu.IncPC())
			return 12
		},
		0x20: func() int { // JR NZ,r8.
			return cpu.opJr(!cpu.FlagZ(), cpu.IncPC())
		},
		0x28: func() int { // JR Z,r8.
			return cpu.opJr(cpu.FlagZ(), cpu.IncPC())
		},
		0x30: func() int { // JR NC,r8.
			return cpu.opJr(!cpu.FlagC(), cpu.IncPC())
		},
		0x38: func() int { // JR C,r8.
			return cpu.opJr(cpu.FlagC(), cpu.IncPC())
		},
		0xc0: func() int { // RET NZ.
			return cpu.opRet(!cpu.FlagZ())
		},
		0xc2: func() int { // JP NZ,a16.
			return cpu.opJp(!cpu.FlagZ(), cpu.IncPC16())
		},
		0xc3: func() int { // JP a16.
			cpu.opJp(true, cpu.IncPC16())
			return 16
		},
		0xc4: func() int { // CALL NZ,a16.
			return cpu.opCall(!cpu.FlagZ(), cpu.IncPC16())
		},
		0xc7: func() int { // RST 00H.
			cpu.opCall(true, 0x0000)
			return 16
		},
		0xc8: func() int { // RET Z.
			return cpu.opRet(cpu.FlagZ())
		},
		0xc9: func() int { // RET.
			cpu.opRet(true)
			return 16
		},
		0xca: func() int { // JP Z,a16.
			return cpu.opJp(cpu.FlagZ(), cpu.IncPC16())
		},
		0xcc: func() int { // CALL Z,a16.
			return cpu.opCall(cpu.FlagZ(), cpu.IncPC16())
		},
		0xcd: func() int { // CALL a16.
			cpu.opCall(true, cpu.IncPC16())
			return 24
		},
		0xcf: func() int { // RST 08H.
			cpu.opCall(true, 0x0008)
			return 16
		},
		0xd0: func() int { // RET NC.
			return cpu.opRet(!cpu.FlagC())
		},
		0xd2: func() int { // JP NC,a16.
			return cpu.opJp(!cpu.FlagC(), cpu.IncPC16())
		},
		0xd4: func() int { // CALL NC,a16.
			return cpu.opCall(!cpu.FlagC(), cpu.IncPC16())
		},
		0xd7: func() int { // RST 10H.
			cpu.opCall(true, 0x0010)
			return 16
		},
		0xd8: func() int { // RET C.
			return cpu.opRet(cpu.FlagC())
		},
		0xd9: func() int { // RETI.
			cpu.opRet(true)
			cpu.SetIME(true)
			return 16
		},
		0xda: func() int { // JP C,a16.
			return cpu.opJp(cpu.FlagC(), cpu.IncPC16())
		},
		0xdc: func() int { // CALL C,a16.
			return cpu.opCall(cpu.FlagC(), cpu.IncPC16())
		},
		0xdf: func() int { // RST 18H.
			cpu.opCall(true, 0x0018)
			return 16
		},
		0xe7: func() int { // RST 20H.
			cpu.opCall(true, 0x0020)
			return 16
		},
		0xe9: func() int { // JP (HL).
			cpu.opJp(true, cpu.HL())
			return 4
		},
		0xef: func() int { // RST 28H.
			cpu.opCall(true, 0x0028)
			return 16
		},
		0xf7: func() int { // RST 30H.
			cpu.opCall(true, 0x0030)
			return 16
		},
		0xff: func() int { // RST 38H.
			cpu.opCall(true, 0x0038)
			return 16
		},

		// Control instructions.
		0x00: func() int { // NOP.
			return 4
		},
		0x10: func() int { // STOP.
			// TODO: Properly implement STOP.
			cpu.SetHalt(true)
			cpu.IncPC()
			return 4
		},
		0x76: func() int { // HALT.
			cpu.opHalt(cpu.IME(), mem.Read(AddrIE), mem.Read(AddrIF))
			return 4
		},
		0xcb: func() int { // PREFIX CB.
			return cpu.instructions[uint16(cpu.IncPC())+0x100]()
		},
		0xf3: func() int { // DI.
			cpu.SetIME(false)
			return 4
		},
		0xfb: func() int { // EI.
			cpu.SetIME(true)
			return 4
		},

		// CB prefix extensions.
		0x100: func() int { // RLC B.
			cpu.SetB(cpu.opRl(cpu.B(), false))
			return 8
		},
		0x101: func() int { // RLC C.
			cpu.SetC(cpu.opRl(cpu.C(), false))
			return 8
		},
		0x102: func() int { // RLC D.
			cpu.SetD(cpu.opRl(cpu.D(), false))
			return 8
		},
		0x103: func() int { // RLC E.
			cpu.SetE(cpu.opRl(cpu.E(), false))
			return 8
		},
		0x104: func() int { // RLC H.
			cpu.SetH(cpu.opRl(cpu.H(), false))
			return 8
		},
		0x105: func() int { // RLC L.
			cpu.SetL(cpu.opRl(cpu.L(), false))
			return 8
		},
		0x106: func() int { // RLC (HL).
			mem.Write(cpu.HL(), cpu.opRl(mem.Read(cpu.HL()), false))
			return 16
		},
		0x107: func() int { // RLC A.
			cpu.SetA(cpu.opRl(cpu.A(), false))
			return 8
		},
		0x108: func() int { // RRC B.
			cpu.SetB(cpu.opRr(cpu.B(), false))
			return 8
		},
		0x109: func() int { // RRC C.
			cpu.SetC(cpu.opRr(cpu.C(), false))
			return 8
		},
		0x10a: func() int { // RRC D.
			cpu.SetD(cpu.opRr(cpu.D(), false))
			return 8
		},
		0x10b: func() int { // RRC E.
			cpu.SetE(cpu.opRr(cpu.E(), false))
			return 8
		},
		0x10c: func() int { // RRC H.
			cpu.SetH(cpu.opRr(cpu.H(), false))
			return 8
		},
		0x10d: func() int { // RRC L.
			cpu.SetL(cpu.opRr(cpu.L(), false))
			return 8
		},
		0x10e: func() int { // RRC (HL).
			mem.Write(cpu.HL(), cpu.opRr(mem.Read(cpu.HL()), false))
			return 16
		},
		0x10f: func() int { // RRC A.
			cpu.SetA(cpu.opRr(cpu.A(), false))
			return 8
		},
		0x110: func() int { // RL B.
			cpu.SetB(cpu.opRl(cpu.B(), true))
			return 8
		},
		0x111: func() int { // RL C.
			cpu.SetC(cpu.opRl(cpu.C(), true))
			return 8
		},
		0x112: func() int { // RL D.
			cpu.SetD(cpu.opRl(cpu.D(), true))
			return 8
		},
		0x113: func() int { // RL E.
			cpu.SetE(cpu.opRl(cpu.E(), true))
			return 8
		},
		0x114: func() int { // RL H.
			cpu.SetH(cpu.opRl(cpu.H(), true))
			return 8
		},
		0x115: func() int { // RL L.
			cpu.SetL(cpu.opRl(cpu.L(), true))
			return 8
		},
		0x116: func() int { // RL (HL).
			mem.Write(cpu.HL(), cpu.opRl(mem.Read(cpu.HL()), true))
			return 16
		},
		0x117: func() int { // RL A.
			cpu.SetA(cpu.opRl(cpu.A(), true))
			return 8
		},
		0x118: func() int { // RR B.
			cpu.SetB(cpu.opRr(cpu.B(), true))
			return 8
		},
		0x119: func() int { // RR C.
			cpu.SetC(cpu.opRr(cpu.C(), true))
			return 8
		},
		0x11a: func() int { // RR D.
			cpu.SetD(cpu.opRr(cpu.D(), true))
			return 8
		},
		0x11b: func() int { // RR E.
			cpu.SetE(cpu.opRr(cpu.E(), true))
			return 8
		},
		0x11c: func() int { // RR H.
			cpu.SetH(cpu.opRr(cpu.H(), true))
			return 8
		},
		0x11d: func() int { // RR L.
			cpu.SetL(cpu.opRr(cpu.L(), true))
			return 8
		},
		0x11e: func() int { // RR (HL).
			mem.Write(cpu.HL(), cpu.opRr(mem.Read(cpu.HL()), true))
			return 16
		},
		0x11f: func() int { // RR A.
			cpu.SetA(cpu.opRr(cpu.A(), true))
			return 8
		},
		0x120: func() int { // SLA B.
			cpu.SetB(cpu.opSl(cpu.B()))
			return 8
		},
		0x121: func() int { // SLA C.
			cpu.SetC(cpu.opSl(cpu.C()))
			return 8
		},
		0x122: func() int { // SLA D.
			cpu.SetD(cpu.opSl(cpu.D()))
			return 8
		},
		0x123: func() int { // SLA E.
			cpu.SetE(cpu.opSl(cpu.E()))
			return 8
		},
		0x124: func() int { // SLA H.
			cpu.SetH(cpu.opSl(cpu.H()))
			return 8
		},
		0x125: func() int { // SLA L.
			cpu.SetL(cpu.opSl(cpu.L()))
			return 8
		},
		0x126: func() int { // SLA (HL).
			mem.Write(cpu.HL(), cpu.opSl(mem.Read(cpu.HL())))
			return 16
		},
		0x127: func() int { // SLA A.
			cpu.SetA(cpu.opSl(cpu.A()))
			return 8
		},
		0x128: func() int { // SRA B.
			cpu.SetB(cpu.opSr(cpu.B(), true))
			return 8
		},
		0x129: func() int { // SRA C.
			cpu.SetC(cpu.opSr(cpu.C(), true))
			return 8
		},
		0x12a: func() int { // SRA D.
			cpu.SetD(cpu.opSr(cpu.D(), true))
			return 8
		},
		0x12b: func() int { // SRA E.
			cpu.SetE(cpu.opSr(cpu.E(), true))
			return 8
		},
		0x12c: func() int { // SRA H.
			cpu.SetH(cpu.opSr(cpu.H(), true))
			return 8
		},
		0x12d: func() int { // SRA L.
			cpu.SetL(cpu.opSr(cpu.L(), true))
			return 8
		},
		0x12e: func() int { // SRA (HL).
			mem.Write(cpu.HL(), cpu.opSr(mem.Read(cpu.HL()), true))
			return 16
		},
		0x12f: func() int { // SRA A.
			cpu.SetA(cpu.opSr(cpu.A(), true))
			return 8
		},
		0x130: func() int { // SWAP B.
			cpu.SetB(cpu.opSwap(cpu.B()))
			return 8
		},
		0x131: func() int { // SWAP C.
			cpu.SetC(cpu.opSwap(cpu.C()))
			return 8
		},
		0x132: func() int { // SWAP D.
			cpu.SetD(cpu.opSwap(cpu.D()))
			return 8
		},
		0x133: func() int { // SWAP E.
			cpu.SetE(cpu.opSwap(cpu.E()))
			return 8
		},
		0x134: func() int { // SWAP H.
			cpu.SetH(cpu.opSwap(cpu.H()))
			return 8
		},
		0x135: func() int { // SWAP L.
			cpu.SetL(cpu.opSwap(cpu.L()))
			return 8
		},
		0x136: func() int { // SWAP (HL).
			mem.Write(cpu.HL(), cpu.opSwap(mem.Read(cpu.HL())))
			return 16
		},
		0x137: func() int { // SWAP A.
			cpu.SetA(cpu.opSwap(cpu.A()))
			return 8
		},
		0x138: func() int { // SRL B.
			cpu.SetB(cpu.opSr(cpu.B(), false))
			return 8
		},
		0x139: func() int { // SRL C.
			cpu.SetC(cpu.opSr(cpu.C(), false))
			return 8
		},
		0x13a: func() int { // SRL D.
			cpu.SetD(cpu.opSr(cpu.D(), false))
			return 8
		},
		0x13b: func() int { // SRL E.
			cpu.SetE(cpu.opSr(cpu.E(), false))
			return 8
		},
		0x13c: func() int { // SRL H.
			cpu.SetH(cpu.opSr(cpu.H(), false))
			return 8
		},
		0x13d: func() int { // SRL L.
			cpu.SetL(cpu.opSr(cpu.L(), false))
			return 8
		},
		0x13e: func() int { // SRL (HL).
			mem.Write(cpu.HL(), cpu.opSr(mem.Read(cpu.HL()), false))
			return 16
		},
		0x13f: func() int { // SRL A.
			cpu.SetA(cpu.opSr(cpu.A(), false))
			return 8
		},
		0x140: func() int { // BIT 0,B.
			cpu.opBit(cpu.B(), 0)
			return 8
		},
		0x141: func() int { // BIT 0,C.
			cpu.opBit(cpu.C(), 0)
			return 8
		},
		0x142: func() int { // BIT 0,D.
			cpu.opBit(cpu.D(), 0)
			return 8
		},
		0x143: func() int { // BIT 0,E.
			cpu.opBit(cpu.E(), 0)
			return 8
		},
		0x144: func() int { // BIT 0,H.
			cpu.opBit(cpu.H(), 0)
			return 8
		},
		0x145: func() int { // BIT 0,L.
			cpu.opBit(cpu.L(), 0)
			return 8
		},
		0x146: func() int { // BIT 0,(HL).
			cpu.opBit(mem.Read(cpu.HL()), 0)
			return 12
		},
		0x147: func() int { // BIT 0,A.
			cpu.opBit(cpu.A(), 0)
			return 8
		},
		0x148: func() int { // BIT 1,B.
			cpu.opBit(cpu.B(), 1)
			return 8
		},
		0x149: func() int { // BIT 1,C.
			cpu.opBit(cpu.C(), 1)
			return 8
		},
		0x14a: func() int { // BIT 1,D.
			cpu.opBit(cpu.D(), 1)
			return 8
		},
		0x14b: func() int { // BIT 1,E.
			cpu.opBit(cpu.E(), 1)
			return 8
		},
		0x14c: func() int { // BIT 1,H.
			cpu.opBit(cpu.H(), 1)
			return 8
		},
		0x14d: func() int { // BIT 1,L.
			cpu.opBit(cpu.L(), 1)
			return 8
		},
		0x14e: func() int { // BIT 1,(HL).
			cpu.opBit(mem.Read(cpu.HL()), 1)
			return 12
		},
		0x14f: func() int { // BIT 1,A.
			cpu.opBit(cpu.A(), 1)
			return 8
		},
		0x150: func() int { // BIT 2,B.
			cpu.opBit(cpu.B(), 2)
			return 8
		},
		0x151: func() int { // BIT 2,C.
			cpu.opBit(cpu.C(), 2)
			return 8
		},
		0x152: func() int { // BIT 2,D.
			cpu.opBit(cpu.D(), 2)
			return 8
		},
		0x153: func() int { // BIT 2,E.
			cpu.opBit(cpu.E(), 2)
			return 8
		},
		0x154: func() int { // BIT 2,H.
			cpu.opBit(cpu.H(), 2)
			return 8
		},
		0x155: func() int { // BIT 2,L.
			cpu.opBit(cpu.L(), 2)
			return 8
		},
		0x156: func() int { // BIT 2,(HL).
			cpu.opBit(mem.Read(cpu.HL()), 2)
			return 12
		},
		0x157: func() int { // BIT 2,A.
			cpu.opBit(cpu.A(), 2)
			return 8
		},
		0x158: func() int { // BIT 3,B.
			cpu.opBit(cpu.B(), 3)
			return 8
		},
		0x159: func() int { // BIT 3,C.
			cpu.opBit(cpu.C(), 3)
			return 8
		},
		0x15a: func() int { // BIT 3,D.
			cpu.opBit(cpu.D(), 3)
			return 8
		},
		0x15b: func() int { // BIT 3,E.
			cpu.opBit(cpu.E(), 3)
			return 8
		},
		0x15c: func() int { // BIT 3,H.
			cpu.opBit(cpu.H(), 3)
			return 8
		},
		0x15d: func() int { // BIT 3,L.
			cpu.opBit(cpu.L(), 3)
			return 8
		},
		0x15e: func() int { // BIT 3,(HL).
			cpu.opBit(mem.Read(cpu.HL()), 3)
			return 12
		},
		0x15f: func() int { // BIT 3,A.
			cpu.opBit(cpu.A(), 3)
			return 8
		},
		0x160: func() int { // BIT 4,B.
			cpu.opBit(cpu.B(), 4)
			return 8
		},
		0x161: func() int { // BIT 4,C.
			cpu.opBit(cpu.C(), 4)
			return 8
		},
		0x162: func() int { // BIT 4,D.
			cpu.opBit(cpu.D(), 4)
			return 8
		},
		0x163: func() int { // BIT 4,E.
			cpu.opBit(cpu.E(), 4)
			return 8
		},
		0x164: func() int { // BIT 4,H.
			cpu.opBit(cpu.H(), 4)
			return 8
		},
		0x165: func() int { // BIT 4,L.
			cpu.opBit(cpu.L(), 4)
			return 8
		},
		0x166: func() int { // BIT 4,(HL).
			cpu.opBit(mem.Read(cpu.HL()), 4)
			return 12
		},
		0x167: func() int { // BIT 4,A.
			cpu.opBit(cpu.A(), 4)
			return 8
		},
		0x168: func() int { // BIT 5,B.
			cpu.opBit(cpu.B(), 5)
			return 8
		},
		0x169: func() int { // BIT 5,C.
			cpu.opBit(cpu.C(), 5)
			return 8
		},
		0x16a: func() int { // BIT 5,D.
			cpu.opBit(cpu.D(), 5)
			return 8
		},
		0x16b: func() int { // BIT 5,E.
			cpu.opBit(cpu.E(), 5)
			return 8
		},
		0x16c: func() int { // BIT 5,H.
			cpu.opBit(cpu.H(), 5)
			return 8
		},
		0x16d: func() int { // BIT 5,L.
			cpu.opBit(cpu.L(), 5)
			return 8
		},
		0x16e: func() int { // BIT 5,(HL).
			cpu.opBit(mem.Read(cpu.HL()), 5)
			return 12
		},
		0x16f: func() int { // BIT 5,A.
			cpu.opBit(cpu.A(), 5)
			return 8
		},
		0x170: func() int { // BIT 6,B.
			cpu.opBit(cpu.B(), 6)
			return 8
		},
		0x171: func() int { // BIT 6,C.
			cpu.opBit(cpu.C(), 6)
			return 8
		},
		0x172: func() int { // BIT 6,D.
			cpu.opBit(cpu.D(), 6)
			return 8
		},
		0x173: func() int { // BIT 6,E.
			cpu.opBit(cpu.E(), 6)
			return 8
		},
		0x174: func() int { // BIT 6,H.
			cpu.opBit(cpu.H(), 6)
			return 8
		},
		0x175: func() int { // BIT 6,L.
			cpu.opBit(cpu.L(), 6)
			return 8
		},
		0x176: func() int { // BIT 6,(HL).
			cpu.opBit(mem.Read(cpu.HL()), 6)
			return 12
		},
		0x177: func() int { // BIT 6,A.
			cpu.opBit(cpu.A(), 6)
			return 8
		},
		0x178: func() int { // BIT 7,B.
			cpu.opBit(cpu.B(), 7)
			return 8
		},
		0x179: func() int { // BIT 7,C.
			cpu.opBit(cpu.C(), 7)
			return 8
		},
		0x17a: func() int { // BIT 7,D.
			cpu.opBit(cpu.D(), 7)
			return 8
		},
		0x17b: func() int { // BIT 7,E.
			cpu.opBit(cpu.E(), 7)
			return 8
		},
		0x17c: func() int { // BIT 7,H.
			cpu.opBit(cpu.H(), 7)
			return 8
		},
		0x17d: func() int { // BIT 7,L.
			cpu.opBit(cpu.L(), 7)
			return 8
		},
		0x17e: func() int { // BIT 7,(HL).
			cpu.opBit(mem.Read(cpu.HL()), 7)
			return 12
		},
		0x17f: func() int { // BIT 7,A.
			cpu.opBit(cpu.A(), 7)
			return 8
		},
		0x180: func() int { // RES 0,B.
			cpu.SetB(utils.SetBit(cpu.B(), 0, false))
			return 8
		},
		0x181: func() int { // RES 0,C.
			cpu.SetC(utils.SetBit(cpu.C(), 0, false))
			return 8
		},
		0x182: func() int { // RES 0,D.
			cpu.SetD(utils.SetBit(cpu.D(), 0, false))
			return 8
		},
		0x183: func() int { // RES 0,E.
			cpu.SetE(utils.SetBit(cpu.E(), 0, false))
			return 8
		},
		0x184: func() int { // RES 0,H.
			cpu.SetH(utils.SetBit(cpu.H(), 0, false))
			return 8
		},
		0x185: func() int { // RES 0,L.
			cpu.SetL(utils.SetBit(cpu.L(), 0, false))
			return 8
		},
		0x186: func() int { // RES 0,(HL).
			mem.Write(cpu.HL(), utils.SetBit(mem.Read(cpu.HL()), 0, false))
			return 16
		},
		0x187: func() int { // RES 0,A.
			cpu.SetA(utils.SetBit(cpu.A(), 0, false))
			return 8
		},
		0x188: func() int { // RES 1,B.
			cpu.SetB(utils.SetBit(cpu.B(), 1, false))
			return 8
		},
		0x189: func() int { // RES 1,C.
			cpu.SetC(utils.SetBit(cpu.C(), 1, false))
			return 8
		},
		0x18a: func() int { // RES 1,D.
			cpu.SetD(utils.SetBit(cpu.D(), 1, false))
			return 8
		},
		0x18b: func() int { // RES 1,E.
			cpu.SetE(utils.SetBit(cpu.E(), 1, false))
			return 8
		},
		0x18c: func() int { // RES 1,H.
			cpu.SetH(utils.SetBit(cpu.H(), 1, false))
			return 8
		},
		0x18d: func() int { // RES 1,L.
			cpu.SetL(utils.SetBit(cpu.L(), 1, false))
			return 8
		},
		0x18e: func() int { // RES 1,(HL).
			mem.Write(cpu.HL(), utils.SetBit(mem.Read(cpu.HL()), 1, false))
			return 16
		},
		0x18f: func() int { // RES 1,A.
			cpu.SetA(utils.SetBit(cpu.A(), 1, false))
			return 8
		},
		0x190: func() int { // RES 2,B.
			cpu.SetB(utils.SetBit(cpu.B(), 2, false))
			return 8
		},
		0x191: func() int { // RES 2,C.
			cpu.SetC(utils.SetBit(cpu.C(), 2, false))
			return 8
		},
		0x192: func() int { // RES 2,D.
			cpu.SetD(utils.SetBit(cpu.D(), 2, false))
			return 8
		},
		0x193: func() int { // RES 2,E.
			cpu.SetE(utils.SetBit(cpu.E(), 2, false))
			return 8
		},
		0x194: func() int { // RES 2,H.
			cpu.SetH(utils.SetBit(cpu.H(), 2, false))
			return 8
		},
		0x195: func() int { // RES 2,L.
			cpu.SetL(utils.SetBit(cpu.L(), 2, false))
			return 8
		},
		0x196: func() int { // RES 2,(HL).
			mem.Write(cpu.HL(), utils.SetBit(mem.Read(cpu.HL()), 2, false))
			return 16
		},
		0x197: func() int { // RES 2,A.
			cpu.SetA(utils.SetBit(cpu.A(), 2, false))
			return 8
		},
		0x198: func() int { // RES 3,B.
			cpu.SetB(utils.SetBit(cpu.B(), 3, false))
			return 8
		},
		0x199: func() int { // RES 3,C.
			cpu.SetC(utils.SetBit(cpu.C(), 3, false))
			return 8
		},
		0x19a: func() int { // RES 3,D.
			cpu.SetD(utils.SetBit(cpu.D(), 3, false))
			return 8
		},
		0x19b: func() int { // RES 3,E.
			cpu.SetE(utils.SetBit(cpu.E(), 3, false))
			return 8
		},
		0x19c: func() int { // RES 3,H.
			cpu.SetH(utils.SetBit(cpu.H(), 3, false))
			return 8
		},
		0x19d: func() int { // RES 3,L.
			cpu.SetL(utils.SetBit(cpu.L(), 3, false))
			return 8
		},
		0x19e: func() int { // RES 3,(HL).
			mem.Write(cpu.HL(), utils.SetBit(mem.Read(cpu.HL()), 3, false))
			return 16
		},
		0x19f: func() int { // RES 3,A.
			cpu.SetA(utils.SetBit(cpu.A(), 3, false))
			return 8
		},
		0x1a0: func() int { // RES 4,B.
			cpu.SetB(utils.SetBit(cpu.B(), 4, false))
			return 8
		},
		0x1a1: func() int { // RES 4,C.
			cpu.SetC(utils.SetBit(cpu.C(), 4, false))
			return 8
		},
		0x1a2: func() int { // RES 4,D.
			cpu.SetD(utils.SetBit(cpu.D(), 4, false))
			return 8
		},
		0x1a3: func() int { // RES 4,E.
			cpu.SetE(utils.SetBit(cpu.E(), 4, false))
			return 8
		},
		0x1a4: func() int { // RES 4,H.
			cpu.SetH(utils.SetBit(cpu.H(), 4, false))
			return 8
		},
		0x1a5: func() int { // RES 4,L.
			cpu.SetL(utils.SetBit(cpu.L(), 4, false))
			return 8
		},
		0x1a6: func() int { // RES 4,(HL).
			mem.Write(cpu.HL(), utils.SetBit(mem.Read(cpu.HL()), 4, false))
			return 16
		},
		0x1a7: func() int { // RES 4,A.
			cpu.SetA(utils.SetBit(cpu.A(), 4, false))
			return 8
		},
		0x1a8: func() int { // RES 5,B.
			cpu.SetB(utils.SetBit(cpu.B(), 5, false))
			return 8
		},
		0x1a9: func() int { // RES 5,C.
			cpu.SetC(utils.SetBit(cpu.C(), 5, false))
			return 8
		},
		0x1aa: func() int { // RES 5,D.
			cpu.SetD(utils.SetBit(cpu.D(), 5, false))
			return 8
		},
		0x1ab: func() int { // RES 5,E.
			cpu.SetE(utils.SetBit(cpu.E(), 5, false))
			return 8
		},
		0x1ac: func() int { // RES 5,H.
			cpu.SetH(utils.SetBit(cpu.H(), 5, false))
			return 8
		},
		0x1ad: func() int { // RES 5,L.
			cpu.SetL(utils.SetBit(cpu.L(), 5, false))
			return 8
		},
		0x1ae: func() int { // RES 5,(HL).
			mem.Write(cpu.HL(), utils.SetBit(mem.Read(cpu.HL()), 5, false))
			return 16
		},
		0x1af: func() int { // RES 5,A.
			cpu.SetA(utils.SetBit(cpu.A(), 5, false))
			return 8
		},
		0x1b0: func() int { // RES 6,B.
			cpu.SetB(utils.SetBit(cpu.B(), 6, false))
			return 8
		},
		0x1b1: func() int { // RES 6,C.
			cpu.SetC(utils.SetBit(cpu.C(), 6, false))
			return 8
		},
		0x1b2: func() int { // RES 6,D.
			cpu.SetD(utils.SetBit(cpu.D(), 6, false))
			return 8
		},
		0x1b3: func() int { // RES 6,E.
			cpu.SetE(utils.SetBit(cpu.E(), 6, false))
			return 8
		},
		0x1b4: func() int { // RES 6,H.
			cpu.SetH(utils.SetBit(cpu.H(), 6, false))
			return 8
		},
		0x1b5: func() int { // RES 6,L.
			cpu.SetL(utils.SetBit(cpu.L(), 6, false))
			return 8
		},
		0x1b6: func() int { // RES 6,(HL).
			mem.Write(cpu.HL(), utils.SetBit(mem.Read(cpu.HL()), 6, false))
			return 16
		},
		0x1b7: func() int { // RES 6,A.
			cpu.SetA(utils.SetBit(cpu.A(), 6, false))
			return 8
		},
		0x1b8: func() int { // RES 7,B.
			cpu.SetB(utils.SetBit(cpu.B(), 7, false))
			return 8
		},
		0x1b9: func() int { // RES 7,C.
			cpu.SetC(utils.SetBit(cpu.C(), 7, false))
			return 8
		},
		0x1ba: func() int { // RES 7,D.
			cpu.SetD(utils.SetBit(cpu.D(), 7, false))
			return 8
		},
		0x1bb: func() int { // RES 7,E.
			cpu.SetE(utils.SetBit(cpu.E(), 7, false))
			return 8
		},
		0x1bc: func() int { // RES 7,H.
			cpu.SetH(utils.SetBit(cpu.H(), 7, false))
			return 8
		},
		0x1bd: func() int { // RES 7,L.
			cpu.SetL(utils.SetBit(cpu.L(), 7, false))
			return 8
		},
		0x1be: func() int { // RES 7,(HL).
			mem.Write(cpu.HL(), utils.SetBit(mem.Read(cpu.HL()), 7, false))
			return 16
		},
		0x1bf: func() int { // RES 7,A.
			cpu.SetA(utils.SetBit(cpu.A(), 7, false))
			return 8
		},
		0x1c0: func() int { // SET 0,B.
			cpu.SetB(utils.SetBit(cpu.B(), 0, true))
			return 8
		},
		0x1c1: func() int { // SET 0,C.
			cpu.SetC(utils.SetBit(cpu.C(), 0, true))
			return 8
		},
		0x1c2: func() int { // SET 0,D.
			cpu.SetD(utils.SetBit(cpu.D(), 0, true))
			return 8
		},
		0x1c3: func() int { // SET 0,E.
			cpu.SetE(utils.SetBit(cpu.E(), 0, true))
			return 8
		},
		0x1c4: func() int { // SET 0,H.
			cpu.SetH(utils.SetBit(cpu.H(), 0, true))
			return 8
		},
		0x1c5: func() int { // SET 0,L.
			cpu.SetL(utils.SetBit(cpu.L(), 0, true))
			return 8
		},
		0x1c6: func() int { // SET 0,(HL).
			mem.Write(cpu.HL(), utils.SetBit(mem.Read(cpu.HL()), 0, true))
			return 16
		},
		0x1c7: func() int { // SET 0,A.
			cpu.SetA(utils.SetBit(cpu.A(), 0, true))
			return 8
		},
		0x1c8: func() int { // SET 1,B.
			cpu.SetB(utils.SetBit(cpu.B(), 1, true))
			return 8
		},
		0x1c9: func() int { // SET 1,C.
			cpu.SetC(utils.SetBit(cpu.C(), 1, true))
			return 8
		},
		0x1ca: func() int { // SET 1,D.
			cpu.SetD(utils.SetBit(cpu.D(), 1, true))
			return 8
		},
		0x1cb: func() int { // SET 1,E.
			cpu.SetE(utils.SetBit(cpu.E(), 1, true))
			return 8
		},
		0x1cc: func() int { // SET 1,H.
			cpu.SetH(utils.SetBit(cpu.H(), 1, true))
			return 8
		},
		0x1cd: func() int { // SET 1,L.
			cpu.SetL(utils.SetBit(cpu.L(), 1, true))
			return 8
		},
		0x1ce: func() int { // SET 1,(HL).
			mem.Write(cpu.HL(), utils.SetBit(mem.Read(cpu.HL()), 1, true))
			return 16
		},
		0x1cf: func() int { // SET 1,A.
			cpu.SetA(utils.SetBit(cpu.A(), 1, true))
			return 8
		},
		0x1d0: func() int { // SET 2,B.
			cpu.SetB(utils.SetBit(cpu.B(), 2, true))
			return 8
		},
		0x1d1: func() int { // SET 2,C.
			cpu.SetC(utils.SetBit(cpu.C(), 2, true))
			return 8
		},
		0x1d2: func() int { // SET 2,D.
			cpu.SetD(utils.SetBit(cpu.D(), 2, true))
			return 8
		},
		0x1d3: func() int { // SET 2,E.
			cpu.SetE(utils.SetBit(cpu.E(), 2, true))
			return 8
		},
		0x1d4: func() int { // SET 2,H.
			cpu.SetH(utils.SetBit(cpu.H(), 2, true))
			return 8
		},
		0x1d5: func() int { // SET 2,L.
			cpu.SetL(utils.SetBit(cpu.L(), 2, true))
			return 8
		},
		0x1d6: func() int { // SET 2,(HL).
			mem.Write(cpu.HL(), utils.SetBit(mem.Read(cpu.HL()), 2, true))
			return 16
		},
		0x1d7: func() int { // SET 2,A.
			cpu.SetA(utils.SetBit(cpu.A(), 2, true))
			return 8
		},
		0x1d8: func() int { // SET 3,B.
			cpu.SetB(utils.SetBit(cpu.B(), 3, true))
			return 8
		},
		0x1d9: func() int { // SET 3,C.
			cpu.SetC(utils.SetBit(cpu.C(), 3, true))
			return 8
		},
		0x1da: func() int { // SET 3,D.
			cpu.SetD(utils.SetBit(cpu.D(), 3, true))
			return 8
		},
		0x1db: func() int { // SET 3,E.
			cpu.SetE(utils.SetBit(cpu.E(), 3, true))
			return 8
		},
		0x1dc: func() int { // SET 3,H.
			cpu.SetH(utils.SetBit(cpu.H(), 3, true))
			return 8
		},
		0x1dd: func() int { // SET 3,L.
			cpu.SetL(utils.SetBit(cpu.L(), 3, true))
			return 8
		},
		0x1de: func() int { // SET 3,(HL).
			mem.Write(cpu.HL(), utils.SetBit(mem.Read(cpu.HL()), 3, true))
			return 16
		},
		0x1df: func() int { // SET 3,A.
			cpu.SetA(utils.SetBit(cpu.A(), 3, true))
			return 8
		},
		0x1e0: func() int { // SET 4,B.
			cpu.SetB(utils.SetBit(cpu.B(), 4, true))
			return 8
		},
		0x1e1: func() int { // SET 4,C.
			cpu.SetC(utils.SetBit(cpu.C(), 4, true))
			return 8
		},
		0x1e2: func() int { // SET 4,D.
			cpu.SetD(utils.SetBit(cpu.D(), 4, true))
			return 8
		},
		0x1e3: func() int { // SET 4,E.
			cpu.SetE(utils.SetBit(cpu.E(), 4, true))
			return 8
		},
		0x1e4: func() int { // SET 4,H.
			cpu.SetH(utils.SetBit(cpu.H(), 4, true))
			return 8
		},
		0x1e5: func() int { // SET 4,L.
			cpu.SetL(utils.SetBit(cpu.L(), 4, true))
			return 8
		},
		0x1e6: func() int { // SET 4,(HL).
			mem.Write(cpu.HL(), utils.SetBit(mem.Read(cpu.HL()), 4, true))
			return 16
		},
		0x1e7: func() int { // SET 4,A.
			cpu.SetA(utils.SetBit(cpu.A(), 4, true))
			return 8
		},
		0x1e8: func() int { // SET 5,B.
			cpu.SetB(utils.SetBit(cpu.B(), 5, true))
			return 8
		},
		0x1e9: func() int { // SET 5,C.
			cpu.SetC(utils.SetBit(cpu.C(), 5, true))
			return 8
		},
		0x1ea: func() int { // SET 5,D.
			cpu.SetD(utils.SetBit(cpu.D(), 5, true))
			return 8
		},
		0x1eb: func() int { // SET 5,E.
			cpu.SetE(utils.SetBit(cpu.E(), 5, true))
			return 8
		},
		0x1ec: func() int { // SET 5,H.
			cpu.SetH(utils.SetBit(cpu.H(), 5, true))
			return 8
		},
		0x1ed: func() int { // SET 5,L.
			cpu.SetL(utils.SetBit(cpu.L(), 5, true))
			return 8
		},
		0x1ee: func() int { // SET 5,(HL).
			mem.Write(cpu.HL(), utils.SetBit(mem.Read(cpu.HL()), 5, true))
			return 16
		},
		0x1ef: func() int { // SET 5,A.
			cpu.SetA(utils.SetBit(cpu.A(), 5, true))
			return 8
		},
		0x1f0: func() int { // SET 6,B.
			cpu.SetB(utils.SetBit(cpu.B(), 6, true))
			return 8
		},
		0x1f1: func() int { // SET 6,C.
			cpu.SetC(utils.SetBit(cpu.C(), 6, true))
			return 8
		},
		0x1f2: func() int { // SET 6,D.
			cpu.SetD(utils.SetBit(cpu.D(), 6, true))
			return 8
		},
		0x1f3: func() int { // SET 6,E.
			cpu.SetE(utils.SetBit(cpu.E(), 6, true))
			return 8
		},
		0x1f4: func() int { // SET 6,H.
			cpu.SetH(utils.SetBit(cpu.H(), 6, true))
			return 8
		},
		0x1f5: func() int { // SET 6,L.
			cpu.SetL(utils.SetBit(cpu.L(), 6, true))
			return 8
		},
		0x1f6: func() int { // SET 6,(HL).
			mem.Write(cpu.HL(), utils.SetBit(mem.Read(cpu.HL()), 6, true))
			return 16
		},
		0x1f7: func() int { // SET 6,A.
			cpu.SetA(utils.SetBit(cpu.A(), 6, true))
			return 8
		},
		0x1f8: func() int { // SET 7,B.
			cpu.SetB(utils.SetBit(cpu.B(), 7, true))
			return 8
		},
		0x1f9: func() int { // SET 7,C.
			cpu.SetC(utils.SetBit(cpu.C(), 7, true))
			return 8
		},
		0x1fa: func() int { // SET 7,D.
			cpu.SetD(utils.SetBit(cpu.D(), 7, true))
			return 8
		},
		0x1fb: func() int { // SET 7,E.
			cpu.SetE(utils.SetBit(cpu.E(), 7, true))
			return 8
		},
		0x1fc: func() int { // SET 7,H.
			cpu.SetH(utils.SetBit(cpu.H(), 7, true))
			return 8
		},
		0x1fd: func() int { // SET 7,L.
			cpu.SetL(utils.SetBit(cpu.L(), 7, true))
			return 8
		},
		0x1fe: func() int { // SET 7,(HL).
			mem.Write(cpu.HL(), utils.SetBit(mem.Read(cpu.HL()), 7, true))
			return 16
		},
		0x1ff: func() int { // SET 7,A.
			cpu.SetA(utils.SetBit(cpu.A(), 7, true))
			return 8
		},
	}
}

// Perform an add, update flags, and return the result.
func (c *Cpu) opAdd(a uint8, b uint8, carry bool) uint8 {
	cy := uint8(0)
	if carry {
		cy = 1
	}

	r16 := uint16(a) + uint16(b) + uint16(cy)
	r := uint8(r16)

	c.SetFlagZ(r == 0)
	c.SetFlagN(false)
	c.SetFlagH((a&0xf)+(b&0xf)+(cy&0xf) > 0xf)
	c.SetFlagC(r16 > 0xff)

	return r
}

// Perform a subtract, update flags, and return the result.
func (c *Cpu) opSub(a uint8, b uint8, borrow bool) uint8 {
	bw := uint8(0)
	if borrow {
		bw = 1
	}
	r16 := uint16(a) - uint16(b) - uint16(bw)
	r := uint8(r16)

	c.SetFlagZ(r == 0)
	c.SetFlagN(true)
	c.SetFlagH((a&0xf)-(b&0xf)-(bw&0xf) > 0xf)
	c.SetFlagC(r16 > 0xff)

	return r
}

// Perform a signed add, update flags, and return the result.
func (c *Cpu) opSignedAdd(a uint16, b uint8) uint16 {
	r := uint16(int32(a) + int32(int8(b)))

	// Get the flags from doing an ordinary add.
	c.opAdd(uint8(a), b, false)
	c.SetFlagZ(false)
	c.SetFlagN(false)

	return r
}

// Perform an increment, update flags, and return the result.
func (c *Cpu) opInc(v uint8) uint8 {
	// Do not update the carry flag.
	cy := c.FlagC()
	r := c.opAdd(v, 1, false)
	c.SetFlagC(cy)
	return r
}

// Perform a decrement, update flags, and return the result.
func (c *Cpu) opDec(v uint8) uint8 {
	// Do not update the carry flag.
	cy := c.FlagC()
	r := c.opSub(v, 1, false)
	c.SetFlagC(cy)
	return r
}

// Perform decimal adjust on register A, update flags, and return the result.
func (c *Cpu) opDaa() uint8 {
	a := c.A()

	// Stolen from https://forums.nesdev.com/viewtopic.php?t=15944#p196282.
	if !c.FlagN() {
		if c.FlagC() || a > 0x99 {
			a += 0x60
			c.SetFlagC(true)
		}
		if c.FlagH() || (a&0xf) > 0x09 {
			a += 0x06
		}
	} else {
		if c.FlagC() {
			a -= 0x60
			c.SetFlagC(true)
		}
		if c.FlagH() {
			a -= 0x06
		}
	}

	c.SetFlagZ(a == 0)
	c.SetFlagH(false)

	return a
}

// Perform an AND operation, update flags, and return the result.
func (c *Cpu) opAnd(a uint8, b uint8) uint8 {
	r := a & b

	c.SetFlagZ(r == 0)
	c.SetFlagN(false)
	c.SetFlagH(true)
	c.SetFlagC(false)

	return r
}

// Perform an XOR operation, update flags, and return the result.
func (c *Cpu) opXor(a uint8, b uint8) uint8 {
	r := a ^ b

	c.SetFlagZ(r == 0)
	c.SetFlagN(false)
	c.SetFlagH(false)
	c.SetFlagC(false)

	return r
}

// Perform an OR operation, update flags, and return the result.
func (c *Cpu) opOr(a uint8, b uint8) uint8 {
	r := a | b

	c.SetFlagZ(r == 0)
	c.SetFlagN(false)
	c.SetFlagH(false)
	c.SetFlagC(false)

	return r
}

// Perform a CP operation, update flags, and return the result.
func (c *Cpu) opCp(a uint8, b uint8) uint8 {
	r := a - b

	c.SetFlagZ(r == 0)
	c.SetFlagN(true)
	c.SetFlagH(a&0xf < b&0xf)
	c.SetFlagC(a < b)

	return r
}

// Perform a 16 bit add, update flags, and return the result.
func (c *Cpu) opAdd16(a uint16, b uint16) uint16 {
	r32 := uint32(a) + uint32(b)
	r := uint16(r32)

	c.SetFlagN(false)
	c.SetFlagH(uint32(a&0x0fff) > r32&0x0fff)
	c.SetFlagC(r32 > 0xffff)

	return r
}

// Perform a rotate left, update flags, and return the result.
func (c *Cpu) opRl(a uint8, thruC bool) uint8 {
	r := a << 1
	if !thruC {
		r |= a >> 7
	} else if c.FlagC() {
		r |= 0x1
	}

	c.SetFlagZ(r == 0)
	c.SetFlagN(false)
	c.SetFlagH(false)
	c.SetFlagC(utils.GetBit(a, 7))

	return r
}

// Perform a rotate right, update flags, and return the result.
func (c *Cpu) opRr(a uint8, thruC bool) uint8 {
	r := a >> 1
	if !thruC {
		r |= (a & 0x1) << 7
	} else if c.FlagC() {
		r |= 0x80
	}

	c.SetFlagZ(r == 0)
	c.SetFlagN(false)
	c.SetFlagH(false)
	c.SetFlagC(utils.GetBit(a, 0))

	return r
}

// Perform a relative jump if the given condition is true. Returns how many cycles it took.
func (c *Cpu) opJr(cond bool, r uint8) int {
	if cond {
		c.SetPC(c.PC() + uint16(int8(r)))
		return 12
	}
	return 8
}

// Perform a return if the given condition is true. Returns how many cycles it took.
func (c *Cpu) opRet(cond bool) int {
	if cond {
		c.SetPC(c.PopSP())
		return 20
	}
	return 8
}

// Perform a jump if the given condition is true. Returns how many cycles it took.
func (c *Cpu) opJp(cond bool, a uint16) int {
	if cond {
		c.SetPC(a)
		return 16
	}
	return 12
}

// Perform a call if the given condition is true. Returns how many cycles it took.
func (c *Cpu) opCall(cond bool, a uint16) int {
	if cond {
		c.PushSP(c.PC())
		c.SetPC(a)
		return 24
	}
	return 12
}

// Perform a halt. Handles HALT bug, documented in section 4.10 of
// https://github.com/AntonioND/giibiiadvance/blob/master/docs/TCAGBD.pdf.
func (c *Cpu) opHalt(ime bool, iE uint8, iF uint8) {
	if ime {
		// HALT is executed normally.
		c.SetHalt(true)
	} else {
		if iE&iF == 0 {
			// HALT is executed normally.
			c.SetHalt(true)
		} else {
			// HALT is not executed. Instead, the halt bug is triggered, and the CPU will fail to
			// increment the program counter on the next instruction.
			c.TriggerHaltBug()
		}
	}
}

// Perform a shift left, update flags, and return the result.
func (c *Cpu) opSl(a uint8) uint8 {
	r := a << 1

	c.SetFlagZ(r == 0)
	c.SetFlagN(false)
	c.SetFlagH(false)
	c.SetFlagC(utils.GetBit(a, 7))

	return r
}

// Perform a shift right, update flags, and return the result.
func (c *Cpu) opSr(a uint8, keepMSB bool) uint8 {
	r := a >> 1
	if keepMSB {
		r = utils.SetBit(r, 7, utils.GetBit(a, 7))
	}

	c.SetFlagZ(r == 0)
	c.SetFlagN(false)
	c.SetFlagH(false)
	c.SetFlagC(utils.GetBit(a, 0))

	return r
}

// Perform a swap, update flags, and return the result.
func (c *Cpu) opSwap(a uint8) uint8 {
	r := (a&0x0f)<<4 | (a&0xf0)>>4

	c.SetFlagZ(r == 0)
	c.SetFlagN(false)
	c.SetFlagH(false)
	c.SetFlagC(false)

	return r
}

// Perform a bit test and update flags.
func (c *Cpu) opBit(a uint8, b int) {
	c.SetFlagZ(!utils.GetBit(a, b))
	c.SetFlagN(false)
	c.SetFlagH(true)
}
