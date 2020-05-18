package mmu

type Cartridge interface {
	ReadROM(uint16) uint8
	ReadRAM(uint16) uint8

	WriteROM(uint8)
	WriteRAM(uint8)
}
