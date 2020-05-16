package gb

import (
	"fmt"
	"strings"

	"github.com/ruiqimao/go-gb-emu/gb/cpu"
)

// Load a debug ROM.
func (gb *GameBoy) LoadDebugRom(r []uint8) {
	gb.dbgRom = make([]uint8, len(r))
	copy(gb.dbgRom, r)

	// Initialize manually due to no Boot ROM.
	gb.cpu.SetSP(0xfffe)
	gb.cpu.SetPC(0x0100)
	gb.cpu.SetRegister16(cpu.RegisterBC, 0x0013)
	gb.cpu.SetRegister16(cpu.RegisterDE, 0x00d8)
	gb.cpu.SetRegister16(cpu.RegisterHL, 0x014d)
	gb.cpu.SetRegister16(cpu.RegisterAF, 0x01b0)
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

// Get a readable version of the current instruction.
func (gb *GameBoy) InstructionName() string {
	opCode := uint16(gb.mem.Read(gb.cpu.PC()))
	if opCode == 0xcb {
		opCode = uint16(gb.mem.Read(gb.cpu.PC()+1)) + 0x100
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

// Get the raw contents of the VRAM.
func (p *PPU) VRAM() []uint8 {
	return p.vram[:]
}

// Get the raw contents of the OAM.
func (p *PPU) OAM() []uint8 {
	return p.oam[:]
}
