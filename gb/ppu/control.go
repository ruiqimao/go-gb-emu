package ppu

// LCDC flags.
const (
	FlagLCDPower      = 7
	FlagWindowTileMap = 6
	FlagWindowEnable  = 5
	FlagTileset       = 4
	FlagBgTileMap     = 3
	FlagSpriteSize    = 2
	FlagSpritesEnable = 1
	FlagBgEnable      = 0
)

// STAT flags.
const (
	FlagLYCCheck      = 6
	FlagMode2Check    = 5
	FlagMode1Check    = 4
	FlagMode0Check    = 3
	FlagLYCComparison = 2
	ModeHBlank        = 0
	ModeVBlank        = 1
	ModeOAM           = 2
	ModeTransfer      = 3
)

func (p *PPU) LCDC() uint8 {
	return p.lcdc
}

func (p *PPU) SetLCDC(v uint8) {
	// TODO.
}

func (p *PPU) STAT() uint8 {
	return p.stat
}

func (p *PPU) SetSTAT(v uint8) {
	// TODO.
}

func (p *PPU) LY() uint8 {
	return p.ly
}

func (p *PPU) SetLY(v uint8) {
	// TODO.
}

func (p *PPU) LYC() uint8 {
	return p.lyc
}

func (p *PPU) SetLYC(v uint8) {
	// TODO.
}
