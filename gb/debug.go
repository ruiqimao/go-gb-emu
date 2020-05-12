package gb

import (
	"fmt"
	"strings"
)

// A snapshot of the Game Boy state.
type Snapshot struct {
	// Registers.
	B  uint8
	C  uint8
	D  uint8
	E  uint8
	H  uint8
	L  uint8
	A  uint8
	F  uint8
	BC uint16
	DE uint16
	HL uint16
	AF uint16

	// Flags.
	FlagZ bool
	FlagN bool
	FlagH bool
	FlagC bool

	// Stack pointer.
	SP uint16

	// Program counter.
	PC              uint16
	InstructionName string

	// Memory.
	BootROM [0x100]uint8
	VRAM    [0x4000]uint8
	WRAM    [0x2000]uint8
	OAM     [0x100]uint8
	HRAM    [0xff]uint8
	IE      uint8
	Memory  [0x10000]uint8 // Full map.

	// Misc.
	Halt    bool
	HaltBug bool
	IME     bool
}

// Load a debug ROM.
func (gb *GameBoy) LoadDebugRom(r []uint8) {
	gb.dbgRom = make([]uint8, len(r))
	copy(gb.dbgRom, r)

	// Initialize manually due to no Boot ROM.
	gb.cpu.SetSP(0xfffe)
	gb.cpu.SetPC(0x0100)
	gb.cpu.SetBC(0x0013)
	gb.cpu.SetDE(0x00d8)
	gb.cpu.SetHL(0x014d)
	gb.cpu.SetAF(0x01b0)
}

// Whether debug is enabled.
func (gb *GameBoy) Debugging() bool {
	return gb.dbgRom != nil
}

// Run the Game Boy clock.
func (gb *GameBoy) Resume() {
	gb.clk.Resume()
}

// Pause the Game Boy clock.
func (gb *GameBoy) Pause() {
	gb.clk.Pause()
}

// Step forward by one instruction. Returns how many cycles were taken.
func (gb *GameBoy) Step() int {
	return gb.RunCycles(1) + 1
}

// Take a snapshot of the Game Boy state.
func (gb *GameBoy) Snapshot() Snapshot {
	ss := Snapshot{}

	// Registers.
	ss.B = gb.cpu.B()
	ss.C = gb.cpu.C()
	ss.D = gb.cpu.D()
	ss.E = gb.cpu.E()
	ss.H = gb.cpu.H()
	ss.L = gb.cpu.L()
	ss.A = gb.cpu.A()
	ss.F = gb.cpu.F()
	ss.BC = gb.cpu.BC()
	ss.DE = gb.cpu.DE()
	ss.HL = gb.cpu.HL()
	ss.AF = gb.cpu.AF()

	// Stack pointer.
	ss.SP = gb.cpu.SP()

	// Program counter.
	ss.PC = gb.cpu.PC()
	ss.InstructionName = gb.InstructionName()

	// Flags.
	ss.FlagZ = gb.cpu.FlagZ()
	ss.FlagN = gb.cpu.FlagN()
	ss.FlagH = gb.cpu.FlagH()
	ss.FlagC = gb.cpu.FlagC()

	// Memory.
	copy(ss.BootROM[:], gb.boot[:])
	copy(ss.VRAM[:], gb.mem.vram[:])
	copy(ss.WRAM[:], gb.mem.wram[:])
	copy(ss.OAM[:], gb.mem.oam[:])
	copy(ss.HRAM[:], gb.mem.hram[:])
	ss.IE = gb.mem.ie
	for i := 0; i < 0x10000; i++ {
		ss.Memory[i] = gb.mem.Read(uint16(i))
	}

	// Misc.
	ss.Halt = gb.cpu.Halted()
	ss.HaltBug = gb.cpu.haltBug
	ss.IME = gb.cpu.IME()

	return ss
}

// Get the program counter for breakpoints.
func (gb *GameBoy) PC() uint16 {
	return gb.cpu.PC()
}

// Get a readable version of the current instruction.
func (gb *GameBoy) InstructionName() string {
	opCode := uint16(gb.mem.Read(gb.cpu.PC()))
	if opCode == 0xcb {
		opCode = uint16(gb.mem.Read(gb.cpu.PC())+1) + 0x100
	}
	name := InstructionNames[opCode]

	// Get all the possible components.
	d16 := gb.mem.Read16(gb.cpu.PC() + 0x1)
	d8 := uint8(d16)
	a16 := fmt.Sprintf("$%04x", d16)
	a8 := fmt.Sprintf("$%02x", d8)
	r8 := int8(d8)

	// Replace tokens.
	if r8 < 0 {
		name = strings.ReplaceAll(name, "+r8", "r8")
		name = strings.ReplaceAll(name, "r8", "-r8")
		r8 *= -1
	}
	name = strings.ReplaceAll(name, "d16", fmt.Sprintf("%04x", d16))
	name = strings.ReplaceAll(name, "d8", fmt.Sprintf("%02x", d8))
	name = strings.ReplaceAll(name, "a16", a16)
	name = strings.ReplaceAll(name, "a8", a8)
	name = strings.ReplaceAll(name, "r8", fmt.Sprintf("%02x", r8))

	return name
}
