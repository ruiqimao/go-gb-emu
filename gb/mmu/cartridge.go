package mmu

type Cartridge interface {
	Read(uint16) uint8
	Write(uint8)
}
