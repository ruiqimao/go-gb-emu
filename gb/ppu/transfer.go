package ppu

// Start pixel transfer.
func (p *PPU) startPixelTransfer() {
	// Reset fetcher.
	p.fetcher.Reset(p.scx, p.ly+p.scy, p.bgMapAddr())

	// Reset the x position.
	p.lx = 0
}

// Run a step of pixel transfer.
func (p *PPU) stepPixelTransfer() {
	// Check for a window.
	if p.winEnable && p.lx == p.wx && p.ly >= p.wy {
		p.fetcher.Reset(0, p.ly-p.wy, p.winMapAddr())
	}

	// Check for a sprite.
	// TODO.

	// Try to pop a pixel off the fetcher.
	if color, ok := p.fetcher.Pop(); ok {
		// Put the color in the frame.
		p.frame[int(p.ly)*FrameWidth+int(p.lx)] = color

		// Move to the next pixel.
		p.lx++
	}

	// If all pixels in the scanline have been filled, move to HBlank.
	if p.lx == FrameWidth {
		p.mode = ModeHBlank
		return
	}

	// Step the fetcher every 2 cycles.
	if p.sc%2 == 0 {
		p.fetcher.Step()
	}
}

// Resolve the color of a pixel.
func (p *PPU) resolve(px Pixel) uint8 {
	var palette uint8
	switch {
	case px.bg:
		palette = p.bgp
	case px.palette == 0:
		palette = p.obp0
	case px.palette == 1:
		palette = p.obp1
	}
	return (palette >> (px.data * 2)) & 0x3
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
