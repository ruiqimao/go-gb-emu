package mmu

// CPU interface.
type CPU interface {
	DIV() uint8
	TIMA() uint8
	TMA() uint8
	TAC() uint8
	IF() uint8
	IE() uint8

	SetDIV(uint8)
	SetTIMA(uint8)
	SetTMA(uint8)
	SetTAC(uint8)
	SetIF(uint8)
	SetIE(uint8)
}

type CPUBus struct {
	Read  func(addr uint16) uint8
	Write func(addr uint16, v uint8)
}

// Handle read operations from the CPU.
func (m *MMU) cpuRead(addr uint16) uint8 {
	return m.read(addr)
}

// Handle write operations from the CPU.
func (m *MMU) cpuWrite(addr uint16, v uint8) {
	m.write(addr, v)
}
