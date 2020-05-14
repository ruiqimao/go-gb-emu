package cart

// Cartridge for Game Boy.
// TODO: Most cartridge functionality.
type Cartridge struct {
	data []uint8
}

func NewCartridge(data []uint8) (*Cartridge, error) {
	c := &Cartridge{
		data: data,
	}

	return c, nil
}

// Read a byte from the cartridge ROM.
func (c *Cartridge) ReadROM(addr uint16) uint8 {
	return c.data[addr]
}

// Write a byte to the cartridge ROM. This is used for memory banking.
func (c *Cartridge) WriteROM(addr uint16, v uint8) {
	// TODO.
}

// Read a byte from the cartridge RAM.
func (c *Cartridge) ReadRAM(addr uint16) uint8 {
	// TODO.
	return 0x00
}

// Write a byte to the cartridge RAM.
func (c *Cartridge) WriteRAM(addr uint16, v uint8) {
	// TODO.
}
