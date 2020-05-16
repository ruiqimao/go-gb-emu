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
	AddrBoot     = 0x0000
	AddrCartROM0 = 0x0000
	AddrCartROMN = 0x4000
	AddrVRAM     = 0x8000
	AddrCartRAM  = 0xa000
	AddrWRAM0    = 0xc000
	AddrWRAMN    = 0xd000
	AddrEcho     = 0xe000
	AddrOAM      = 0xfe00
	AddrEmpty    = 0xfea0
	AddrIO       = 0xff00
	AddrHRAM     = 0xff80
)

type Memory struct {
	gb *GameBoy

	wram [0x2000]uint8
	hram [0xff]uint8

	// Scratch space for I/O registers.
	IO [0x80]uint8
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

	// Cartridge ROM banks.
	case addr < AddrVRAM:
		if m.gb.cart != nil {
			return m.gb.cart.ReadROM(addr)
		}

	// Video RAM.
	case addr < AddrCartRAM:
		if m.gb.cart != nil {
			return m.gb.ppu.ReadVRAM(addr - AddrVRAM)
		}

	// Cartridge RAM.
	case addr < AddrWRAM0:
		return m.gb.cart.ReadRAM(addr - AddrCartRAM)

	// Work RAM banks 0 and 1.
	case addr < AddrEcho:
		return m.wram[addr-AddrWRAM0]

	// Mirror of C000 - DDFF.
	case addr < AddrOAM:
		return m.Read(addr - 0x2000)

	// Sprite attribute table.
	case addr < AddrEmpty:
		return m.gb.ppu.ReadOAM(addr - AddrOAM)

	// Empty.
	case addr < AddrIO:
		return 0x00

	// I/O registers.
	case addr < AddrHRAM:
		return m.ReadIO(addr)

	// High RAM.
	case addr < AddrIE:
		return m.hram[addr-AddrHRAM]

	// Interrupt enable register.
	case addr == AddrIE:
		return m.gb.cpu.IE()

	}

	return 0x00
}

// Write a byte into memory.
func (m *Memory) Write(addr uint16, v uint8) {
	switch {

	// Cartridge ROM.
	case addr < AddrVRAM:
		if m.gb.cart != nil {
			m.gb.cart.WriteROM(addr, v)
		}

	// Video RAM.
	case addr >= AddrVRAM && addr < AddrCartRAM:
		m.gb.ppu.WriteVRAM(addr-AddrVRAM, v)

	// Cartridge RAM.
	case addr >= AddrCartRAM && addr < AddrWRAM0:
		if m.gb.cart != nil {
			m.gb.cart.WriteROM(addr-AddrCartRAM, v)
		}

	// Work RAM banks 0 and 1.
	case addr >= AddrWRAM0 && addr < AddrEcho:
		m.wram[addr-AddrWRAM0] = v

	// Mirror of C000 - DDFF.
	case addr >= AddrEcho && addr < AddrOAM:
		m.Write(addr-0x2000, v)

	// Sprite attribute table.
	case addr >= AddrOAM && addr < AddrIO:
		m.gb.ppu.WriteOAM(addr-AddrOAM, v)

	// I/O registers.
	case addr >= AddrIO && addr < AddrHRAM:
		m.WriteIO(addr, v)

	// High RAM.
	case addr >= AddrHRAM && addr < AddrIE:
		m.hram[addr-AddrHRAM] = v

	// Interrupt enable register.
	case addr == AddrIE:
		m.gb.cpu.SetIE(v)

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
