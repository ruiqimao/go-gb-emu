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
	mode         Mode
	check0Enable bool
	check1Enable bool
	check2Enable bool
	lycEnable    bool

	// Memory.
	vram [0x4000]uint8
	oam  [0x100]uint8

	// Scanline counter.
	sc uint16
}

func NewPPU() *PPU {
	p := &PPU{}
	return p
}

// Do a PPU step. Returns how many clocks were used.
// The PPU runs at a resolution of 2 clocks, so always returns 2.
func (p *PPU) Step() int {
	// If the LCD is off, reset the PPU.
	if !p.lcdPower {
		p.reset()
		return 2
	}

	// Run a step.
	p.step()

	return 2
}

// Attach an MMU.
func (p *PPU) AttachMMU(mmu MMU) {
	p.mmu = mmu
}
