package gb

// Start pixel transfer.
func (p *PPU) startPixelTransfer() {
	// Take a snapshot of the registers.
	p.snapSCX = p.gb.mem.Read(AddrSCX)
	p.snapSCY = p.gb.mem.Read(AddrSCY)
	p.snapBGP = p.gb.mem.Read(AddrBGP)
	p.snapOBP0 = p.gb.mem.Read(AddrOBP0)
	p.snapOBP1 = p.gb.mem.Read(AddrOBP1)
	p.snapWX = p.gb.mem.Read(AddrWX)
	// Window Y is actually not dynamic and should be snapped at the beginning of the frame instead.
	// However, it is done here for simplicity.
	p.snapWY = p.gb.mem.Read(AddrWY)

	// Reset fetcher.
	p.fetcher.Reset(p.snapSCX, p.ly+p.snapSCY, p.BgTileMap())

	// Reset the x position.
	p.lx = 0
}

// Execute a step of pixel transfer.
func (p *PPU) stepPixelTransfer() {
	for c := 0; c < 4; c++ {
		// Check for a window.
		if p.WindowEnabled() && p.lx == p.snapWX && p.ly >= p.snapWY {
			p.fetcher.Reset(0, p.ly-p.snapWY, p.WindowTileMap())
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
			p.setMode(ModeHBlank)
			break
		}

		// Step the fetcher every 2 cycles.
		if c%2 == 0 {
			p.fetcher.Step()
		}
	}
}

// Push the current frame into the channel.
func (p *PPU) pushFrame() {
	// Make a copy of the frame.
	frame := make([]byte, len(p.frame))
	copy(frame, p.frame[:])

	// If the LCD is off, send nil instead.
	if !p.LCDPower() {
		frame = nil
	}

	// Try to push the frame. If the channel is full, drop the frame.
	select {
	case p.gb.F <- frame:
	default:
	}
}

// Resolve the color of a pixel.
func (p *PPU) Resolve(px Pixel) uint8 {
	var palette uint8
	switch px.src {
	case PixelSrc00, PixelSrc10:
		palette = p.snapOBP0
	case PixelSrc01, PixelSrc11:
		palette = p.snapOBP1
	case PixelSrcBG:
		palette = p.snapBGP
	}
	return (palette >> (px.data * 2)) & 0x3
}
