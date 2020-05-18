package mmu

// The functions in this file should be used only for debugging purposes.

import (
	"github.com/ruiqimao/go-gb-emu/utils"
)

func (m *MMU) Read(addr uint16) uint8 {
	return m.read(addr)
}

func (m *MMU) Write(addr uint16, v uint8) {
	m.write(addr, v)
}

func (m *MMU) Read16(addr uint16) uint16 {
	hi := m.read(addr + 1)
	lo := m.read(addr)
	return utils.CombineBytes(hi, lo)
}

func (m *MMU) Write16(addr uint16, v uint16) {
	hi, lo := utils.SplitShort(v)
	m.write(addr, lo)
	m.write(addr+1, hi)
}
