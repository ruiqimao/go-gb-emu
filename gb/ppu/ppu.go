package ppu

// Pixel processing unit.
type PPU struct {
	mmu MMU

	// Registers.
	scy  uint8
	scx  uint8
	ly   uint8
	lyc  uint8
	bgp  uint8
	obp0 uint8
	obp1 uint8
	wx   uint8
	wy   uint8

	// LCDC flags.
	lcdPower      bool
	winMap        TileMap
	winEnable     bool
	tileset       Tileset
	bgMap         TileMap
	spriteSize    SpriteSize
	spritesEnable bool
	bgEnable      bool

	// STAT flags.
	mode     Mode
	hblCheck bool
	vblCheck bool
	oamCheck bool
	lycCheck bool

	// STAT signal.
	statSig uint8

	// Memory.
	vram [0x4000]uint8
	oam  [0x100]uint8

	// Scanline counter.
	sc uint16

	// OAM cache.
	oamCache []Sprite

	// Pixel transfer state.
	fetcher *Fetcher
	lx      uint8
	frame   [FrameWidth * FrameHeight]uint8

	// Latest rendered frame.
	F chan []uint8
}

func NewPPU() *PPU {
	p := &PPU{
		F: make(chan []uint8, 1),
	}
	p.fetcher = NewFetcher(p)
	return p
}

// Do a PPU step. Returns how many clocks were used.
// The PPU runs at a resolution of 1 clock, so always returns 1.
func (p *PPU) Step() int {
	// If the LCD is off, reset the PPU.
	if !p.lcdPower {
		p.reset()
		return 1
	}

	// Update the mode.
	switch {
	case p.ly < FrameHeight && p.sc == 0:
		p.mode = ModeOAM
		p.startOAMSearch()
	case p.ly < FrameHeight && p.sc == OAMClocks:
		p.mode = ModeTransfer
		p.startPixelTransfer()
	case p.ly == FrameHeight && p.sc == 0:
		p.mode = ModeVBlank
		p.pushFrame()
		p.mmu.RequestInterrupt(InterruptVBlank)
	}

	// Update the STAT signal.
	p.updateSTAT()

	// Perform a step of OAM search or pixel transfer.
	if p.mode == ModeOAM {
		p.stepOAMSearch()
	}
	if p.mode == ModeTransfer {
		p.stepPixelTransfer()
	}

	// Increment the scanline counter and scanline.
	p.sc = (p.sc + 1) % HClocks
	if p.sc == 0 {
		p.ly = (p.ly + 1) % VLines
	}

	return 1
}

// Attach an MMU.
func (p *PPU) AttachMMU(mmu MMU) {
	p.mmu = mmu
}
