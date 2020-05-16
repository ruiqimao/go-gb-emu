package ppu

// Start pixel transfer.
func (p *PPU) startPixelTransfer() {
	// TODO.
}

// Run a step of pixel transfer.
func (p *PPU) stepPixelTransfer() {
	// TODO.
}

func (p *PPU) SCY() uint8 {
	return p.scy
}

func (p *PPU) SetSCY(v uint8) {
	p.scy = v
}

func (p *PPU) SCX() uint8 {
	return p.scx
}

func (p *PPU) SetSCX(v uint8) {
	p.scx = v
}

func (p *PPU) WY() uint8 {
	return p.wy
}

func (p *PPU) SetWY(v uint8) {
	p.wy = v
}

func (p *PPU) WX() uint8 {
	return p.wx
}

func (p *PPU) SetWX(v uint8) {
	p.wx = v
}
