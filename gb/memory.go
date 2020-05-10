package gb

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

type Memory struct {
	gb *GameBoy

	vram [0x4000]uint8
	wram [0x2000]uint8
	oam  [0x100]uint8
	hram [0xff]uint8
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
	switch {

	// Boot ROM.
	case addr < 0x0100 && m.BootRomEnabled():
		return m.gb.boot[addr]

	// Cartridge ROM bank 0.
	case addr < 0x4000:
		// TODO.
		return 0x00

	// Cartridge ROM bank 1 - N.
	case addr < 8000:
		// TODO.
		return 0x00

	// Video RAM.
	case addr < 0xa000:
		return m.vram[addr-0x8000]

	// Cartridge RAM.
	case addr < 0xc000:
		// TODO.
		return 0x00

	// Work RAM banks 0 and 1.
	case addr < 0xe000:
		return m.wram[addr-0xc000]

	// Mirror of C000 - DDFF.
	case addr < 0xfe00:
		return m.Read(addr - 0x2000)

	// Sprite attribute table.
	case addr < 0xfea0:
		return m.oam[addr-0xfe00]

	// Empty.
	case addr < 0xff00:
		return 0x00

	// I/O registers.
	case addr < 0xff80:
		return m.ReadIO(addr)

	// High RAM.
	case addr < 0xffff:
		return m.hram[addr-0xff80]

	// Interrupt enable register.
	case addr == 0xffff:
		return m.ie

	}

	// This line will never be reached.
	return 0x00
}

// Write a byte into memory.
func (m *Memory) Write(addr uint16, v uint8) {
	switch {

	// Video RAM.
	case addr >= 0x8000 && addr < 0xa000:
		m.vram[addr-0x8000] = v

	// Cartridge RAM.
	case addr >= 0xa000 && addr < 0xc000:
		// TODO.

	// Work RAM banks 0 and 1.
	case addr >= 0xc000 && addr < 0xe000:
		m.wram[addr-0xc000] = v

	// Mirror of C000 - DDFF.
	case addr >= 0xe000 && addr < 0xfe00:
		m.Write(addr-0x2000, v)

	// Sprite attribute table.
	case addr >= 0xfe00 && addr < 0xfea0:
		m.oam[addr-0xfe00] = v

	// I/O registers.
	case addr >= 0xff00 && addr < 0xff80:
		m.WriteIO(addr, v)

	// High RAM.
	case addr >= 0xff80 && addr < 0xffff:
		m.hram[addr-0xff80] = v

	// Interrupt enable register.
	case addr == 0xffff:
		m.ie = v

	}
}
