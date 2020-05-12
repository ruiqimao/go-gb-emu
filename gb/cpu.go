package gb

import (
	"fmt"

	"github.com/ruiqimao/go-gb-emu/utils"
)

type Cpu struct {
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

func NewCpu(gb *GameBoy) *Cpu {
	c := &Cpu{
		gb: gb,
		ic: 0xabcc, // Initial value of internal counter.
	}

	return c
}

// Perform a step and return how many cycles were used.
func (c *Cpu) Step() (int, error) {
	// Get whether interrupts will be handled this cycle. This gives us the behavior of EI taking
	// effect in the following cycle.
	ime := c.ime

	// Execute the next instruction.
	cycles, err := c.execNextInstruction()
	if err != nil {
		return cycles, err
	}

	// Handle interrupts.
	cycles += c.handleInterrupts(ime)

	// Update timers.
	c.updateTimers(cycles)

	return cycles, nil
}

// Execute the next instruction and return how many cycles were used.
func (c *Cpu) execNextInstruction() (int, error) {
	// If the CPU is halted, do nothing.
	if c.Halted() {
		return 4, nil
	}

	// Read an instruction.
	opCode := c.IncPC()
	inst := c.instructions[opCode]
	if inst == nil {
		return 0, fmt.Errorf("Unimplemented OP code: %02x", opCode)
	}

	// Execute the instruction.
	return inst(), nil
}

// Getters/setters for registers.
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

func (c *Cpu) A() uint8 {
	return c.a
}

func (c *Cpu) SetA(v uint8) {
	c.a = v
}

func (c *Cpu) F() uint8 {
	return c.f
}

func (c *Cpu) SetF(v uint8) {
	// Last 4 bits of F must be 0.
	c.f = v & 0xf0
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

func (c *Cpu) AF() uint16 {
	return utils.CombineBytes(c.a, c.f)
}

func (c *Cpu) SetAF(v uint16) {
	// Last 4 bits of F must be 0.
	c.a, c.f = utils.SplitShort(v & 0xfff0)
}

// Getters/setters for flags.
func (c *Cpu) FlagZ() bool {
	return utils.GetBit(c.f, 7)
}

func (c *Cpu) SetFlagZ(v bool) {
	c.f = utils.SetBit(c.f, 7, v)
}

func (c *Cpu) FlagN() bool {
	return utils.GetBit(c.f, 6)
}

func (c *Cpu) SetFlagN(v bool) {
	c.f = utils.SetBit(c.f, 6, v)
}

func (c *Cpu) FlagH() bool {
	return utils.GetBit(c.f, 5)
}

func (c *Cpu) SetFlagH(v bool) {
	c.f = utils.SetBit(c.f, 5, v)
}

func (c *Cpu) FlagC() bool {
	return utils.GetBit(c.f, 4)
}

func (c *Cpu) SetFlagC(v bool) {
	c.f = utils.SetBit(c.f, 4, v)
}

// Get the stack pointer.
func (c *Cpu) SP() uint16 {
	return c.sp
}

// Push a value to the stack.
func (c *Cpu) PushSP(v uint16) {
	c.sp -= 2
	c.gb.mem.Write16(c.sp, v)
}

// Pop a value from the stack.
func (c *Cpu) PopSP() uint16 {
	v := c.gb.mem.Read16(c.sp)
	c.sp += 2
	return v
}

// Set the stack pointer.
func (c *Cpu) SetSP(v uint16) {
	c.sp = v
}

// Get the program counter.
func (c *Cpu) PC() uint16 {
	return c.pc
}

// Increment the program counter by a byte and return the read value.
func (c *Cpu) IncPC() uint8 {
	v := c.gb.mem.Read(c.pc)

	// If the halt bug is active, the program counter is not incremented.
	if !c.haltBug {
		c.pc++
	}
	c.haltBug = false
	return v
}

// Increment the program counter by a short and return the read value.
func (c *Cpu) IncPC16() uint16 {
	lo := c.IncPC()
	hi := c.IncPC()
	return utils.CombineBytes(hi, lo)
}

// Set the program counter.
func (c *Cpu) SetPC(v uint16) {
	c.pc = v
}

// Get the internal counter.
func (c *Cpu) IC() uint16 {
	return c.ic
}

// Set the internal counter.
func (c *Cpu) SetIC(v uint16) {
	c.ic = v
}

// Get interrupt master.
func (c *Cpu) IME() bool {
	return c.ime
}

// Set interrupt master.
func (c *Cpu) SetIME(v bool) {
	c.ime = v
}

// Get the halt flag.
func (c *Cpu) Halted() bool {
	return c.halt
}

// Set the halt flag.
func (c *Cpu) SetHalt(v bool) {
	c.halt = v
}

// Trigger the halt bug.
func (c *Cpu) TriggerHaltBug() {
	c.haltBug = true
}
