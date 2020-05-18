package gb

import (
	"fmt"
	"strings"

	"github.com/ruiqimao/go-gb-emu/gb/cpu"
	"github.com/ruiqimao/go-gb-emu/gb/joypad"
	"github.com/ruiqimao/go-gb-emu/gb/mmu"
	"github.com/ruiqimao/go-gb-emu/gb/ppu"
)

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
	return gb.RunClocks(1) + 1
}

// Get a readable version of the current instruction.
func (gb *GameBoy) InstructionName() string {
	opCode := uint16(gb.mmu.Read(gb.cpu.PC()))
	if opCode == 0xcb {
		opCode = uint16(gb.mmu.Read(gb.cpu.PC()+1)) + 0x100
	}
	name := InstructionNames[opCode]

	// Get all the possible components.
	d16 := gb.mmu.Read16(gb.cpu.PC() + 0x1)
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

// Get the CPU.
func (gb *GameBoy) CPU() *cpu.CPU {
	return gb.cpu
}

// Get the PPU.
func (gb *GameBoy) PPU() *ppu.PPU {
	return gb.ppu
}

// Get the joypad.
func (gb *GameBoy) Joypad() *joypad.Joypad {
	return gb.jp
}

// Get the memory controller.
func (gb *GameBoy) MMU() *mmu.MMU {
	return gb.mmu
}

// Get the clock.
func (gb *GameBoy) Clock() *Clock {
	return gb.clk
}
