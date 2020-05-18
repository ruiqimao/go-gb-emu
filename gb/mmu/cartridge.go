package mmu

type Cartridge interface {
	ReadROM(uint16) uint8
	ReadRAM(uint16) uint8

	WriteROM(uint16, uint8)
	WriteRAM(uint16, uint8)
}
