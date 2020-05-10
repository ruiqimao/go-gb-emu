package gb

import (
	"github.com/ruiqimao/go-gb-emu/utils"
)

type Cpu struct {
	// Registers.
	a uint8
	b uint8
	c uint8
	d uint8
	e uint8
	f uint8
	h uint8
	l uint8

	// Stack pointer.
	sp uint16

	// Program counter.
	pc uint16
}

func NewCpu() *Cpu {
	c := &Cpu{}
	return c
}

// Getters/setters for general registers.
func (c *Cpu) A() uint8 {
	return c.a
}

func (c *Cpu) SetA(v uint8) {
	c.a = v
}

func (c *Cpu) B() uint8 {
	return c.b
}

func (c *Cpu) SetB(v uint8) {
	c.b = v
}

func (c *Cpu) C() uint8 {
	return c.c
}

func (c *Cpu) SetC(v uint8) {
	c.c = v
}

func (c *Cpu) D() uint8 {
	return c.d
}

func (c *Cpu) SetD(v uint8) {
	c.d = v
}

func (c *Cpu) E() uint8 {
	return c.e
}

func (c *Cpu) SetE(v uint8) {
	c.e = v
}

func (c *Cpu) H() uint8 {
	return c.h
}

func (c *Cpu) SetH(v uint8) {
	c.h = v
}

func (c *Cpu) L() uint8 {
	return c.l
}

func (c *Cpu) SetL(v uint8) {
	c.l = v
}

func (c *Cpu) AF() uint16 {
	return utils.CombineBytes(c.a, c.f)
}

func (c *Cpu) SetAF(v uint16) {
	c.a, c.f = utils.SplitShort(v)
	c.f &= 0xf0 // Last 4 bits of F register are always 0.
}

func (c *Cpu) BC() uint16 {
	return utils.CombineBytes(c.b, c.c)
}

func (c *Cpu) SetBC(v uint16) {
	c.b, c.c = utils.SplitShort(v)
}

func (c *Cpu) DE() uint16 {
	return utils.CombineBytes(c.d, c.e)
}

func (c *Cpu) SetDE(v uint16) {
	c.d, c.e = utils.SplitShort(v)
}

func (c *Cpu) HL() uint16 {
	return utils.CombineBytes(c.h, c.l)
}

func (c *Cpu) SetHL(v uint16) {
	c.h, c.l = utils.SplitShort(v)
}

// Getters/setters for flag (F) register.
func (c *Cpu) FC() bool {
	return utils.GetBit(c.f, 4)
}

func (c *Cpu) SetFC(v bool) {
	utils.SetBit(c.f, 4, v)
}

func (c *Cpu) FH() bool {
	return utils.GetBit(c.f, 5)
}

func (c *Cpu) SetFH(v bool) {
	utils.SetBit(c.f, 5, v)
}

func (c *Cpu) FN() bool {
	return utils.GetBit(c.f, 6)
}

func (c *Cpu) SetFN(v bool) {
	utils.SetBit(c.f, 6, v)
}

func (c *Cpu) FZ() bool {
	return utils.GetBit(c.f, 7)
}

func (c *Cpu) SetFZ(v bool) {
	utils.SetBit(c.f, 7, v)
}

// Getters/setters for stack pointer and program counter.
func (c *Cpu) SP() uint16 {
	return c.sp
}

func (c *Cpu) IncrementSP() {
	c.sp += 2
}

func (c *Cpu) DecrementSP() {
	c.sp -= 2
}

func (c *Cpu) PC() uint16 {
	return c.pc
}

func (c *Cpu) IncrementPC() {
	c.pc += 2
}

func (c *Cpu) SetPC(v uint16) {
	c.pc = v
}
