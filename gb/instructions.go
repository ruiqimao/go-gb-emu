package gb

// An Instruction returns how many cycles it takes to execute.
type Instruction func() int

func (gb *GameBoy) createInstructionSet() {
	cpu := gb.cpu
	mem := gb.mem

	gb.instructions = [0x200]Instruction{
		// 8 bit loads.
		0x02: func() int { // LD (BC),A.
			mem.Write(cpu.BC(), cpu.A())
			return 8
		},
		0x06: func() int { // LD B,d8.
			cpu.SetB(cpu.IncPc())
			return 8
		},
		0x0a: func() int { // LD A,(BC).
			cpu.SetA(mem.Read(cpu.BC()))
			return 8
		},
		0x0e: func() int { // LD C,d8.
			cpu.SetC(cpu.IncPc())
			return 8
		},
		0x12: func() int { // LD (DE),A.
			mem.Write(cpu.DE(), cpu.A())
			return 8
		},
		0x16: func() int { // LD D,d8.
			cpu.SetD(cpu.IncPc())
			return 8
		},
		0x1a: func() int { // LD A,(DE).
			cpu.SetA(mem.Read(cpu.DE()))
			return 8
		},
		0x1e: func() int { // LD E,d8.
			cpu.SetE(cpu.IncPc())
			return 8
		},
		0x22: func() int { // LD (HL+),A.
			hl := cpu.HL()
			mem.Write(hl, cpu.A())
			cpu.SetHL(hl+1)
			return 8
		},
		0x26: func() int { // LD H,d8.
			cpu.SetH(cpu.IncPc())
			return 8
		},
		0x2a: func() int { // LD A,(HL+).
			hl := cpu.HL()
			cpu.SetA(mem.Read(hl))
			cpu.SetHL(hl+1)
			return 8
		},
		0x2e: func() int { // LD L,d8.
			cpu.SetL(cpu.IncPc())
			return 8
		},
		0x32: func() int { // LD (HL-),A.
			hl := cpu.HL()
			mem.Write(hl, cpu.A())
			cpu.SetHL(hl-1)
			return 8
		},
		0x36: func() int { // LD (HL),d8.
			mem.Write(cpu.HL(), cpu.IncPc())
			return 12
		},
		0x3a: func() int { // LD A,(HL-).
			hl := cpu.HL()
			cpu.SetA(mem.Read(hl))
			cpu.SetHL(hl-1)
			return 8
		},
		0x3e: func() int { // LD A,d8.
			cpu.SetA(cpu.IncPc())
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
			mem.Write(0xff00+uint16(cpu.IncPc()), cpu.A())
			return 12
		},
		0xe2: func() int { // LD ($FF00+C),A.
			mem.Write(0xff00+uint16(cpu.C()), cpu.A())
			return 8
		},
		0xea: func() int { // LD (a16),A.
			mem.Write(cpu.IncPc16(), cpu.A())
			return 16
		},
		0xf0: func() int { // LD A,($FF00+a8).
			cpu.SetA(mem.Read(0xff00+uint16(cpu.IncPc())))
			return 12
		},
		0xf2: func() int { // LD A,(C).
			cpu.SetA(mem.Read(0xff00+uint16(cpu.C())))
			return 8
		},
		0xfa: func() int { // LD A,(a16).
			cpu.SetA(mem.Read(cpu.IncPc16()))
			return 16
		},

		// 16 bit loads.
		0x01: func() int { // LD BC,d16.
			cpu.SetBC(cpu.IncPc16())
			return 12
		},
		0x08: func() int { // LD (a16),SP.
			mem.Write16(cpu.IncPc16(), cpu.Sp())
			return 20
		},
		0x11: func() int { // LD DE,d16.
			cpu.SetDE(cpu.IncPc16())
			return 12
		},
		0x21: func() int { // LD HL,d16.
			cpu.SetHL(cpu.IncPc16())
			return 12
		},
		0x31: func() int { // LD SP,d16.
			cpu.SetSp(cpu.IncPc16())
			return 12
		},
		0xc1: func() int { // POP BC.
			cpu.SetBC(cpu.PopSp())
			return 12
		},
		0xc5: func() int { // PUSH BC.
			cpu.PushSp(cpu.BC())
			return 16
		},
		0xd1: func() int { // POP DE.
			cpu.SetDE(cpu.PopSp())
			return 12
		},
		0xd5: func() int { // PUSH DE.
			cpu.PushSp(cpu.DE())
			return 16
		},
		0xe1: func() int { // POP HL.
			cpu.SetHL(cpu.PopSp())
			return 12
		},
		0xe5: func() int { // PUSH HL.
			cpu.PushSp(cpu.HL())
			return 16
		},
		0xf1: func() int { // POP AF.
			cpu.SetAF(cpu.PopSp())
			return 12
		},
		0xf5: func() int { // PUSH AF.
			cpu.PushSp(cpu.AF())
			return 16
		},
		0xf8: func() int { // LD HL,SP+r8.
			cpu.SetHL(cpu.opSignedAdd(cpu.Sp(), cpu.IncPc()))
			return 12
		},
		0xf9: func() int { // LD SP,HL.
			cpu.SetSp(cpu.HL())
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
			cpu.SetA(cpu.opAdd(cpu.A(), cpu.B(), 0))
			return 4
		},
		0x81: func() int { // ADD A,C.
			cpu.SetA(cpu.opAdd(cpu.A(), cpu.C(), 0))
			return 4
		},
		0x82: func() int { // ADD A,D.
			cpu.SetA(cpu.opAdd(cpu.A(), cpu.D(), 0))
			return 4
		},
		0x83: func() int { // ADD A,E.
			cpu.SetA(cpu.opAdd(cpu.A(), cpu.E(), 0))
			return 4
		},
		0x84: func() int { // ADD A,H.
			cpu.SetA(cpu.opAdd(cpu.A(), cpu.H(), 0))
			return 4
		},
		0x85: func() int { // ADD A,L.
			cpu.SetA(cpu.opAdd(cpu.A(), cpu.L(), 0))
			return 4
		},
		0x86: func() int { // ADD A,(HL).
			cpu.SetA(cpu.opAdd(cpu.A(), mem.Read(cpu.HL()), 0))
			return 8
		},
		0x87: func() int { // ADD A,A.
			cpu.SetA(cpu.opAdd(cpu.A(), cpu.A(), 0))
			return 4
		},
		0x88: func() int { // ADC A,B.
			cpu.SetA(cpu.opAdd(cpu.A(), cpu.B(), 1))
			return 4
		},
		0x89: func() int { // ADC A,C.
			cpu.SetA(cpu.opAdd(cpu.A(), cpu.C(), 1))
			return 4
		},
		0x8a: func() int { // ADC A,D.
			cpu.SetA(cpu.opAdd(cpu.A(), cpu.D(), 1))
			return 4
		},
		0x8b: func() int { // ADC A,E.
			cpu.SetA(cpu.opAdd(cpu.A(), cpu.E(), 1))
			return 4
		},
		0x8c: func() int { // ADC A,H.
			cpu.SetA(cpu.opAdd(cpu.A(), cpu.H(), 1))
			return 4
		},
		0x8d: func() int { // ADC A,L.
			cpu.SetA(cpu.opAdd(cpu.A(), cpu.L(), 1))
			return 4
		},
		0x8e: func() int { // ADC A,(HL).
			cpu.SetA(cpu.opAdd(cpu.A(), mem.Read(cpu.HL()), 1))
			return 8
		},
		0x8f: func() int { // ADC A,A.
			cpu.SetA(cpu.opAdd(cpu.A(), cpu.A(), 1))
			return 4
		},
		0x90: func() int { // SUB B.
			cpu.SetA(cpu.opSub(cpu.A(), cpu.B(), 0))
			return 4
		},
		0x91: func() int { // SUB C.
			cpu.SetA(cpu.opSub(cpu.A(), cpu.C(), 0))
			return 4
		},
		0x92: func() int { // SUB D.
			cpu.SetA(cpu.opSub(cpu.A(), cpu.D(), 0))
			return 4
		},
		0x93: func() int { // SUB E.
			cpu.SetA(cpu.opSub(cpu.A(), cpu.E(), 0))
			return 4
		},
		0x94: func() int { // SUB H.
			cpu.SetA(cpu.opSub(cpu.A(), cpu.H(), 0))
			return 4
		},
		0x95: func() int { // SUB L.
			cpu.SetA(cpu.opSub(cpu.A(), cpu.L(), 0))
			return 4
		},
		0x96: func() int { // SUB (HL).
			cpu.SetA(cpu.opSub(cpu.A(), mem.Read(cpu.HL()), 0))
			return 8
		},
		0x97: func() int { // SUB A.
			cpu.SetA(cpu.opSub(cpu.A(), cpu.A(), 0))
			return 4
		},
		0x98: func() int { // SBC A,B.
			cpu.SetA(cpu.opSub(cpu.A(), cpu.B(), 1))
			return 4
		},
		0x99: func() int { // SBC A,C.
			cpu.SetA(cpu.opSub(cpu.A(), cpu.C(), 1))
			return 4
		},
		0x9a: func() int { // SBC A,D.
			cpu.SetA(cpu.opSub(cpu.A(), cpu.D(), 1))
			return 4
		},
		0x9b: func() int { // SBC A,E.
			cpu.SetA(cpu.opSub(cpu.A(), cpu.E(), 1))
			return 4
		},
		0x9c: func() int { // SBC A,H.
			cpu.SetA(cpu.opSub(cpu.A(), cpu.H(), 1))
			return 4
		},
		0x9d: func() int { // SBC A,L.
			cpu.SetA(cpu.opSub(cpu.A(), cpu.L(), 1))
			return 4
		},
		0x9e: func() int { // SBC A,(HL).
			cpu.SetA(cpu.opSub(cpu.A(), mem.Read(cpu.HL()), 1))
			return 8
		},
		0x9f: func() int { // SBC A,A.
			cpu.SetA(cpu.opSub(cpu.A(), cpu.A(), 1))
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
			cpu.SetA(cpu.opCp(cpu.A(), cpu.B()))
			return 4
		},
		0xb9: func() int { // CP C.
			cpu.SetA(cpu.opCp(cpu.A(), cpu.C()))
			return 4
		},
		0xba: func() int { // CP D.
			cpu.SetA(cpu.opCp(cpu.A(), cpu.D()))
			return 4
		},
		0xbb: func() int { // CP E.
			cpu.SetA(cpu.opCp(cpu.A(), cpu.E()))
			return 4
		},
		0xbc: func() int { // CP H.
			cpu.SetA(cpu.opCp(cpu.A(), cpu.H()))
			return 4
		},
		0xbd: func() int { // CP L.
			cpu.SetA(cpu.opCp(cpu.A(), cpu.L()))
			return 4
		},
		0xbe: func() int { // CP (HL).
			cpu.SetA(cpu.opCp(cpu.A(), mem.Read(cpu.HL())))
			return 8
		},
		0xbf: func() int { // CP A.
			cpu.SetA(cpu.opCp(cpu.A(), cpu.A()))
			return 4
		},
		0xc6: func() int { // ADD A,d8.
			cpu.SetA(cpu.opAdd(cpu.A(), cpu.IncPc(), 0))
			return 8
		},
		0xce: func() int { // ADC A,d8.
			cpu.SetA(cpu.opAdd(cpu.A(), cpu.IncPc(), 1))
			return 8
		},
		0xd6: func() int { // SUB d8.
			cpu.SetA(cpu.opSub(cpu.A(), cpu.IncPc(), 0))
			return 8
		},
		0xde: func() int { // SBC A,d8.
			cpu.SetA(cpu.opSub(cpu.A(), cpu.IncPc(), 1))
			return 8
		},
		0xe6: func() int { // AND d8.
			cpu.SetA(cpu.opAnd(cpu.A(), cpu.IncPc()))
			return 8
		},
		0xee: func() int { // XOR d8.
			cpu.SetA(cpu.opXor(cpu.A(), cpu.IncPc()))
			return 8
		},
		0xf6: func() int { // OR d8.
			cpu.SetA(cpu.opOr(cpu.A(), cpu.IncPc()))
			return 8
		},
		0xfe: func() int { // CP d8.
			cpu.SetA(cpu.opCp(cpu.A(), cpu.IncPc()))
			return 8
		},

		// 16 bit arithmetic.
		0x03: func() int { // INC BC.
			cpu.SetBC(cpu.BC()+1)
			return 8
		},
		0x09: func() int { // ADD HL,BC.
			cpu.SetHL(cpu.opAdd16(cpu.HL(), cpu.BC()))
			return 8
		},
		0x0b: func() int { // DEC BC.
			cpu.SetBC(cpu.BC()-1)
			return 8
		},
		0x13: func() int { // INC DE.
			cpu.SetDE(cpu.DE()+1)
			return 8
		},
		0x19: func() int { // ADD HL,DE.
			cpu.SetHL(cpu.opAdd16(cpu.HL(), cpu.DE()))
			return 8
		},
		0x1b: func() int { // DEC DE.
			cpu.SetDE(cpu.DE()-1)
			return 8
		},
		0x23: func() int { // INC HL.
			cpu.SetHL(cpu.HL()+1)
			return 8
		},
		0x29: func() int { // ADD HL,HL.
			cpu.SetHL(cpu.opAdd16(cpu.HL(), cpu.HL()))
			return 8
		},
		0x2b: func() int { // DEC HL.
			cpu.SetHL(cpu.HL()-1)
			return 8
		},
		0x33: func() int { // INC SP.
			cpu.SetSp(cpu.Sp()+1)
			return 8
		},
		0x39: func() int { // ADD HL,SP.
			cpu.SetHL(cpu.opAdd16(cpu.HL(), cpu.Sp()))
			return 8
		},
		0x3b: func() int { // DEC SP.
			cpu.SetSp(cpu.Sp()-1)
			return 8
		},
		0xe8: func() int { // ADD SP,r8.
			cpu.SetSp(cpu.opSignedAdd(cpu.Sp(), cpu.IncPc()))
			return 16
		},
	}
}

