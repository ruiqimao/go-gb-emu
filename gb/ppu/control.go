package ppu

import (
	"github.com/ruiqimao/go-gb-emu/utils"
)

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
	FlagHBLCheck      = 5
	FlagVBLCheck      = 4
	FlagOAMCheck      = 3
	FlagLYCComparison = 2
)

type Mode int

const (
	ModeHBlank   Mode = 0
	ModeVBlank        = 1
	ModeOAM           = 2
	ModeTransfer      = 3
)

// Timing constants.
const (
	OAMClocks = 80
	HClocks   = 456
	VLines    = 154
)

// Buffer constants.
const (
	FrameWidth  = 160
	FrameHeight = 144
)

// Resets the PPU.
func (p *PPU) reset() {
	p.sc = 0
	p.ly = 0
}

// Push a frame to the MMU.
func (p *PPU) pushFrame() {
	// Make a copy of the frame.
	frame := make([]byte, len(p.frame))
	copy(frame, p.frame[:])

	// If the LCD is off, send nil instead.
	if !p.lcdPower {
		frame = nil
	}

	// Try to push the frame. If the channel is full, drop the frame.
	select {
	case p.F <- frame:
	default:
	}
}

// Update the STAT signal.
func (p *PPU) updateSTAT() {
	statSig := uint8(0)

	// Calculate new signal.
	if p.lycCheck && p.ly == p.lyc {
		statSig = 0x1
	}
	if p.hblCheck && p.mode == ModeHBlank {
		statSig = 0x1
	}
	if (p.vblCheck || p.oamCheck) && p.mode == ModeVBlank { // Possible bug in Game Boy hardware.
		statSig = 0x1
	}
	if p.oamCheck && p.mode == ModeOAM {
		statSig = 0x1
	}

	// If there is a rising edge, trigger the STAT interrupt.
	if p.statSig == 0x0 && statSig == 0x1 {
		p.interruptSTAT()
	}

	p.statSig = statSig
}

// Get the LCDC register.
func (p *PPU) LCDC() uint8 {
	lcdc := uint8(0)
	lcdc = utils.SetBit(lcdc, FlagBgEnable, p.bgEnable)
	lcdc = utils.SetBit(lcdc, FlagSpritesEnable, p.spritesEnable)
	lcdc = utils.SetBit(lcdc, FlagSpriteSize, bool(p.spriteSize))
	lcdc = utils.SetBit(lcdc, FlagBgTileMap, bool(p.bgMap))
	lcdc = utils.SetBit(lcdc, FlagTileset, bool(p.tileset))
	lcdc = utils.SetBit(lcdc, FlagWindowEnable, p.winEnable)
	lcdc = utils.SetBit(lcdc, FlagWindowTileMap, bool(p.winMap))
	lcdc = utils.SetBit(lcdc, FlagLCDPower, p.lcdPower)
	return lcdc
}

// Set the LCDC register.
func (p *PPU) SetLCDC(v uint8) {
	p.bgEnable = utils.GetBit(v, FlagBgEnable)
	p.spritesEnable = utils.GetBit(v, FlagSpritesEnable)
	p.spriteSize = SpriteSize(utils.GetBit(v, FlagSpriteSize))
	p.bgMap = TileMap(utils.GetBit(v, FlagBgTileMap))
	p.tileset = Tileset(utils.GetBit(v, FlagTileset))
	p.winEnable = utils.GetBit(v, FlagWindowEnable)
	p.winMap = TileMap(utils.GetBit(v, FlagWindowTileMap))
	p.lcdPower = utils.GetBit(v, FlagLCDPower)
}

// Get the STAT register.
func (p *PPU) STAT() uint8 {
	stat := uint8(p.mode) & 0x3
	stat = utils.SetBit(stat, FlagLYCComparison, p.ly == p.lyc)
	stat = utils.SetBit(stat, FlagHBLCheck, p.hblCheck)
	stat = utils.SetBit(stat, FlagVBLCheck, p.vblCheck)
	stat = utils.SetBit(stat, FlagOAMCheck, p.oamCheck)
	stat = utils.SetBit(stat, FlagLYCCheck, p.lycCheck)
	stat |= 0x80 // Bit 7 always reads 1.
	return stat
}

// Set the STAT register.
func (p *PPU) SetSTAT(v uint8) {
	p.hblCheck = utils.GetBit(v, FlagHBLCheck)
	p.vblCheck = utils.GetBit(v, FlagVBLCheck)
	p.oamCheck = utils.GetBit(v, FlagOAMCheck)
	p.lycCheck = utils.GetBit(v, FlagLYCCheck)
}

// Get the LY register.
func (p *PPU) LY() uint8 {
	return p.ly
}

// Set the LY register. This does nothing, as LY is not writable.
func (p *PPU) SetLY(v uint8) {
	// LY is read-only.
}

// Get the LYC register.
func (p *PPU) LYC() uint8 {
	return p.lyc
}

// Set the LYC register.
func (p *PPU) SetLYC(v uint8) {
	p.lyc = v
}
