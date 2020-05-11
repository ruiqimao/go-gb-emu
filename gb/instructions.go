package gb

import (
	"github.com/ruiqimao/go-gb-emu/utils"
)

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
			r, flags := opSignedAdd(cpu.Sp(), cpu.IncPc())
			cpu.Set16(RegHL, r)
			cpu.Set(RegF, flags)
			return 12
		},
		0xf9: func() int { // LD SP,HL.
			cpu.SetSp(cpu.Get16(RegHL))
			return 8
		},
	}
}

// Perform an add and return the result and updated flags.
func opAdd(a uint8, b uint8, c uint8) (uint8, uint8) {
	r16 := uint16(a) + uint16(b) + uint16(c)
	r := uint8(r16)

	flags := uint8(0)
	flags = utils.SetBit(flags, FlagZ, r == 0)
	flags = utils.SetBit(flags, FlagN, false)
	flags = utils.SetBit(flags, FlagH, (a & 0xf) + (b & 0xf) + (c & 0xf) > 0xf)
	flags = utils.SetBit(flags, FlagC, r16 > 0xff)

	return r, flags
}

// Perform a signed add and return the result and updated flags.
func opSignedAdd(a uint16, b uint8) (uint16, uint8) {
	r := uint16(int32(a) + int32(int8(b)))

	// Get the flags from doing an ordinary add.
	_, flags := opAdd(uint8(a), b, 0)
	flags = utils.SetBit(flags, FlagZ, false)
	flags = utils.SetBit(flags, FlagN, false)

	return r, flags
}