// Perform an add, update flags, and return the result.
func (c *Cpu) opAdd(a uint8, b uint8, cy uint8) uint8 {
	r16 := uint16(a) + uint16(b) + uint16(cy)
	r := uint8(r16)

	c.SetFlagZ(r == 0)
	c.SetFlagN(false)
	c.SetFlagH((a & 0xf) + (b & 0xf) + (cy & 0xf) > 0xf)
	c.SetFlagC(r16 > 0xff)

	return r
}

// Perform a subtract, update flags, and return the result.
func (c *Cpu) opSub(a uint8, b uint8, bw uint8) uint8 {
	r16 := uint16(a) - uint16(b) - uint16(bw)
	r := uint8(r16)

	c.SetFlagZ(r == 0)
	c.SetFlagN(true)
	c.SetFlagH((a & 0xf) - (b & 0xf) - (bw & 0xf) > 0xf)
	c.SetFlagC(r16 > 0xff)

	return r
}

// Perform a signed add, update flags, and return the result.
func (c *Cpu) opSignedAdd(a uint16, b uint8) uint16 {
	r := uint16(int32(a) + int32(int8(b)))

	// Get the flags from doing an ordinary add.
	c.opAdd(uint8(a), b, 0)
	c.SetFlagZ(false)
	c.SetFlagN(false)

	return r
}

