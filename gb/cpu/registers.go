package cpu

import (
	"github.com/ruiqimao/go-gb-emu/utils"
)

// Explicit types for type safety.
type Register int
type Register16 int
type Flag int

const (
	RegisterB Register = 0
	RegisterC          = 1
	RegisterD          = 2
	RegisterE          = 3
	RegisterH          = 4
	RegisterL          = 5
	RegisterA          = 6
	RegisterF          = 7
)
const (
	RegisterBC Register16 = 0
	RegisterDE            = 2
	RegisterHL            = 4
	RegisterAF            = 6
)
const (
	FlagC = 4
	FlagH = 5
	FlagN = 6
	FlagZ = 7
)

func (c *CPU) getRegister(reg Register) uint8 {
	return c.rg[reg]
}

func (c *CPU) getRegister16(reg Register16) uint16 {
	return utils.CombineBytes(c.rg[reg], c.rg[reg+1])
}

func (c *CPU) setRegister(reg Register, v uint8) {
	if reg == RegisterF {
		v &= 0xf0 // Lower 4 bits of F are always zero.
	}
	c.rg[reg] = v
}

func (c *CPU) setRegister16(reg Register16, v uint16) {
	if reg == RegisterAF {
		v &= 0xfff0 // Lower 4 bits of F are always zero.
	}
	c.rg[reg], c.rg[reg+1] = utils.SplitShort(v)
}

func (c *CPU) getFlag(flag Flag) bool {
	return utils.GetBit(c.rg[RegisterF], int(flag))
}

func (c *CPU) setFlag(flag Flag, v bool) {
	c.rg[RegisterF] = utils.SetBit(c.rg[RegisterF], int(flag), v)
}