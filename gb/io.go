package gb

import (
	"fmt"

	"github.com/ruiqimao/go-gb-emu/utils"
)

// Memory addresses for IO registers.
const (
	AddrJOYP = 0xff00 // Joypad.
	AddrSB   = 0xff01 // Serial byte.
	AddrSC   = 0xff02 // Serial control.
	// Unmapped: FF03.
	AddrDIV  = 0xff04 // Clock divider.
	AddrTIMA = 0xff05 // Timer value.
	AddrTMA  = 0xff06 // Timer reload.
	AddrTAC  = 0xff07 // Timer control.
	// Unmapped: FF08 - FF0E.
	AddrIF   = 0xff0f // Interrupt asserted.
	AddrNR10 = 0xff10 // Audio channel 1 sweep.
	AddrNR11 = 0xff11 // Audio channel 1 sound length/wave duty.
	AddrNR12 = 0xff12 // Audio channel 1 envelope.
	AddrNR13 = 0xff13 // Audio channel 1 frequency.
	AddrNR14 = 0xff14 // Audio channel 1 control.
	// Unmapped: FF15.
	AddrNR21 = 0xff16 // Audio channel 2 sound length/wave duty.
	AddrNR22 = 0xff17 // Audio channel 2 envelope.
	AddrNR23 = 0xff18 // Audio channel 2 frequency.
	AddrNR24 = 0xff19 // Audio channel 2 control.
	AddrNR30 = 0xff1a // Audio channel 3 enable.
	AddrNR31 = 0xff1b // Audio channel 3 sound length/wave duty.
	AddrNR32 = 0xff1c // Audio channel 3 envelope.
	AddrNR33 = 0xff1d // Audio channel 3 frequency.
	AddrNR34 = 0xff1e // Audio channel 3 control.
	// Unmapped: FF1F.
	AddrNR41 = 0xff20 // Audio channel 4 sound length.
	AddrNR42 = 0xff21 // Audio channel 4 volume.
	AddrNR43 = 0xff22 // Audio channel 4 frequency.
	AddrNR44 = 0xff23 // Audio channel 4 control.
	AddrNR50 = 0xff24 // Audio output mapping.
	AddrNR51 = 0xff25 // Audio channel mapping.
	AddrNR52 = 0xff26 // Audio channel control.
	// Unmapped: FF27 - FF2F.
	AddrWAVE = 0xff27 // Wave pattern.
	// AddrWAVE: FF28 - FF3F.
	AddrLCDC = 0xff40 // LCD control.
	AddrSTAT = 0xff41 // LCD status.
	AddrSCY  = 0xff42 // Background vertical scroll.
	AddrSCX  = 0xff43 // Background horizontal scroll.
	AddrLY   = 0xff44 // LCD Y coordinate.
	AddrLYC  = 0xff45 // LCD Y compare.
	AddrDMA  = 0xff46 // OAM DMA source address.
	AddrBGP  = 0xff47 // Background palette.
	AddrOBP0 = 0xff48 // OBJ palette 0.
	AddrOBP1 = 0xff49 // OBJ palette 1.
	AddrWY   = 0xff4a // Window Y coordinate.
	AddrWX   = 0xff4b // Window X coordinate.
	// Unmapped: FF4C - FF4F.
	AddrBOOT = 0xff50 // Boot ROM control.
	// Unmapped: FF51 - FF7F.
	// High RAM: FF80 - FFFE.
	AddrIE = 0xffff // Interrupts enabled.
)

// Whether the boot ROM is enabled.
func (m *Memory) BootRomEnabled() bool {
	return !utils.GetBit(m.Read(AddrBOOT), 0)
}

// Read an I/O register.
func (m *Memory) ReadIO(addr uint16) uint8 {
	switch addr {

	// TODO.

	case AddrDIV:
		// DIV register is upper 8 bits of internal counter.
		return uint8(m.gb.cpu.IC() >> 8)

	case AddrIF:
		// Upper 3 bits are always high.
		return m.io[addr-AddrIO] | 0xe0

	case AddrLCDC:
		// Retrieve from PPU.
		return m.gb.ppu.LCDC()

	case AddrSTAT:
		// Retrieve from PPU.
		return m.gb.ppu.STAT()

	case AddrLY:
		// Retrieve from PPU.
		return m.gb.ppu.LY()

	default:
		return m.io[addr-AddrIO]

	}
}

// Write to an I/O register.
func (m *Memory) WriteIO(addr uint16, v uint8) {
	if m.gb.Debugging() {
		if addr == AddrSC && v == 0x81 {
			fmt.Printf("%c", m.Read(AddrSB))
		}
	}

	switch addr {

	// TODO.

	case AddrDIV:
		// DIV maps directly to internal counter.
		lo := uint8(m.gb.cpu.IC())
		m.gb.cpu.SetIC(utils.CombineBytes(v, lo))

	case AddrTAC:
		// Only lower 3 bits are writable.
		m.io[addr-AddrIO] = v & 0x07

	case AddrLCDC:
		// Redirect to PPU.
		m.gb.ppu.SetLCDC(v)

	case AddrSTAT:
		// Redirect to PPU.
		m.gb.ppu.SetSTAT(v)

	case AddrLY:
		// Read only.

	default:
		m.io[addr-AddrIO] = v

	}
}
