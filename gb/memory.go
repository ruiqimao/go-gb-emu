package gb

import (
	"github.com/ruiqimao/go-gb-emu/utils"
)

// Memory layout:
// 0000 - 00FF
//   Boot ROM - This gets turned off at the end of execution.
// 0000 - 3FFF
//   Cartridge ROM bank 0.
// 4000 - 7FFF
//   Cartridge ROM bank 1 - N.
// 8000 - 9FFF
//   Video RAM.
// A000 - BFFF
//   Cartridge RAM.
// C000 - CFFF
//   Work RAM bank 0.
// D000 - DFFF
//   Work RAM bank 1.
// E000 - FDFF
//   Mirror of C000 - DDFF.
// FE00 - FE9F
//   Sprite attribute table (OAM).
// FEA0 - FEFF
//   Empty. All reads are 0, and all writes are no-op.
// FF00 - FF7F
//   I/O registers.
// FF80 - FFFE
//   High RAM.
// FFFF
//   Interrupts enable register.
const (
	AddrBoot = 0x0000
	AddrCartROM0 = 0x0000
	AddrCartROMN = 0x4000
	AddrVRAM = 0x8000
	AddrCartRAM = 0xa000
	AddrWRAM0 = 0xc000
	AddrWRAMN = 0xd000
	AddrEcho = 0xe000
	AddrOAM = 0xfe00
	AddrEmpty = 0xfea0
	AddrIO = 0xff00
	AddrHRAM = 0xff80
)

type Memory struct {
	gb *GameBoy

	vram [0x4000]uint8
	wram [0x2000]uint8
	oam  [0x100]uint8
	hram [0xff]uint8
	io   [0x80]uint8
	ie   uint8
}

func NewMemory(gb *GameBoy) *Memory {
	m := &Memory{
		gb: gb,
	}
	return m
}

// Read a byte from memory.
func (m *Memory) Read(addr uint16) uint8 {
	if m.gb.Debugging() && int(addr) < len(m.gb.dbgRom) {
		return m.gb.dbgRom[addr]
	}

	switch {

	// Boot ROM.
	case addr < 0x0100 && m.BootRomEnabled():
		return m.gb.boot[addr]

	// Cartridge ROM bank 0.
	case addr < AddrCartROMN:
		// TODO.
		return 0x00

	// Cartridge ROM bank 1 - N.
	case addr < AddrVRAM:
		// TODO.
		return 0x00

	// Video RAM.
	case addr < AddrCartRAM:
		// Make sure VRAM is currently accessible.
		mode := m.gb.ppu.Mode()
		if mode == ModeTransfer {
			return 0x00
		}
		return m.vram[addr-AddrVRAM]

	// Cartridge RAM.
	case addr < AddrWRAM0:
		// TODO.
		return 0x00

	// Work RAM banks 0 and 1.
	case addr < AddrEcho:
		return m.wram[addr-AddrWRAM0]

	// Mirror of C000 - DDFF.
	case addr < AddrOAM:
		return m.Read(addr - 0x2000)

	// Sprite attribute table.
	case addr < AddrEmpty:
		// Make sure the OAM is currently accessible.
		mode := m.gb.ppu.Mode()
		if mode == ModeTransfer || mode == ModeOAM {
			return 0x00
		}
		return m.oam[addr-AddrOAM]

	// Empty.
	case addr < AddrIO:
		return 0x00

	// I/O registers.
	case addr < AddrHRAM:
		return m.ReadIO(addr)

	// High RAM.
	case addr < AddrIE:
		return m.hram[addr-AddrIO]

	// Interrupt enable register.
	case addr == AddrIE:
		return m.ie

	}

	// This line will never be reached.
	return 0x00
}

// Write a byte into memory.
func (m *Memory) Write(addr uint16, v uint8) {
	switch {

	// Video RAM.
	case addr >= AddrVRAM && addr < AddrCartRAM:
		// Make sure VRAM is currently accessible.
		mode := m.gb.ppu.Mode()
		if mode == ModeTransfer {
			break
		}
		m.vram[addr-AddrVRAM] = v

	// Cartridge RAM.
	case addr >= AddrCartRAM && addr < AddrWRAM0:
		// TODO.

	// Work RAM banks 0 and 1.
	case addr >= AddrWRAM0 && addr < AddrEcho:
		m.wram[addr-AddrWRAM0] = v

	// Mirror of C000 - DDFF.
	case addr >= AddrEcho && addr < AddrOAM:
		m.Write(addr-0x2000, v)

	// Sprite attribute table.
	case addr >= AddrOAM && addr < AddrIO:
		// Make sure the OAM is currently accessible.
		mode := m.gb.ppu.Mode()
		if mode == ModeTransfer || mode == ModeOAM {
			break
		}
		m.oam[addr-AddrOAM] = v

	// I/O registers.
	case addr >= AddrIO && addr < AddrHRAM:
		m.WriteIO(addr, v)

	// High RAM.
	case addr >= AddrHRAM && addr < AddrIE:
		m.hram[addr-AddrHRAM] = v

	// Interrupt enable register.
	case addr == AddrIE:
		m.ie = v

	}
}

// Read a short from memory.
func (m *Memory) Read16(addr uint16) uint16 {
	return utils.CombineBytes(m.Read(addr+1), m.Read(addr))
}

// Write a short into memory.
func (m *Memory) Write16(addr uint16, v uint16) {
	hi, lo := utils.SplitShort(v)
	m.Write(addr, lo)
	m.Write(addr+1, hi)
}
