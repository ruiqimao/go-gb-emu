package ppu

type TileMap bool

const (
	TileMap0 TileMap = false
	TileMap1         = true
)

type Tileset bool

const (
	Tileset0 Tileset = false
	Tileset1         = true
)

// Get the address of the tile map the window is using relative to VRAM.
func (p *PPU) winMapAddr() uint16 {
	if p.winMap == TileMap1 {
		return 0x1c00
	} else {
		return 0x1800
	}
}

// Get the data of a tile given the tile id and offset.
// If the tile is a sprite tile, unsigned addressing mode will be forced.
// The offset must be in the range [0, 16).
func (p *PPU) tileData(id uint8, offset uint8, sprite bool) uint8 {
	if p.tileset == Tileset1 || sprite {
		// Unsigned addressing mode.
		return p.vram[uint16(id)*16+uint16(offset)]
	} else {
		// Signed addressing mode.
		return p.vram[uint16(0x1000+int32(int8(id))*16)+uint16(offset)]
	}
}

// Get the address of the tile map the background is using relative to VRAM.
func (p *PPU) bgMapAddr() uint16 {
	if p.bgMap == TileMap1 {
		return 0x1c00
	} else {
		return 0x1800
	}
}

// Get the height of the sprites.
func (p *PPU) spriteHeight() uint8 {
	if p.spriteSize == SpriteLarge {
		return 16
	} else {
		return 8
	}
}

func (p *PPU) BGP() uint8 {
	return p.bgp
}

func (p *PPU) SetBGP(v uint8) {
	p.bgp = v
}

func (p *PPU) OBP0() uint8 {
	return p.obp0
}

func (p *PPU) SetOBP0(v uint8) {
	p.obp0 = v
}

func (p *PPU) OBP1() uint8 {
	return p.obp1
}

func (p *PPU) SetOBP1(v uint8) {
	p.obp1 = v
}
