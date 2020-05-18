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

	RequestInterrupt(int)
}

type CPUBus struct {
	mmu *MMU
}

// Handle read operations from the CPU.
func (b *CPUBus) Read(addr uint16) uint8 {
	// TODO: CPU bus gates.
	return b.mmu.read(addr)
}

// Handle write operations from the CPU.
func (b *CPUBus) Write(addr uint16, v uint8) {
	// TODO: CPU bus gates.
	b.mmu.write(addr, v)
}

// Request an interrupt from the CPU.
func (m *MMU) requestInterrupt(interrupt int) {
	if m.cpu != nil {
		m.cpu.RequestInterrupt(interrupt)
	}
}
