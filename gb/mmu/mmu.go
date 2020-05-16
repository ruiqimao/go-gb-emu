package mmu

type MMU struct {
	cpu    CPU
	cpuBus CPUBus

	ppu    PPU
	ppuBus PPUBus

	// RAM.
	wram [0x2000]uint8
	hram [0xff]uint8
}

func NewMMU() *MMU {
	m := &MMU{}

	// Create the CPU bus.
	m.cpuBus = CPUBus{
		Read:  m.cpuRead,
		Write: m.cpuWrite,
	}

	// Create the PPU bus.
	m.ppuBus = PPUBus{
		Read:  m.ppuRead,
		Write: m.ppuWrite,
	}

	return m
}

// Attach a CPU.
func (m *MMU) AttachCPU(cpu CPU) {
	m.cpu = cpu
}

// Attach a PPU.
func (m *MMU) AttachPPU(ppu PPU) {
	m.ppu = ppu
}

// Get the CPU bus.
func (m *MMU) CPUBus() CPUBus {
	return m.cpuBus
}

// Get the PPU bus.
func (m *MMU) PPUBus() PPUBus {
	return m.ppuBus
}
