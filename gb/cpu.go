package gb

import (
	"fmt"

	"github.com/ruiqimao/go-gb-emu/utils"
)

type CPU struct {
	gb *GameBoy

	// Instruction set.
	instructions [0x200]Instruction

	// Registers.
	b uint8
	c uint8
	d uint8
	e uint8
	h uint8
	l uint8
	a uint8
	f uint8

	// Stack pointer.
	sp uint16

	// Program counter.
	pc uint16

	// Internal counter.
	ic uint16

	// Interrupt master.
	ime bool

	// Halt and bug flags.
	halt    bool
	haltBug bool

	// Misc flags.
	of bool // Timer overflow.

}

func NewCPU(gb *GameBoy) *CPU {
	c := &CPU{
		gb: gb,
		ic: 0xabcc, // Initial value of internal counter.
	}

	return c
}

// Perform a step and return how many cycles were used.
func (c *CPU) Step() (int, error) {
	// Get whether interrupts will be handled this cycle. This gives us the behavior of EI taking
	// effect in the following cycle.
	ime := c.ime

	// Execute the next instruction.
	var cycles int = 0
	if c.Halted() {
		cycles = 4
	} else {
		// Read an instruction.
		opCode := c.IncPC()
		inst := c.instructions[opCode]
		if inst == nil {
			return 0, fmt.Errorf("Unimplemented OP code: %02x", opCode)
		}

		// Execute the instruction.
		cycles = inst()
	}

	// Handle interrupts.
	cycles += c.handleInterrupts(ime)

	// Update timers.
	c.updateTimers(cycles)

	return cycles, nil
}

// Getters/setters for registers.
func (c *CPU) B() uint8 {
	return c.b
}

func (c *CPU) SetB(v uint8) {
	c.b = v
}

func (c *CPU) C() uint8 {
	return c.c
}

func (c *CPU) SetC(v uint8) {
	c.c = v
}

func (c *CPU) D() uint8 {
	return c.d
}

func (c *CPU) SetD(v uint8) {
	c.d = v
}

func (c *CPU) E() uint8 {
	return c.e
}

func (c *CPU) SetE(v uint8) {
	c.e = v
}

func (c *CPU) H() uint8 {
	return c.h
}

func (c *CPU) SetH(v uint8) {
	c.h = v
}

func (c *CPU) L() uint8 {
	return c.l
}

func (c *CPU) SetL(v uint8) {
	c.l = v
}

func (c *CPU) A() uint8 {
	return c.a
}

func (c *CPU) SetA(v uint8) {
	c.a = v
}

func (c *CPU) F() uint8 {
	return c.f
}

func (c *CPU) SetF(v uint8) {
	// Last 4 bits of F must be 0.
	c.f = v & 0xf0
}

func (c *CPU) BC() uint16 {
	return utils.CombineBytes(c.b, c.c)
}

func (c *CPU) SetBC(v uint16) {
	c.b, c.c = utils.SplitShort(v)
}

func (c *CPU) DE() uint16 {
	return utils.CombineBytes(c.d, c.e)
}

func (c *CPU) SetDE(v uint16) {
	c.d, c.e = utils.SplitShort(v)
}

func (c *CPU) HL() uint16 {
	return utils.CombineBytes(c.h, c.l)
}

func (c *CPU) SetHL(v uint16) {
	c.h, c.l = utils.SplitShort(v)
}

func (c *CPU) AF() uint16 {
	return utils.CombineBytes(c.a, c.f)
}

func (c *CPU) SetAF(v uint16) {
	// Last 4 bits of F must be 0.
	c.a, c.f = utils.SplitShort(v & 0xfff0)
}

// Getters/setters for flags.
func (c *CPU) FlagZ() bool {
	return utils.GetBit(c.f, 7)
}

func (c *CPU) SetFlagZ(v bool) {
	c.f = utils.SetBit(c.f, 7, v)
}

func (c *CPU) FlagN() bool {
	return utils.GetBit(c.f, 6)
}

func (c *CPU) SetFlagN(v bool) {
	c.f = utils.SetBit(c.f, 6, v)
}

func (c *CPU) FlagH() bool {
	return utils.GetBit(c.f, 5)
}

func (c *CPU) SetFlagH(v bool) {
	c.f = utils.SetBit(c.f, 5, v)
}

func (c *CPU) FlagC() bool {
	return utils.GetBit(c.f, 4)
}

func (c *CPU) SetFlagC(v bool) {
	c.f = utils.SetBit(c.f, 4, v)
}

// Get the stack pointer.
func (c *CPU) SP() uint16 {
	return c.sp
}

// Push a value to the stack.
func (c *CPU) PushSP(v uint16) {
	c.sp -= 2
	c.gb.mem.Write16(c.sp, v)
}

// Pop a value from the stack.
func (c *CPU) PopSP() uint16 {
	v := c.gb.mem.Read16(c.sp)
	c.sp += 2
	return v
}

// Set the stack pointer.
func (c *CPU) SetSP(v uint16) {
	c.sp = v
}

// Get the program counter.
func (c *CPU) PC() uint16 {
	return c.pc
}

// Increment the program counter by a byte and return the read value.
func (c *CPU) IncPC() uint8 {
	v := c.gb.mem.Read(c.pc)

	// If the halt bug is active, the program counter is not incremented.
	if !c.haltBug {
		c.pc++
	}
	c.haltBug = false
	return v
}

// Increment the program counter by a short and return the read value.
func (c *CPU) IncPC16() uint16 {
	lo := c.IncPC()
	hi := c.IncPC()
	return utils.CombineBytes(hi, lo)
}

// Set the program counter.
func (c *CPU) SetPC(v uint16) {
	c.pc = v
}

// Get the internal counter.
func (c *CPU) IC() uint16 {
	return c.ic
}

// Set the internal counter.
func (c *CPU) SetIC(v uint16) {
	c.ic = v
}

// Get interrupt master.
func (c *CPU) IME() bool {
	return c.ime
}

// Set interrupt master.
func (c *CPU) SetIME(v bool) {
	c.ime = v
}

// Get the halt flag.
func (c *CPU) Halted() bool {
	return c.halt
}

// Set the halt flag.
func (c *CPU) SetHalt(v bool) {
	c.halt = v
}

// Trigger the halt bug.
func (c *CPU) TriggerHaltBug() {
	c.haltBug = true
}
