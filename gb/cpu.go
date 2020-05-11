package gb

import (
	"github.com/ruiqimao/go-gb-emu/utils"
)

const (
	// 8 bit registers.
	RegB = 0
	RegC = 1
	RegD = 2
	RegE = 3
	RegH = 4
	RegL = 5
	RegA = 6
	RegF = 7

	// 16 bit registers.
	RegBC = 0
	RegDE = 2
	RegHL = 4
	RegAF = 6

	// Flags.
	FlagC = 4
	FlagH = 5
	FlagN = 6
	FlagZ = 7
)

type Cpu struct {
	gb *GameBoy

	// Registers.
	r [8]uint8

	// Stack pointer.
	sp uint16

	// Program counter.
	pc uint16
}

func NewCpu(gb *GameBoy) *Cpu {
	c := &Cpu{
		gb: gb,
	}
	return c
}

// Get the value of a register.
func (c *Cpu) Get(r uint8) uint8 {
	return c.r[r]
}

// Set the value of a register.
func (c *Cpu) Set(r uint8, v uint8) {
	if r == RegF {
		// Last 4 bits of F cannot be set.
		v &= 0xf0
	}
	c.r[r] = v
}

// Get the value of a 16 bit register.
func (c *Cpu) Get16(r uint8) uint16 {
	return utils.CombineBytes(c.r[r], c.r[r+1])
}

// Set the value of a 16 bit register.
func (c *Cpu) Set16(r uint8, v uint16) {
	if r == RegAF {
		// Last 4 bits of F cannot be set.
		v &= 0xfff0
	}
	c.r[r], c.r[r+1] = utils.SplitShort(v)
}

// Get whether a flag is enabled.
func (c *Cpu) GetFlag(f int) bool {
	return utils.GetBit(c.r[RegF], f)
}

// Set whether a flag is enabled.
func (c *Cpu) SetFlag(f int, v bool) {
	utils.SetBit(c.r[RegF], f, v)
}

// Get the stack pointer.
func (c *Cpu) Sp() uint16 {
	return c.sp
}

// Push a value to the stack.
func (c *Cpu) PushSp(v uint16) {
	c.sp -= 2
	c.gb.mem.Write16(c.sp, v)
}

// Pop a value from the stack.
func (c *Cpu) PopSp() uint16 {
	v := c.gb.mem.Read16(c.sp)
	c.sp += 2
	return v
}

// Set the stack pointer.
func (c *Cpu) SetSp(v uint16) {
	c.sp = v
}

// Get the program counter.
func (c *Cpu) Pc() uint16 {
	return c.pc
}

// Increment the program counter by a byte and return the read value.
func (c *Cpu) IncPc() uint8 {
	v := c.gb.mem.Read(c.pc)
	c.pc++
	return v
}

// Increment the program counter by a short and return the read value.
func (c *Cpu) IncPc16() uint16 {
	lo := c.IncPc()
	hi := c.IncPc()
	return utils.CombineBytes(hi, lo)
}

// Set the program counter.
func (c *Cpu) SetPc(v uint16) {
	c.pc = v
}
