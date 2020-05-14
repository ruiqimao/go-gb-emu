package gb

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
	FlagMode2Check    = 5
	FlagMode1Check    = 4
	FlagMode0Check    = 3
	FlagLYCComparison = 2
	ModeHBlank        = 0
	ModeVBlank        = 1
	ModeOAM           = 2
	ModeTransfer      = 3
)

// DMA constants.
const (
	DMACycles = 644
)

// Timing constants.
const (
	HSteps = 114
	VLines = 154
)

// Buffer constants.
const (
	FrameWidth  = 160
	FrameHeight = 144
)

// Pixel processing unit.
type PPU struct {
	gb *GameBoy

	sc uint16 // Scanline counter.
	lx uint8  // Current position in scanline (0-160).
	ly uint8  // Current scanline.

	lcdc uint8 // LCD control.
	stat uint8 // LCD status.

	// VRAM and OAM.
	vram [0x4000]uint8
	oam  [0x100]uint8

	// Internal STAT signal.
	statSignal uint8

	// DMA address and remaining cycles.
	dmaAddr   uint16
	dmaCycles uint16

	// OAM cache for OAM searching and pixel transfer.
	oamCache []Sprite

	// Register snapshot.
	snapSCX  uint8
	snapSCY  uint8
	snapBGP  uint8
	snapOBP0 uint8
	snapOBP1 uint8
	snapWX   uint8
	snapWY   uint8

	// Pixel fetcher.
	fetcher *Fetcher

	// Current frame.
	frame [FrameWidth * FrameHeight]uint8
}

func NewPPU(gb *GameBoy) *PPU {
	p := &PPU{
		gb: gb,
	}
	p.fetcher = NewFetcher(p)
	return p
}

// Catch the PPU up to the CPU.
func (p *PPU) Update(cycles int) {
	// Process 4 cycles at a time.
	for cycles > 0 {

		// If DMA is being done, update the number of cycles left.
		if p.dmaCycles > 0 {
			p.dmaCycles -= 4
		}

		// If the LCD is off, reset LY and exit.
		if !p.LCDPower() {
			p.ly = 0
			break
		}

		// Update the mode.
		// This is done at the beginning of the step because mode changes are delayed by a step.
		switch {
		case p.ly < FrameHeight && p.sc == 0:
			p.setMode(ModeOAM)
			p.startOAMSearch()
		case p.ly < FrameHeight && p.sc == 20:
			p.setMode(ModeTransfer)
			p.startPixelTransfer()
		case p.ly == FrameHeight && p.sc == 0:
			p.setMode(ModeVBlank)
			p.pushFrame()
			p.gb.cpu.RequestInterrupt(IntVBlank)
		}

		// Update the STAT signal.
		p.updateSTAT()

		// If the PPU is not in a blanking mode, execute a step.
		mode := p.Mode()
		if mode == ModeOAM {
			p.stepOAMSearch()
		}
		if mode == ModeTransfer {
			p.stepPixelTransfer()
		}

		// Increment the scanline counter and scanline.
		p.sc = (p.sc + 1) % HSteps
		if p.sc == 0 {
			p.ly = (p.ly + 1) % VLines
		}

		cycles -= 4
	}
}

// Set the mode.
func (p *PPU) setMode(m uint8) {
	p.stat = p.stat&0xf8 | m&0x03
}

// Update the STAT signal.
func (p *PPU) updateSTAT() {
	statSignal := uint8(0)

	mode := p.Mode()
	lycCheck := utils.GetBit(p.stat, FlagLYCCheck)
	hblCheck := utils.GetBit(p.stat, FlagMode0Check)
	vblCheck := utils.GetBit(p.stat, FlagMode1Check)
	oamCheck := utils.GetBit(p.stat, FlagMode2Check)

	// Calculate new signal.
	if lycCheck && p.ly == p.gb.mem.Read(AddrLYC) {
		statSignal = 0x1
	}
	if hblCheck && mode == ModeHBlank {
		statSignal = 0x1
	}
	if (vblCheck || oamCheck) && mode == ModeVBlank { // Possible bug in Game Boy hardware.
		statSignal = 0x1
	}
	if oamCheck && mode == ModeOAM {
		statSignal = 0x1
	}

	// If there is a rising edge, trigger the STAT interrupt.
	if p.statSignal == 0x0 && statSignal == 0x1 {
		p.gb.cpu.RequestInterrupt(IntStat)
	}

	p.statSignal = statSignal
}

