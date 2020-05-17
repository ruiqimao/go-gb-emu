package ppu

import (
	"github.com/ruiqimao/go-gb-emu/utils"
)

type SpriteSize bool

const (
	SpriteSmall = false
	SpriteLarge = true
)

// OAM search constants.
const (
	MaxSpritesPerScanline = 10
)

// Flags for sprites.
const (
	FlagPalette  = 4
	FlagFlipX    = 5
	FlagFlipY    = 6
	FlagPriority = 7
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

func (p *PPU) newSprite(addr uint16) Sprite {
	posY := p.oam[addr]
	posX := p.oam[addr+0x1]
	tileN := p.oam[addr+0x2]
	flags := p.oam[addr+0x3]

	return Sprite{
		posY:     posY,
		posX:     posX,
		tileN:    tileN,
		palette:  utils.GetBit(flags, FlagPalette),
		flipX:    utils.GetBit(flags, FlagFlipX),
		flipY:    utils.GetBit(flags, FlagFlipY),
		priority: utils.GetBit(flags, FlagPriority),
	}
}

// Start OAM search.
func (p *PPU) startOAMSearch() {
	// Clear the OAM cache.
	p.oamCache = nil
}

// Run a step of OAM search.
func (p *PPU) stepOAMSearch() {
	// Do OAM search every other clock.
	if p.sc%2 == 0 {
		return
	}

	// If the OAM cache is full, exit.
	if len(p.oamCache) == MaxSpritesPerScanline {
		return
	}

	// Decode the next entry into a Sprite.
	sprite := p.newSprite(p.sc * 2)

	// Determine if the sprite is in the current scanline.
	inFrame := sprite.posX != 0
	inUpperBound := p.ly+0x10 >= sprite.posY
	inLowerBound := p.ly+0x10 < sprite.posY+p.spriteHeight()

	// If the sprite is visible, add it to the OAM cache.
	if inFrame && inUpperBound && inLowerBound {
		p.oamCache = append(p.oamCache, sprite)
	}
}
