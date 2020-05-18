package mmu

type MMU struct {
	cpu    CPU
	ppu    PPU
	joypad Joypad

	cpuBus    *CPUBus
	ppuBus    *PPUBus
	joypadBus *JoypadBus

	bootrom BootROM

	cartridge Cartridge

	// RAM.
	wram [0x2000]uint8
	hram [0xff]uint8

	// DMA.
	dma       uint8
	dmaClocks uint16
}

func NewMMU() *MMU {
	m := &MMU{}

	// Create the buses.
	m.cpuBus = &CPUBus{m}
	m.ppuBus = &PPUBus{m}
	m.joypadBus = &JoypadBus{m}

	return m
}

// Do a step of the MMU. Consumes 1 clock.
func (m *MMU) Step() {
	if m.dmaClocks > 0 {
		m.stepDMA()
		m.dmaClocks--
	}
}

// Attach a CPU.
func (m *MMU) AttachCPU(cpu CPU) {
	m.cpu = cpu
}

// Attach a PPU.
func (m *MMU) AttachPPU(ppu PPU) {
	m.ppu = ppu
}

// Attach a joypad.
func (m *MMU) AttachJoypad(joypad Joypad) {
	m.joypad = joypad
}

// Attach a boot ROM.
func (m *MMU) AttachBootROM(bootrom BootROM) {
	m.bootrom = bootrom
}

// Attach a cartridge.
func (m *MMU) AttachCartridge(cartridge Cartridge) {
	m.cartridge = cartridge
}

// Get the CPU bus.
func (m *MMU) CPUBus() *CPUBus {
	return m.cpuBus
}

// Get the PPU bus.
func (m *MMU) PPUBus() *PPUBus {
	return m.ppuBus
}

// Get the joypad bus.
func (m *MMU) JoypadBus() *JoypadBus {
	return m.joypadBus
}
