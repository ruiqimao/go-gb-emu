package mmu

const (
	DMAClocks = 644
)

// Get the value of the DMA register.
func (m *MMU) DMA() uint8 {
	return m.dma
}

// Set the value of the DMA register.
func (m *MMU) SetDMA(v uint8) {
	m.dma = v

	// Start DMA.
	m.dmaClocks = DMAClocks
}

// Do a step of DMA.
func (m *MMU) stepDMA() {
	// Calculate how many clocks have elapsed.
	clocks := DMAClocks - m.dmaClocks

	// Copies are done after the 4th clock and on every 4 clocks.
	if clocks < 4 || clocks%4 != 0 {
		return
	}

	// Copy the byte.
	offset := clocks/4 - 1
	dst := AddrOAM + offset
	src := uint16(m.dma)*0x100 + offset
	m.write(dst, m.read(src))
}
