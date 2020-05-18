package mmu

// Handle a read.
func (m *MMU) read(addr uint16) uint8 {
	switch {

	// Boot ROM.
	case addr < 0x0100 && m.bootrom != nil && m.bootrom.BOOT() == 0x0:
		return m.bootrom.Read(addr)

	// Cartridge ROM banks.
	case addr >= AddrCartROM0 && addr < AddrVRAM && m.cart != nil:
		return m.cart.ReadROM(addr)

	// Video RAM.
	case addr >= AddrVRAM && addr < AddrCartRAM && m.ppu != nil:
		return m.ppu.ReadVRAM(addr - AddrVRAM)

	// Cartridge RAM.
	case addr >= AddrCartRAM && addr < AddrWRAM0:
		return m.cart.ReadRAM(addr - AddrCartRAM)

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
	case addr >= addrHRAM && addr < AddrIE:
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
	case addr < AddrVRAM && m.cart != nil:
		m.gb.cart.WriteROM(addr, v)

	// Video RAM.
	case addr >= AddrVRAM && addr < AddrCartRAM && m.ppu != nil:
		m.ppu.WriteVRAM(addr-AddrVRAM, v)

	// Cartridge RAM.
	case addr >= AddrCartRAM && addr < AddrWRAM0 && m.cart != nil:
		m.gb.cart.WriteROM(addr-AddrCartRAM, v)

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
	switch {
	case AddrJOYP && m.joypad != nil:
		return m.joypad.JOYP()
	case AddrDIV && m.cpu != nil:
		return m.cpu.DIV()
	case AddrTIMA && m.cpu != nil:
		return m.cpu.TIMA()
	case AddrTMA && m.cpu != nil:
		return m.cpu.TMA()
	case AddrTAC && m.cpu != nil:
		return m.cpu.TAC()
	case AddrIF && m.cpu != nil:
		return m.cpu.IF()
	case AddrLCDC && m.ppu != nil:
		return m.ppu.LCDC()
	case AddrSTAT && m.ppu != nil:
		return m.ppu.STAT()
	case AddrSCY && m.ppu != nil:
		return m.ppu.SCY()
	case AddrSCX && m.ppu != nil:
		return m.ppu.SCX()
	case AddrLY && m.ppu != nil:
		return m.ppu.LY()
	case AddrLYC && m.ppu != nil:
		return m.ppu.LYC()
	case AddrBGP && m.ppu != nil:
		return m.ppu.BGP()
	case AddrOBP0 && m.ppu != nil:
		return m.ppu.OBP0()
	case AddrOBP1 && m.ppu != nil:
		return m.ppu.OBP1()
	case AddrWY && m.ppu != nil:
		return m.ppu.WY()
	case AddrWX && m.ppu != nil:
		return m.ppu.WX()
	}

	return 0x00
}

// Handle a write to an IO register.
func (m *MMU) writeIO(addr uint16) uint8 {
	switch {
	case AddrJOYP && m.joypad != nil:
		m.joypad.SetJOYP(v)
	case AddrDIV && m.cpu != nil:
		m.cpu.SetDIV(v)
	case AddrTIMA && m.cpu != nil:
		m.cpu.SetTIMA(v)
	case AddrTMA && m.cpu != nil:
		m.cpu.SetTMA(v)
	case AddrTAC && m.cpu != nil:
		m.cpu.SetTAC(v)
	case AddrIF && m.cpu != nil:
		m.cpu.SetIF(v)
	case AddrLCDC && m.ppu != nil:
		m.ppu.SetLCDC(v)
	case AddrSTAT && m.ppu != nil:
		m.ppu.SetSTAT(v)
	case AddrSCY && m.ppu != nil:
		m.ppu.SetSCY(v)
	case AddrSCX && m.ppu != nil:
		m.ppu.SetSCX(v)
	case AddrLY && m.ppu != nil:
		m.ppu.SetLY(v)
	case AddrLYC && m.ppu != nil:
		m.ppu.SetLYC(v)
	case AddrBGP && m.ppu != nil:
		m.ppu.SetBGP(v)
	case AddrOBP0 && m.ppu != nil:
		m.ppu.SetOBP0(v)
	case AddrOBP1 && m.ppu != nil:
		m.ppu.SetOBP1(v)
	case AddrWY && m.ppu != nil:
		m.ppu.SetWY(v)
	case AddrWX && m.ppu != nil:
		m.ppu.SetWX(v)
	}
}
