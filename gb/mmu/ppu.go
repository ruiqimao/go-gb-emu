package mmu

// PPU interface.
type PPU interface {
	LCDC() uint8
	STAT() uint8
	SCY() uint8
	SCX() uint8
	LY() uint8
	LYC() uint8
	DMA() uint8
	BGP() uint8
	OBP0() uint8
	OBP1() uint8
	WY() uint8
	WX() uint8

	SetLCDC(uint8)
	SetSTAT(uint8)
	SetSCY(uint8)
	SetSCX(uint8)
	SetLY(uint8)
	SetLYC(uint8)
	SetDMA(uint8)
	SetBGP(uint8)
	SetOBP0(uint8)
	SetOBP1(uint8)
	SetWY(uint8)
	SetWX(uint8)
}

type PPUBus struct {
	Read  func(addr uint16) uint8
	Write func(addr uint16, v uint8)
}

// Handle read operations from the PPU.
func (m *MMU) cpuRead(addr uint16) uint8 {
	return m.read(addr)
}

// Handle write operations from the PPU.
func (m *MMU) cpuWrite(addr uint16, v uint8) {
	m.write(addr, v)
}