// Read a byte from the VRAM.
func (p *PPU) ReadVRAM(addr uint16) uint8 {
	if p.LCDPower() && p.Mode() == ModeTransfer {
		return 0x00
	}
	return p.vram[addr]
}

// Write a byte to the VRAM.
func (p *PPU) WriteVRAM(addr uint16, v uint8) {
	if p.LCDPower() && p.Mode() == ModeTransfer {
		return
	}
	p.vram[addr] = v
}

// Read a byte from the OAM.
func (p *PPU) ReadOAM(addr uint16) uint8 {
	if p.LCDPower() && p.Mode() == ModeTransfer || p.Mode() == ModeOAM {
		return 0x00
	}
	return p.oam[addr]
}

// Write a byte to the OAM.
func (p *PPU) WriteOAM(addr uint16, v uint8) {
	if p.LCDPower() && p.Mode() == ModeTransfer || p.Mode() == ModeOAM {
		return
	}
	p.oam[addr] = v
}

// Get the value of LCDC.
func (p *PPU) LCDC() uint8 {
	return p.lcdc
}

// Set the value of LCDC.
func (p *PPU) SetLCDC(v uint8) {
	p.lcdc = v
}

// Get the value of STAT.
func (p *PPU) STAT() uint8 {
	return p.lcdc
}

// Set the value of STAT.
func (p *PPU) SetSTAT(v uint8) {
	// Only set writable values.
	p.stat = p.stat&0x0f | v&0xf0
}

// Get the mode.
func (p *PPU) Mode() uint8 {
	return p.stat & 0x03
}

// Get the value of LY.
func (p *PPU) LY() uint8 {
	return p.ly
}

// Get the value of DMA.
func (p *PPU) DMA() uint8 {
	return uint8(p.dmaAddr / 0x100)
}

// Set the value of DMA.
func (p *PPU) SetDMA(v uint8) {
	p.dmaAddr = uint16(v) * 0x100

	// Perform the entire DMA in one go. We can do this because the CPU should not write any changes
	// to memory other than HRAM during DMA, so it is safe to consider memory to be static during the
	// entire duration of DMA.
	for i := 0; i < 0x100; i++ {
		p.oam[i] = p.gb.mem.Read(p.dmaAddr + uint16(i))
	}

	// Start the DMA duration.
	p.dmaCycles = DMACycles
}

// Get whether DMA is being run.
func (p *PPU) InDMA() bool {
	return p.dmaCycles > 0
}

// Get whether the LCD is on.
func (p *PPU) LCDPower() bool {
	return utils.GetBit(p.lcdc, FlagLCDPower)
}

// Get the address of the tile map the window is using relative to VRAM.
func (p *PPU) WindowTileMap() uint16 {
	if utils.GetBit(p.lcdc, FlagWindowTileMap) {
		return 0x9c00 - AddrVRAM
	} else {
		return 0x9800 - AddrVRAM
	}
}

// Get whether the window is enabled.
func (p *PPU) WindowEnabled() bool {
	return utils.GetBit(p.lcdc, FlagWindowEnable)
}

// Get the data of a tile given the tile id and offset.
// If the tile is a sprite tile, unsigned addressing mode will be forced.
// The offset must be in the range [0, 16).
func (p *PPU) Tile(id uint8, offset uint8, sprite bool) uint8 {
	if utils.GetBit(p.lcdc, FlagTileset) || sprite {
		// Unsigned addressing mode.
		return p.vram[0x8000-AddrVRAM+uint16(id)*16+uint16(offset)]
	} else {
		// Signed addressing mode.
		return p.vram[uint16(0x9000-AddrVRAM+int32(int8(id))*16)+uint16(offset)]
	}
}

// Get the address of the tile map the background is using relative to VRAM.
func (p *PPU) BgTileMap() uint16 {
	if utils.GetBit(p.lcdc, FlagBgTileMap) {
		return 0x9c00 - AddrVRAM
	} else {
		return 0x9800 - AddrVRAM
	}
}

// Get the height of the sprites.
func (p *PPU) SpriteHeight() uint8 {
	if utils.GetBit(p.lcdc, FlagSpriteSize) {
		return 16
	} else {
		return 8
	}
}

// Get whether sprites are enabled.
func (p *PPU) SpritesEnabled() bool {
	return utils.GetBit(p.lcdc, FlagSpritesEnable)
}

// Get whether the background is enabled.
func (p *PPU) BackgroundEnabled() bool {
	return utils.GetBit(p.lcdc, FlagBgEnable)
}
