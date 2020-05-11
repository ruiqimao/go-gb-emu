package gb

// An Instruction returns how many cycles it takes to execute.
type Instruction func() int

func (gb *GameBoy) createInstructionSet() {
	cpu := gb.cpu
	mem := gb.mem

	gb.instructions = [0x200]Instruction{
		// 8 bit loads.
		0x02: func() int { // LD (BC),A.
			mem.Write(cpu.Get16(RegBC), cpu.Get(RegA))
			return 8
		},
		0x06: func() int { // LD B,d8.
			cpu.Set(RegB, cpu.IncPc())
			return 8
		},
		0x0a: func() int { // LD A,(BC).
			cpu.Set(RegA, mem.Read(cpu.Get16(RegBC)))
			return 8
		},
		0x0e: func() int { // LD C,d8.
			cpu.Set(RegC, cpu.IncPc())
			return 8
		},
		0x12: func() int { // LD (DE),A.
			mem.Write(cpu.Get16(RegDE), cpu.Get(RegA))
			return 8
		},
		0x16: func() int { // LD D,d8.
			cpu.Set(RegD, cpu.IncPc())
			return 8
		},
		0x1a: func() int { // LD A,(DE).
			cpu.Set(RegA, mem.Read(cpu.Get16(RegDE)))
			return 8
		},
		0x1e: func() int { // LD E,d8.
			cpu.Set(RegE, cpu.IncPc())
			return 8
		},
		0x22: func() int { // LD (HL+),A.
			hl := cpu.Get16(RegHL)
			mem.Write(hl, cpu.Get(RegA))
			cpu.Set16(RegHL, hl+1)
			return 8
		},
		0x26: func() int { // LD H,d8.
			cpu.Set(RegH, cpu.IncPc())
			return 8
		},
		0x2a: func() int { // LD A,(HL+).
			hl := cpu.Get16(RegHL)
			cpu.Set(RegA, mem.Read(hl))
			cpu.Set16(RegHL, hl+1)
			return 8
		},
		0x2e: func() int { // LD L,d8.
			cpu.Set(RegL, cpu.IncPc())
			return 8
		},
		0x32: func() int { // LD (HL-),A.
			hl := cpu.Get16(RegHL)
			mem.Write(hl, cpu.Get(RegA))
			cpu.Set16(RegHL, hl-1)
			return 8
		},
		0x36: func() int { // LD (HL),d8.
			mem.Write(cpu.Get16(RegHL), cpu.IncPc())
			return 12
		},
		0x3a: func() int { // LD A,(HL-).
			hl := cpu.Get16(RegHL)
			cpu.Set(RegA, mem.Read(hl))
			cpu.Set16(RegH, hl-1)
			return 8
		},
		0x3e: func() int { // LD A,d8.
			cpu.Set(RegA, cpu.IncPc())
			return 8
		},
		0x40: func() int { // LD B,B.
			cpu.Set(RegB, cpu.Get(RegB))
			return 4
		},
		0x41: func() int { // LD B,C.
			cpu.Set(RegB, cpu.Get(RegC))
			return 4
		},
		0x42: func() int { // LD B,D.
			cpu.Set(RegB, cpu.Get(RegD))
			return 4
		},
		0x43: func() int { // LD B,E.
			cpu.Set(RegB, cpu.Get(RegE))
			return 4
		},
		0x44: func() int { // LD B,H.
			cpu.Set(RegB, cpu.Get(RegH))
			return 4
		},
		0x45: func() int { // LD B,L.
			cpu.Set(RegB, cpu.Get(RegL))
			return 4
		},
		0x46: func() int { // LD B,(HL).
			cpu.Set(RegB, mem.Read(cpu.Get16(RegHL)))
			return 8
		},
		0x47: func() int { // LD B,A.
			cpu.Set(RegB, cpu.Get(RegA))
			return 4
		},
		0x48: func() int { // LD C,B.
			cpu.Set(RegC, cpu.Get(RegB))
			return 4
		},
		0x49: func() int { // LD C,C.
			cpu.Set(RegC, cpu.Get(RegC))
			return 4
		},
		0x4a: func() int { // LD C,D.
			cpu.Set(RegC, cpu.Get(RegD))
			return 4
		},
		0x4b: func() int { // LD C,E.
			cpu.Set(RegC, cpu.Get(RegE))
			return 4
		},
		0x4c: func() int { // LD C,H.
			cpu.Set(RegC, cpu.Get(RegH))
			return 4
		},
		0x4d: func() int { // LD C,L.
			cpu.Set(RegC, cpu.Get(RegL))
			return 4
		},
		0x4e: func() int { // LD C,(HL).
			cpu.Set(RegC, mem.Read(cpu.Get16(RegHL)))
			return 8
		},
		0x4f: func() int { // LD C,A.
			cpu.Set(RegC, cpu.Get(RegA))
			return 4
		},
		0x50: func() int { // LD D,B.
			cpu.Set(RegD, cpu.Get(RegB))
			return 4
		},
		0x51: func() int { // LD D,C.
			cpu.Set(RegD, cpu.Get(RegC))
			return 4
		},
		0x52: func() int { // LD D,D.
			cpu.Set(RegD, cpu.Get(RegD))
			return 4
		},
		0x53: func() int { // LD D,E.
			cpu.Set(RegD, cpu.Get(RegE))
			return 4
		},
		0x54: func() int { // LD D,H.
			cpu.Set(RegD, cpu.Get(RegH))
			return 4
		},
		0x55: func() int { // LD D,L.
			cpu.Set(RegD, cpu.Get(RegL))
			return 4
		},
		0x56: func() int { // LD D,(HL).
			cpu.Set(RegD, mem.Read(cpu.Get16(RegHL)))
			return 8
		},
		0x57: func() int { // LD D,A.
			cpu.Set(RegD, cpu.Get(RegA))
			return 4
		},
		0x58: func() int { // LD E,B.
			cpu.Set(RegE, cpu.Get(RegB))
			return 4
		},
		0x59: func() int { // LD E,C.
			cpu.Set(RegE, cpu.Get(RegC))
			return 4
		},
		0x5a: func() int { // LD E,D.
			cpu.Set(RegE, cpu.Get(RegD))
			return 4
		},
		0x5b: func() int { // LD E,E.
			cpu.Set(RegE, cpu.Get(RegE))
			return 4
		},
		0x5c: func() int { // LD E,H.
			cpu.Set(RegE, cpu.Get(RegH))
			return 4
		},
		0x5d: func() int { // LD E,L.
			cpu.Set(RegE, cpu.Get(RegL))
			return 4
		},
		0x5e: func() int { // LD E,(HL).
			cpu.Set(RegE, mem.Read(cpu.Get16(RegHL)))
			return 8
		},
		0x5f: func() int { // LD E,A.
			cpu.Set(RegE, cpu.Get(RegA))
			return 4
		},
		0x60: func() int { // LD H,B.
			cpu.Set(RegH, cpu.Get(RegB))
			return 4
		},
		0x61: func() int { // LD H,C.
			cpu.Set(RegH, cpu.Get(RegC))
			return 4
		},
		0x62: func() int { // LD H,D.
			cpu.Set(RegH, cpu.Get(RegD))
			return 4
		},
		0x63: func() int { // LD H,E.
			cpu.Set(RegH, cpu.Get(RegE))
			return 4
		},
		0x64: func() int { // LD H,H.
			cpu.Set(RegH, cpu.Get(RegH))
			return 4
		},
		0x65: func() int { // LD H,L.
			cpu.Set(RegH, cpu.Get(RegL))
			return 4
		},
		0x66: func() int { // LD H,(HL).
			cpu.Set(RegH, mem.Read(cpu.Get16(RegHL)))
			return 8
		},
		0x67: func() int { // LD H,A.
			cpu.Set(RegH, cpu.Get(RegA))
			return 4
		},
		0x68: func() int { // LD L,B.
			cpu.Set(RegL, cpu.Get(RegB))
			return 4
		},
		0x69: func() int { // LD L,C.
			cpu.Set(RegL, cpu.Get(RegC))
			return 4
		},
		0x6a: func() int { // LD L,D.
			cpu.Set(RegL, cpu.Get(RegD))
			return 4
		},
		0x6b: func() int { // LD L,E.
			cpu.Set(RegL, cpu.Get(RegE))
			return 4
		},
		0x6c: func() int { // LD L,H.
			cpu.Set(RegL, cpu.Get(RegH))
			return 4
		},
		0x6d: func() int { // LD L,L.
			cpu.Set(RegL, cpu.Get(RegL))
			return 4
		},
		0x6e: func() int { // LD L,(HL).
			cpu.Set(RegL, mem.Read(cpu.Get16(RegHL)))
			return 8
		},
		0x6f: func() int { // LD L,A.
			cpu.Set(RegL, cpu.Get(RegA))
			return 4
		},
		0x70: func() int { // LD (HL),B.
			mem.Write(cpu.Get16(RegHL), cpu.Get(RegB))
			return 8
		},
		0x71: func() int { // LD (HL),C.
			mem.Write(cpu.Get16(RegHL), cpu.Get(RegC))
			return 8
		},
		0x72: func() int { // LD (HL),D.
			mem.Write(cpu.Get16(RegHL), cpu.Get(RegD))
			return 8
		},
		0x73: func() int { // LD (HL),E.
			mem.Write(cpu.Get16(RegHL), cpu.Get(RegE))
			return 8
		},
		0x74: func() int { // LD (HL),H.
			mem.Write(cpu.Get16(RegHL), cpu.Get(RegH))
			return 8
		},
		0x75: func() int { // LD (HL),L.
			mem.Write(cpu.Get16(RegHL), cpu.Get(RegL))
			return 8
		},
		0x77: func() int { // LD (HL),A.
			mem.Write(cpu.Get16(RegHL), cpu.Get(RegA))
			return 8
		},
		0x78: func() int { // LD A,B.
			cpu.Set(RegA, cpu.Get(RegB))
			return 4
		},
		0x79: func() int { // LD A,C.
			cpu.Set(RegA, cpu.Get(RegC))
			return 4
		},
		0x7a: func() int { // LD A,D.
			cpu.Set(RegA, cpu.Get(RegD))
			return 4
		},
		0x7b: func() int { // LD A,E.
			cpu.Set(RegA, cpu.Get(RegE))
			return 4
		},
		0x7c: func() int { // LD A,H.
			cpu.Set(RegA, cpu.Get(RegH))
			return 4
		},
		0x7d: func() int { // LD A,L.
			cpu.Set(RegA, cpu.Get(RegL))
			return 4
		},
		0x7e: func() int { // LD A,(HL).
			cpu.Set(RegA, mem.Read(cpu.Get16(RegHL)))
			return 8
		},
		0x7f: func() int { // LD A,A.
			cpu.Set(RegA, cpu.Get(RegA))
			return 4
		},
		0xe0: func() int { // LD ($FF00+a8),A.
			mem.Write(0xff00+uint16(cpu.IncPc()), cpu.Get(RegA))
			return 12
		},
		0xe2: func() int { // LD ($FF00+C),A.
			mem.Write(0xff00+uint16(cpu.Get(RegC)), cpu.Get(RegA))
			return 8
		},
		0xea: func() int { // LD (a16),A.
			mem.Write(cpu.IncPc16(), cpu.Get(RegA))
			return 16
		},
		0xf0: func() int { // LD A,($FF00+a8).
			cpu.Set(RegA, mem.Read(0xff00+uint16(cpu.IncPc())))
			return 12
		},
		0xf2: func() int { // LD A,(C).
			cpu.Set(RegA, mem.Read(0xff00+uint16(cpu.Get(RegC))))
			return 8
		},
		0xfa: func() int { // LD A,(a16).
			cpu.Set(RegA, mem.Read(cpu.IncPc16()))
			return 16
		},

		// 16 bit loads.
		0x01: func() int { // LD BC,d16.
			cpu.Set16(RegBC, cpu.IncPc16())
			return 12
		},
		0x08: func() int { // LD (a16),SP.
			mem.Write16(cpu.IncPc16(), cpu.Sp())
			return 20
		},
		0x11: func() int { // LD DE,d16.
			cpu.Set16(RegDE, cpu.IncPc16())
			return 12
		},
		0x21: func() int { // LD HL,d16.
			cpu.Set16(RegHL, cpu.IncPc16())
			return 12
		},
		0x31: func() int { // LD SP,d16.
			cpu.SetSp(cpu.IncPc16())
			return 12
		},
		0xc1: func() int { // POP BC.
			cpu.Set16(RegBC, cpu.PopSp())
			return 12
		},
		0xc5: func() int { // PUSH BC.
			cpu.PushSp(cpu.Get16(RegBC))
			return 16
		},
		0xd1: func() int { // POP DE.
			cpu.Set16(RegDE, cpu.PopSp())
			return 12
		},
		0xd5: func() int { // PUSH DE.
			cpu.PushSp(cpu.Get16(RegDE))
			return 16
		},
		0xe1: func() int { // POP HL.
			cpu.Set16(RegHL, cpu.PopSp())
			return 12
		},
		0xe5: func() int { // PUSH HL.
			cpu.PushSp(cpu.Get16(RegHL))
			return 16
		},
		0xf1: func() int { // POP AF.
			cpu.Set16(RegAF, cpu.PopSp())
			return 12
		},
		0xf5: func() int { // PUSH AF.
			cpu.PushSp(cpu.Get16(RegAF))
			return 16
		},
		0xf8: func() int { // LD HL,SP+r8.
			cpu.Set16(RegHL, cpu.opSignedAdd(cpu.Sp(), cpu.IncPc()))
			return 12
		},
		0xf9: func() int { // LD SP,HL.
			cpu.SetSp(cpu.Get16(RegHL))
			return 8
		},

		// 8 bit arithmetic.
		0x04: func() int { // INC B.
			cpu.Set(RegB, cpu.opInc(cpu.Get(RegB)))
			return 4
		},
		0x05: func() int { // DEC B.
			cpu.Set(RegB, cpu.opDec(cpu.Get(RegB)))
			return 4
		},
		0x0c: func() int { // INC C.
			cpu.Set(RegC, cpu.opInc(cpu.Get(RegC)))
			return 4
		},
		0x0d: func() int { // DEC C.
			cpu.Set(RegC, cpu.opDec(cpu.Get(RegC)))
			return 4
		},
		0x14: func() int { // INC D.
			cpu.Set(RegD, cpu.opInc(cpu.Get(RegD)))
			return 4
		},
		0x15: func() int { // DEC D.
			cpu.Set(RegD, cpu.opDec(cpu.Get(RegD)))
			return 4
		},
		0x1c: func() int { // INC E.
			cpu.Set(RegE, cpu.opInc(cpu.Get(RegE)))
			return 4
		},
		0x1d: func() int { // DEC E.
			cpu.Set(RegE, cpu.opDec(cpu.Get(RegE)))
			return 4
		},
		0x24: func() int { // INC H.
			cpu.Set(RegH, cpu.opInc(cpu.Get(RegH)))
			return 4
		},
		0x25: func() int { // DEC H.
			cpu.Set(RegH, cpu.opDec(cpu.Get(RegH)))
			return 4
		},
		0x27: func() int { // DAA.
			cpu.Set(RegA, cpu.opDaa())
			return 4
		},
		0x2c: func() int { // INC L.
			cpu.Set(RegL, cpu.opInc(cpu.Get(RegL)))
			return 4
		},
		0x2d: func() int { // DEC L.
			cpu.Set(RegL, cpu.opDec(cpu.Get(RegL)))
			return 4
		},
		0x2f: func() int { // CPL.
			cpu.Set(RegA, ^cpu.Get(RegA))
			cpu.SetFlag(FlagN, true)
			cpu.SetFlag(FlagH, true)
			return 4
		},
		0x34: func() int { // INC (HL).
			hl := cpu.Get16(RegHL)
			mem.Write(hl, cpu.opInc(mem.Read(hl)))
			return 12
		},
		0x35: func() int { // DEC (HL).
			hl := cpu.Get16(RegHL)
			mem.Write(hl, cpu.opDec(mem.Read(hl)))
			return 12
		},
		0x37: func() int { // SCF.
			cpu.SetFlag(FlagN, false)
			cpu.SetFlag(FlagH, false)
			cpu.SetFlag(FlagC, true)
			return 4
		},
		0x3c: func() int { // INC A.
			cpu.Set(RegA, cpu.opInc(cpu.Get(RegA)))
			return 4
		},
		0x3d: func() int { // DEC A.
			cpu.Set(RegA, cpu.opDec(cpu.Get(RegA)))
			return 4
		},
		0x3f: func() int { // CCF.
			cpu.SetFlag(FlagN, false)
			cpu.SetFlag(FlagH, false)
			cpu.SetFlag(FlagC, !cpu.GetFlag(FlagC))
			return 4
		},
		0x80: func() int { // ADD A,B.
			cpu.Set(RegA, cpu.opAdd(cpu.Get(RegA), cpu.Get(RegB), 0))
			return 4
		},
		0x81: func() int { // ADD A,C.
			cpu.Set(RegA, cpu.opAdd(cpu.Get(RegA), cpu.Get(RegC), 0))
			return 4
		},
		0x82: func() int { // ADD A,D.
			cpu.Set(RegA, cpu.opAdd(cpu.Get(RegA), cpu.Get(RegD), 0))
			return 4
		},
		0x83: func() int { // ADD A,E.
			cpu.Set(RegA, cpu.opAdd(cpu.Get(RegA), cpu.Get(RegE), 0))
			return 4
		},
		0x84: func() int { // ADD A,H.
			cpu.Set(RegA, cpu.opAdd(cpu.Get(RegA), cpu.Get(RegH), 0))
			return 4
		},
		0x85: func() int { // ADD A,L.
			cpu.Set(RegA, cpu.opAdd(cpu.Get(RegA), cpu.Get(RegL), 0))
			return 4
		},
		0x86: func() int { // ADD A,(HL).
			cpu.Set(RegA, cpu.opAdd(cpu.Get(RegA), mem.Read(cpu.Get16(RegHL)), 0))
			return 8
		},
		0x87: func() int { // ADD A,A.
			cpu.Set(RegA, cpu.opAdd(cpu.Get(RegA), cpu.Get(RegA), 0))
			return 4
		},
		0x88: func() int { // ADC A,B.
			cpu.Set(RegA, cpu.opAdd(cpu.Get(RegA), cpu.Get(RegB), 1))
			return 4
		},
		0x89: func() int { // ADC A,C.
			cpu.Set(RegA, cpu.opAdd(cpu.Get(RegA), cpu.Get(RegC), 1))
			return 4
		},
		0x8a: func() int { // ADC A,D.
			cpu.Set(RegA, cpu.opAdd(cpu.Get(RegA), cpu.Get(RegD), 1))
			return 4
		},
		0x8b: func() int { // ADC A,E.
			cpu.Set(RegA, cpu.opAdd(cpu.Get(RegA), cpu.Get(RegE), 1))
			return 4
		},
		0x8c: func() int { // ADC A,H.
			cpu.Set(RegA, cpu.opAdd(cpu.Get(RegA), cpu.Get(RegH), 1))
			return 4
		},
		0x8d: func() int { // ADC A,L.
			cpu.Set(RegA, cpu.opAdd(cpu.Get(RegA), cpu.Get(RegL), 1))
			return 4
		},
		0x8e: func() int { // ADC A,(HL).
			cpu.Set(RegA, cpu.opAdd(cpu.Get(RegA), mem.Read(cpu.Get16(RegHL)), 1))
			return 8
		},
		0x8f: func() int { // ADC A,A.
			cpu.Set(RegA, cpu.opAdd(cpu.Get(RegA), cpu.Get(RegA), 1))
			return 4
		},
		0x90: func() int { // SUB B.
			cpu.Set(RegA, cpu.opSub(cpu.Get(RegA), cpu.Get(RegB), 0))
			return 4
		},
		0x91: func() int { // SUB C.
			cpu.Set(RegA, cpu.opSub(cpu.Get(RegA), cpu.Get(RegC), 0))
			return 4
		},
		0x92: func() int { // SUB D.
			cpu.Set(RegA, cpu.opSub(cpu.Get(RegA), cpu.Get(RegD), 0))
			return 4
		},
		0x93: func() int { // SUB E.
			cpu.Set(RegA, cpu.opSub(cpu.Get(RegA), cpu.Get(RegE), 0))
			return 4
		},
		0x94: func() int { // SUB H.
			cpu.Set(RegA, cpu.opSub(cpu.Get(RegA), cpu.Get(RegH), 0))
			return 4
		},
		0x95: func() int { // SUB L.
			cpu.Set(RegA, cpu.opSub(cpu.Get(RegA), cpu.Get(RegL), 0))
			return 4
		},
		0x96: func() int { // SUB (HL).
			cpu.Set(RegA, cpu.opSub(cpu.Get(RegA), mem.Read(cpu.Get16(RegHL)), 0))
			return 8
		},
		0x97: func() int { // SUB A.
			cpu.Set(RegA, cpu.opSub(cpu.Get(RegA), cpu.Get(RegA), 0))
			return 4
		},
		0x98: func() int { // SBC A,B.
			cpu.Set(RegA, cpu.opSub(cpu.Get(RegA), cpu.Get(RegB), 1))
			return 4
		},
		0x99: func() int { // SBC A,C.
			cpu.Set(RegA, cpu.opSub(cpu.Get(RegA), cpu.Get(RegC), 1))
			return 4
		},
		0x9a: func() int { // SBC A,D.
			cpu.Set(RegA, cpu.opSub(cpu.Get(RegA), cpu.Get(RegD), 1))
			return 4
		},
		0x9b: func() int { // SBC A,E.
			cpu.Set(RegA, cpu.opSub(cpu.Get(RegA), cpu.Get(RegE), 1))
			return 4
		},
		0x9c: func() int { // SBC A,H.
			cpu.Set(RegA, cpu.opSub(cpu.Get(RegA), cpu.Get(RegH), 1))
			return 4
		},
		0x9d: func() int { // SBC A,L.
			cpu.Set(RegA, cpu.opSub(cpu.Get(RegA), cpu.Get(RegL), 1))
			return 4
		},
		0x9e: func() int { // SBC A,(HL).
			cpu.Set(RegA, cpu.opSub(cpu.Get(RegA), mem.Read(cpu.Get16(RegHL)), 1))
			return 8
		},
		0x9f: func() int { // SBC A,A.
			cpu.Set(RegA, cpu.opSub(cpu.Get(RegA), cpu.Get(RegA), 1))
			return 4
		},
		0xa0: func() int { // AND B.
			cpu.Set(RegA, cpu.opAnd(cpu.Get(RegA), cpu.Get(RegB)))
			return 4
		},
		0xa1: func() int { // AND C.
			cpu.Set(RegA, cpu.opAnd(cpu.Get(RegA), cpu.Get(RegC)))
			return 4
		},
		0xa2: func() int { // AND D.
			cpu.Set(RegA, cpu.opAnd(cpu.Get(RegA), cpu.Get(RegD)))
			return 4
		},
		0xa3: func() int { // AND E.
			cpu.Set(RegA, cpu.opAnd(cpu.Get(RegA), cpu.Get(RegE)))
			return 4
		},
		0xa4: func() int { // AND H.
			cpu.Set(RegA, cpu.opAnd(cpu.Get(RegA), cpu.Get(RegH)))
			return 4
		},
		0xa5: func() int { // AND L.
			cpu.Set(RegA, cpu.opAnd(cpu.Get(RegA), cpu.Get(RegL)))
			return 4
		},
		0xa6: func() int { // AND (HL).
			cpu.Set(RegA, cpu.opAnd(cpu.Get(RegA), mem.Read(cpu.Get16(RegHL))))
			return 8
		},
		0xa7: func() int { // AND A.
			cpu.Set(RegA, cpu.opAnd(cpu.Get(RegA), cpu.Get(RegA)))
			return 4
		},
		0xa8: func() int { // XOR B.
			cpu.Set(RegA, cpu.opXor(cpu.Get(RegA), cpu.Get(RegB)))
			return 4
		},
		0xa9: func() int { // XOR C.
			cpu.Set(RegA, cpu.opXor(cpu.Get(RegA), cpu.Get(RegC)))
			return 4
		},
		0xaa: func() int { // XOR D.
			cpu.Set(RegA, cpu.opXor(cpu.Get(RegA), cpu.Get(RegD)))
			return 4
		},
		0xab: func() int { // XOR E.
			cpu.Set(RegA, cpu.opXor(cpu.Get(RegA), cpu.Get(RegE)))
			return 4
		},
		0xac: func() int { // XOR H.
			cpu.Set(RegA, cpu.opXor(cpu.Get(RegA), cpu.Get(RegH)))
			return 4
		},
		0xad: func() int { // XOR L.
			cpu.Set(RegA, cpu.opXor(cpu.Get(RegA), cpu.Get(RegL)))
			return 4
		},
		0xae: func() int { // XOR (HL).
			cpu.Set(RegA, cpu.opXor(cpu.Get(RegA), mem.Read(cpu.Get16(RegHL))))
			return 8
		},
		0xaf: func() int { // XOR A.
			cpu.Set(RegA, cpu.opXor(cpu.Get(RegA), cpu.Get(RegA)))
			return 4
		},
		0xb0: func() int { // OR B.
			cpu.Set(RegA, cpu.opOr(cpu.Get(RegA), cpu.Get(RegB)))
			return 4
		},
		0xb1: func() int { // OR C.
			cpu.Set(RegA, cpu.opOr(cpu.Get(RegA), cpu.Get(RegC)))
			return 4
		},
		0xb2: func() int { // OR D.
			cpu.Set(RegA, cpu.opOr(cpu.Get(RegA), cpu.Get(RegD)))
			return 4
		},
		0xb3: func() int { // OR E.
			cpu.Set(RegA, cpu.opOr(cpu.Get(RegA), cpu.Get(RegE)))
			return 4
		},
		0xb4: func() int { // OR H.
			cpu.Set(RegA, cpu.opOr(cpu.Get(RegA), cpu.Get(RegH)))
			return 4
		},
		0xb5: func() int { // OR L.
			cpu.Set(RegA, cpu.opOr(cpu.Get(RegA), cpu.Get(RegL)))
			return 4
		},
		0xb6: func() int { // OR (HL).
			cpu.Set(RegA, cpu.opOr(cpu.Get(RegA), mem.Read(cpu.Get16(RegHL))))
			return 8
		},
		0xb7: func() int { // OR A.
			cpu.Set(RegA, cpu.opOr(cpu.Get(RegA), cpu.Get(RegA)))
			return 4
		},
		0xb8: func() int { // CP B.
			cpu.Set(RegA, cpu.opCp(cpu.Get(RegA), cpu.Get(RegB)))
			return 4
		},
		0xb9: func() int { // CP C.
			cpu.Set(RegA, cpu.opCp(cpu.Get(RegA), cpu.Get(RegC)))
			return 4
		},
		0xba: func() int { // CP D.
			cpu.Set(RegA, cpu.opCp(cpu.Get(RegA), cpu.Get(RegD)))
			return 4
		},
		0xbb: func() int { // CP E.
			cpu.Set(RegA, cpu.opCp(cpu.Get(RegA), cpu.Get(RegE)))
			return 4
		},
		0xbc: func() int { // CP H.
			cpu.Set(RegA, cpu.opCp(cpu.Get(RegA), cpu.Get(RegH)))
			return 4
		},
		0xbd: func() int { // CP L.
			cpu.Set(RegA, cpu.opCp(cpu.Get(RegA), cpu.Get(RegL)))
			return 4
		},
		0xbe: func() int { // CP (HL).
			cpu.Set(RegA, cpu.opCp(cpu.Get(RegA), mem.Read(cpu.Get16(RegHL))))
			return 8
		},
		0xbf: func() int { // CP A.
			cpu.Set(RegA, cpu.opCp(cpu.Get(RegA), cpu.Get(RegA)))
			return 4
		},
		0xc6: func() int { // ADD A,d8.
			cpu.Set(RegA, cpu.opAdd(cpu.Get(RegA), cpu.IncPc(), 0))
			return 8
		},
		0xce: func() int { // ADC A,d8.
			cpu.Set(RegA, cpu.opAdd(cpu.Get(RegA), cpu.IncPc(), 1))
			return 8
		},
		0xd6: func() int { // SUB d8.
			cpu.Set(RegA, cpu.opSub(cpu.Get(RegA), cpu.IncPc(), 0))
			return 8
		},
		0xde: func() int { // SBC A,d8.
			cpu.Set(RegA, cpu.opSub(cpu.Get(RegA), cpu.IncPc(), 1))
			return 8
		},
		0xe6: func() int { // AND d8.
			cpu.Set(RegA, cpu.opAnd(cpu.Get(RegA), cpu.IncPc()))
			return 8
		},
		0xee: func() int { // XOR d8.
			cpu.Set(RegA, cpu.opXor(cpu.Get(RegA), cpu.IncPc()))
			return 8
		},
		0xf6: func() int { // OR d8.
			cpu.Set(RegA, cpu.opOr(cpu.Get(RegA), cpu.IncPc()))
			return 8
		},
		0xfe: func() int { // CP d8.
			cpu.Set(RegA, cpu.opCp(cpu.Get(RegA), cpu.IncPc()))
			return 8
		},

		// 16 bit arithmetic.
		0x03: func() int { // INC BC.
			cpu.Set16(RegBC, cpu.Get16(RegBC)+1)
			return 8
		},
		0x09: func() int { // ADD HL,BC.
			cpu.Set16(RegHL, cpu.opAdd16(cpu.Get16(RegHL), cpu.Get16(RegBC)))
			return 8
		},
		0x0b: func() int { // DEC BC.
			cpu.Set16(RegBC, cpu.Get16(RegBC)-1)
			return 8
		},
		0x13: func() int { // INC DE.
			cpu.Set16(RegDE, cpu.Get16(RegDE)+1)
			return 8
		},
		0x19: func() int { // ADD HL,DE.
			cpu.Set16(RegHL, cpu.opAdd16(cpu.Get16(RegHL), cpu.Get16(RegDE)))
			return 8
		},
		0x1b: func() int { // DEC DE.
			cpu.Set16(RegDE, cpu.Get16(RegDE)-1)
			return 8
		},
		0x23: func() int { // INC HL.
			cpu.Set16(RegHL, cpu.Get16(RegHL)+1)
			return 8
		},
		0x29: func() int { // ADD HL,HL.
			cpu.Set16(RegHL, cpu.opAdd16(cpu.Get16(RegHL), cpu.Get16(RegHL)))
			return 8
		},
		0x2b: func() int { // DEC HL.
			cpu.Set16(RegHL, cpu.Get16(RegHL)-1)
			return 8
		},
		0x33: func() int { // INC SP.
			cpu.SetSp(cpu.Sp()+1)
			return 8
		},
		0x39: func() int { // ADD HL,SP.
			cpu.Set16(RegHL, cpu.opAdd16(cpu.Get16(RegHL), cpu.Sp()))
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

	c.SetFlag(FlagZ, r == 0)
	c.SetFlag(FlagN, false)
	c.SetFlag(FlagH, (a & 0xf) + (b & 0xf) + (cy & 0xf) > 0xf)
	c.SetFlag(FlagC, r16 > 0xff)

	return r
}

// Perform a subtract, update flags, and return the result.
func (c *Cpu) opSub(a uint8, b uint8, bw uint8) uint8 {
	r16 := uint16(a) - uint16(b) - uint16(bw)
	r := uint8(r16)

	c.SetFlag(FlagZ, r == 0)
	c.SetFlag(FlagN, true)
	c.SetFlag(FlagH, (a & 0xf) - (b & 0xf) - (bw & 0xf) > 0xf)
	c.SetFlag(FlagC, r16 > 0xff)

	return r
}

// Perform a signed add, update flags, and return the result.
func (c *Cpu) opSignedAdd(a uint16, b uint8) uint16 {
	r := uint16(int32(a) + int32(int8(b)))

	// Get the flags from doing an ordinary add.
	c.opAdd(uint8(a), b, 0)
	c.SetFlag(FlagZ, false)
	c.SetFlag(FlagN, false)

	return r
}

// Perform an increment, update flags, and return the result.
func (c *Cpu) opInc(v uint8) uint8 {
	// Do not update the carry flag.
	cy := c.GetFlag(FlagC)
	r := c.opAdd(v, 1, 0)
	c.SetFlag(FlagC, cy)
	return r
}

// Perform a decrement, update flags, and return the result.
func (c *Cpu) opDec(v uint8) uint8 {
	// Do not update the carry flag.
	cy := c.GetFlag(FlagC)
	r := c.opSub(v, 1, 0)
	c.SetFlag(FlagC, cy)
	return r
}

// Perform decimal adjust on register A, update flags, and return the result.
func (c *Cpu) opDaa() uint8 {
	a := c.Get(RegA)

	// Stolen from https://forums.nesdev.com/viewtopic.php?t=15944#p196282.
	if !c.GetFlag(FlagN) {
		if c.GetFlag(FlagH) || (a & 0xf) > 9 {
			a += 0x06
		}
		if c.GetFlag(FlagC) || a > 0x99 {
			a += 0x60
			c.SetFlag(FlagC, true)
		}
	} else {
		if c.GetFlag(FlagC) {
			a -= 0x60
		}
		if c.GetFlag(FlagH) {
			a -= 0x6
		}
	}

	c.SetFlag(FlagZ, a == 0)
	c.SetFlag(FlagH, false)

	return a
}


// Perform an AND operation, update flags, and return the result.
func (c *Cpu) opAnd(a uint8, b uint8) uint8 {
	r := a & b

	c.SetFlag(FlagZ, r == 0)
	c.SetFlag(FlagN, false)
	c.SetFlag(FlagH, true)
	c.SetFlag(FlagC, false)

	return r
}

// Perform an XOR operation, update flags, and return the result.
func (c *Cpu) opXor(a uint8, b uint8) uint8 {
	r := a ^ b

	c.SetFlag(FlagZ, r == 0)
	c.SetFlag(FlagN, false)
	c.SetFlag(FlagH, false)
	c.SetFlag(FlagC, false)

	return r
}

// Perform an OR operation, update flags, and return the result.
func (c *Cpu) opOr(a uint8, b uint8) uint8 {
	r := a | b

	c.SetFlag(FlagZ, r == 0)
	c.SetFlag(FlagN, false)
	c.SetFlag(FlagH, false)
	c.SetFlag(FlagC, false)

	return r
}

// Perform a CP operation, update flags, and return the result.
func (c *Cpu) opCp(a uint8, b uint8) uint8 {
	r := a - b

	c.SetFlag(FlagZ, r == 0)
	c.SetFlag(FlagN, true)
	c.SetFlag(FlagH, a & 0xf < b & 0xf)
	c.SetFlag(FlagC, a < b)

	return r
}

// Perform a 16 bit add, update flags, and return the result.
func (c *Cpu) opAdd16(a uint16, b uint16) uint16 {
	r32 := uint32(a) + uint32(b)
	r := uint16(r32)

	c.SetFlag(FlagN, false)
	c.SetFlag(FlagH, uint32(a & 0x0fff) > r32 & 0x0fff)
	c.SetFlag(FlagC, r32 > 0xffff)

	return r
}
