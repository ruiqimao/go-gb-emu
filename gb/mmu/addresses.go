package mmu

// Memory layout:
// 0000 - 00FF
//   Boot ROM - This gets turned off at the end of execution.
// 0000 - 3FFF
//   Cartridge ROM bank 0.
// 4000 - 7FFF
//   Cartridge ROM bank 1 - N.
// 8000 - 9FFF
//   Video RAM.
// A000 - BFFF
//   Cartridge RAM.
// C000 - CFFF
//   Work RAM bank 0.
// D000 - DFFF
//   Work RAM bank 1.
// E000 - FDFF
//   Mirror of C000 - DDFF.
// FE00 - FE9F
//   Sprite attribute table (OAM).
// FEA0 - FEFF
//   Empty. All reads are 0, and all writes are no-op.
// FF00 - FF7F
//   I/O registers.
// FF80 - FFFE
//   High RAM.
// FFFF
//   Interrupts enable register.
const (
	AddrBoot     = 0x0000
	AddrCartROM0 = 0x0000
	AddrCartROMN = 0x4000
	AddrVRAM     = 0x8000
	AddrCartRAM  = 0xa000
	AddrWRAM0    = 0xc000
	AddrWRAMN    = 0xd000
	AddrEcho     = 0xe000
	AddrOAM      = 0xfe00
	AddrEmpty    = 0xfea0
	AddrIO       = 0xff00
	AddrHRAM     = 0xff80
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
