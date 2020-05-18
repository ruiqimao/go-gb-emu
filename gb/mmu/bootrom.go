package mmu

// Boot ROM interface.
type BootROM interface {
	BOOT() uint8

	SetBOOT(uint8)

	Read(uint16) uint8
}
