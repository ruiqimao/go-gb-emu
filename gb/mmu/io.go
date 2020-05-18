package mmu

// Handle a read.
func (m *MMU) read(addr uint16) uint8 {
	switch {

	// Boot ROM.
	case addr < 0x0100 && m.bootrom != nil && m.bootrom.BOOT() == 0x0:
		return m.bootrom.Read(addr)

	// Cartridge ROM banks.
	case addr >= AddrCartROM0 && addr < AddrVRAM && m.cartridge != nil:
		return m.cartridge.ReadROM(addr)

	// Video RAM.
	case addr >= AddrVRAM && addr < AddrCartRAM && m.ppu != nil:
		return m.ppu.ReadVRAM(addr - AddrVRAM)

	// Cartridge RAM.
	case addr >= AddrCartRAM && addr < AddrWRAM0:
		return m.cartridge.ReadRAM(addr - AddrCartRAM)

	// Work RAM banks 0 and 1.
	case addr >= AddrWRAM0 && addr < AddrEcho:
		return m.wram[addr-AddrWRAM0]

	// Mirror of C000 - DDFF.
	case addr >= AddrEcho && addr < AddrOAM:
		return m.read(addr - 0x2000)

	// Sprite attribute table.
	case addr >= AddrOAM && addr < AddrEmpty && m.ppu != nil:
		return m.ppu.ReadOAM(addr - AddrOAM)

	// Empty.
	case addr >= AddrEmpty && addr < AddrIO:
		return 0x00

	// I/O registers.
	case addr >= AddrIO && addr < AddrHRAM:
		return m.readIO(addr)

	// High RAM.
	case addr >= AddrHRAM && addr < AddrIE:
		return m.hram[addr-AddrHRAM]

	// Interrupt enable register.
	case addr == AddrIE && m.cpu != nil:
		return m.cpu.IE()

	}

	return 0x00
}

// Handle a write.
func (m *MMU) write(addr uint16, v uint8) {
	switch {

	// Cartridge ROM.
	case addr < AddrVRAM && m.cartridge != nil:
		m.cartridge.WriteROM(addr, v)

	// Video RAM.
	case addr >= AddrVRAM && addr < AddrCartRAM && m.ppu != nil:
		m.ppu.WriteVRAM(addr-AddrVRAM, v)

	// Cartridge RAM.
	case addr >= AddrCartRAM && addr < AddrWRAM0 && m.cartridge != nil:
		m.cartridge.WriteROM(addr-AddrCartRAM, v)

	// Work RAM banks 0 and 1.
	case addr >= AddrWRAM0 && addr < AddrEcho:
		m.wram[addr-AddrWRAM0] = v

	// Mirror of C000 - DDFF.
	case addr >= AddrEcho && addr < AddrOAM:
		m.write(addr-0x2000, v)

	// Sprite attribute table.
	case addr >= AddrOAM && addr < AddrIO && m.ppu != nil:
		m.ppu.WriteOAM(addr-AddrOAM, v)

	// I/O registers.
	case addr >= AddrIO && addr < AddrHRAM:
		m.writeIO(addr, v)

	// High RAM.
	case addr >= AddrHRAM && addr < AddrIE:
		m.hram[addr-AddrHRAM] = v

	// Interrupt enable register.
	case addr == AddrIE && m.cpu != nil:
		m.cpu.SetIE(v)

	}
}

// Handle a read to an IO register.
func (m *MMU) readIO(addr uint16) uint8 {
	if m.joypad != nil && addr == AddrJOYP {
		return m.joypad.JOYP()
	}

	if m.cpu != nil {
		switch addr {
		case AddrDIV:
			return m.cpu.DIV()
		case AddrTIMA:
			return m.cpu.TIMA()
		case AddrTMA:
			return m.cpu.TMA()
		case AddrTAC:
			return m.cpu.TAC()
		case AddrIF:
			return m.cpu.IF()
		}
	}

	if m.ppu != nil {
		switch addr {
		case AddrLCDC:
			return m.ppu.LCDC()
		case AddrSTAT:
			return m.ppu.STAT()
		case AddrSCY:
			return m.ppu.SCY()
		case AddrSCX:
			return m.ppu.SCX()
		case AddrLY:
			return m.ppu.LY()
		case AddrLYC:
			return m.ppu.LYC()
		case AddrDMA:
			return m.DMA()
		case AddrBGP:
			return m.ppu.BGP()
		case AddrOBP0:
			return m.ppu.OBP0()
		case AddrOBP1:
			return m.ppu.OBP1()
		case AddrWY:
			return m.ppu.WY()
		case AddrWX:
			return m.ppu.WX()
		}
	}

	if m.bootrom != nil && addr == AddrBOOT {
		return m.bootrom.BOOT()
	}

	return 0x00
}

// Handle a write to an IO register.
func (m *MMU) writeIO(addr uint16, v uint8) {
	if m.joypad != nil && addr == AddrJOYP {
		m.joypad.SetJOYP(v)
	}

	if m.cpu != nil {
		switch addr {
		case AddrDIV:
			m.cpu.SetDIV(v)
		case AddrTIMA:
			m.cpu.SetTIMA(v)
		case AddrTMA:
			m.cpu.SetTMA(v)
		case AddrTAC:
			m.cpu.SetTAC(v)
		case AddrIF:
			m.cpu.SetIF(v)
		}
	}

	if m.ppu != nil {
		switch addr {
		case AddrLCDC:
			m.ppu.SetLCDC(v)
		case AddrSTAT:
			m.ppu.SetSTAT(v)
		case AddrSCY:
			m.ppu.SetSCY(v)
		case AddrSCX:
			m.ppu.SetSCX(v)
		case AddrLY:
			m.ppu.SetLY(v)
		case AddrLYC:
			m.ppu.SetLYC(v)
		case AddrDMA:
			m.SetDMA(v)
		case AddrBGP:
			m.ppu.SetBGP(v)
		case AddrOBP0:
			m.ppu.SetOBP0(v)
		case AddrOBP1:
			m.ppu.SetOBP1(v)
		case AddrWY:
			m.ppu.SetWY(v)
		case AddrWX:
			m.ppu.SetWX(v)
		}
	}

	if m.bootrom != nil && addr == AddrBOOT {
		m.bootrom.SetBOOT(v)
	}
}
