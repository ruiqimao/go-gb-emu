package gb

import (
	"github.com/ruiqimao/go-gb-emu/utils"
)

// OAM search constants.
const (
	MaxSpritesPerScanline = 10
)

// Flags for sprites.
const (
	SpritePalette  = 4
	SpriteFlipX    = 5
	SpriteFlipY    = 6
	SpritePriority = 7
)

// A Sprite is a decoded entry in the OAM.
type Sprite struct {
	posY     uint8
	posX     uint8
	tileN    uint8
	palette  bool
	flipX    bool
	flipY    bool
	priority bool
}

func NewSprite(ppu *PPU, addr uint16) Sprite {
	posY := ppu.oam[addr]
	posX := ppu.oam[addr+0x1]
	tileN := ppu.oam[addr+0x2]
	flags := ppu.oam[addr+0x3]

	return Sprite{
		posY:     posY,
		posX:     posX,
		tileN:    tileN,
		palette:  utils.GetBit(flags, SpritePalette),
		flipX:    utils.GetBit(flags, SpriteFlipX),
		flipY:    utils.GetBit(flags, SpriteFlipY),
		priority: utils.GetBit(flags, SpritePriority),
	}
}

// Start OAM search.
func (p *PPU) startOAMSearch() {
	// Clear the OAM cache.
	p.oamCache = nil
}

// Execute a step of OAM searching.
func (p *PPU) stepOAMSearch() {
	// Scan two OAM entries per step.
	for i := 0; i < 2; i++ {
		// If the OAM cache is full, exit.
		if len(p.oamCache) == MaxSpritesPerScanline {
			break
		}

		// Get the address of the entry in OAM.
		entryAddr := p.sc*8 + uint16(i*4)

		// Decode the entry into a Sprite.
		sprite := NewSprite(p, entryAddr)

		// Determine if the sprite is in the current scanline.
		inFrame := sprite.posX != 0
		inUpperBound := p.ly+0x10 >= sprite.posY
		inLowerBound := p.ly+0x10 < sprite.posY+p.SpriteHeight()

		// If the sprite is visible, add it to the OAM cache.
		if inFrame && inUpperBound && inLowerBound {
			p.oamCache = append(p.oamCache, sprite)
		}
	}
}
