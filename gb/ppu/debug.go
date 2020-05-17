package ppu

// The functions in this file should be used only for debugging purposes.

func (p *PPU) VRAM() []uint8 {
	return p.vram[:]
}

func (p *PPU) OAM() []uint8 {
	return p.oam[:]
}

func (p *PPU) BgMapAddr() uint16 {
	return p.bgMapAddr()
}

func (p *PPU) WinMapAddr() uint16 {
	return p.winMapAddr()
}

func (p *PPU) TileData(id uint8, offset uint8, sprite bool) uint8 {
	return p.tileData(id, offset, sprite)
}

func (p *PPU) Resolve(px Pixel) uint8 {
	return p.resolve(px)
}
