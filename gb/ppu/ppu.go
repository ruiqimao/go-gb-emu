package ppu

// Pixel processing unit.
type PPU struct {
	mmu MMU

	// Registers.
	lcdc uint8
	stat uint8
	scy  uint8
	scx  uint8
	ly   uint8
	lyc  uint8
	bgp  uint8
	obp0 uint8
	obp1 uint8
	wx   uint8
	wy   uint8

	// Memory.
	vram [0x4000]uint8
	oam  [0x100]uint8
}
