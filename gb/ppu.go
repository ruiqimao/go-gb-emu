package gb

// LCDC flags.
const (
	FlagLCDPower      = 7
	FlagWindowTileMap = 6
	FlagWindowEnable  = 5
	FlagBgTileset     = 4
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

// Timing constants.
const (
	HSteps = 114
	VLines = 154
)

// Buffer constants.
const (
	FrameWidth = 160
	FrameHeight = 144
)

// Pixel processing unit.
type PPU struct {
	gb *GameBoy

	sc uint16 // Scanline counter.
	ly uint8 // Current scanline.

	lcdc uint8 // LCD control.
	stat uint8 // LCD status.

	// Current frame.
	frame [FrameWidth*FrameHeight]uint8
}

func NewPPU(gb *GameBoy) *PPU {
	p := &PPU{
		gb: gb,
	}
	return p
}

// Catch the PPU up to the CPU.
func (p *PPU) Update(cycles int) {
	// Process 4 cycles at a time.
	for cycles > 0 {

		// Update the mode.
		// This is done at the beginning of the step because mode changes are delayed by a step.
		switch {
		case p.ly < 144 && p.sc == 0:
			p.setMode(ModeOAM)
		case p.ly < 144 && p.sc == 20:
			p.setMode(ModeTransfer)
		case p.ly == 144 && p.sc == 0:
			p.setMode(ModeVBlank)
			p.pushFrame()
			p.gb.cpu.RequestInterrupt(IntVBlank)
		}

		// If the PPU is not in a blanking mode, execute a step.
		mode := p.Mode()
		if mode == ModeOAM {
			p.stepOAM()
		}
		if mode == ModeTransfer {
			p.stepTransfer()
		}

		// Increment the scanline counter and scanline.
		p.sc = (p.sc + 1) % HSteps
		if p.sc == 0 {
			p.ly = (p.ly + 1) % VLines
		}

		cycles -= 4
	}
}

// Execute a step of OAM searching.
func (p *PPU) stepOAM() {
	// TODO.
}

// Execute a step of pixel transfer.
func (p *PPU) stepTransfer() {
	// TODO.
}

// Push the current frame into the channel.
func (p *PPU) pushFrame() {
	// Make a copy of the frame.
	frame := make([]byte, len(p.frame))
	copy(frame, p.frame[:])

	// Try to push the frame. If the channel is full, drop the frame.
	select {
	case p.gb.F <- frame:
	default:
	}
}

// Set the mode.
func (p *PPU) setMode(m uint8) {
	p.stat = p.stat & 0xf8 | m & 0x03
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
	p.stat = p.stat & 0x0f | v & 0xf0
}

// Get the mode.
func (p *PPU) Mode() uint8 {
	return p.stat & 0x03
}

// Get the value of LY.
func (p *PPU) LY() uint8 {
	return p.ly
}