// Perform an increment, update flags, and return the result.
func (c *Cpu) opInc(v uint8) uint8 {
	// Do not update the carry flag.
	cy := c.FlagC()
	r := c.opAdd(v, 1, 0)
	c.SetFlagC(cy)
	return r
}

// Perform a decrement, update flags, and return the result.
func (c *Cpu) opDec(v uint8) uint8 {
	// Do not update the carry flag.
	cy := c.FlagC()
	r := c.opSub(v, 1, 0)
	c.SetFlagC(cy)
	return r
}

// Perform decimal adjust on register A, update flags, and return the result.
func (c *Cpu) opDaa() uint8 {
	a := c.A()

	// Stolen from https://forums.nesdev.com/viewtopic.php?t=15944#p196282.
	if !c.FlagN() {
		if c.FlagH() || (a & 0xf) > 9 {
			a += 0x06
		}
		if c.FlagC() || a > 0x99 {
			a += 0x60
			c.SetFlagC(true)
		}
	} else {
		if c.FlagC() {
			a -= 0x60
		}
		if c.FlagH() {
			a -= 0x6
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
	c.SetFlagH(a & 0xf < b & 0xf)
	c.SetFlagC(a < b)

	return r
}

// Perform a 16 bit add, update flags, and return the result.
func (c *Cpu) opAdd16(a uint16, b uint16) uint16 {
	r32 := uint32(a) + uint32(b)
	r := uint16(r32)

	c.SetFlagN(false)
	c.SetFlagH(uint32(a & 0x0fff) > r32 & 0x0fff)
	c.SetFlagC(r32 > 0xffff)

	return r
}
