package mmu

type MMU struct {
	cpu    CPU
	cpuBus CPUBus

	ppu    PPU
	ppuBus PPUBus

	joypad    Joypad
	joypadBus JoypadBus

	bootrom BootRom

	cartridge Cartridge

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
		RequestInterrupt: m.requestInterrupt,
	}

	// Create the joypad bus.
	m.joypadBus = JoypadBus{
		RequestInterrupt: m.requestInterrupt,
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
func (m *MMU) CPUBus() CPUBus {
	return m.cpuBus
}

// Get the PPU bus.
func (m *MMU) PPUBus() PPUBus {
	return m.ppuBus
}

// Get the joypad bus.
func (m *MMU) JoypadBus() JoypadBus {
	return m.joypadBus
}
