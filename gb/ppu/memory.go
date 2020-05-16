package ppu

// MMU interface.
type MMU interface {
	Read(uint16) uint8
	Write(uint16, uint8)
}

// Reads from the VRAM. The address is in the range [0x0000, 0x4000).
func (p *PPU) ReadVRAM(addr uint16) uint8 {
	return p.vram[addr]
}

// Reads from the OAM. The address is in the range [0x0000, 0x0100).
func (p *PPU) ReadOAM(addr uint16) uint8 {
	return p.oam[addr]
}

// Writes to the VRAM. The address is in the range [0x0000, 0x4000).
func (p *PPU) WriteVRAM(addr uint16, v uint8) {
	p.vram[addr] = v
}

// Writes to the OAM. The address is in the range [0x0000, 0x0100).
func (p *PPU) WriteOAM(addr uint16, v uint8) {
	p.oam[addr] = v
}
