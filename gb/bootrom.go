package gb

import (
	"fmt"
)

// Game Boy boot ROM.
type BootROM struct {
	rom     [0x100]uint8
	enabled bool
}

func NewBootROM(rom []uint8) (*BootROM, error) {
	if len(rom) != 0x100 {
		return nil, fmt.Errorf("Improper Boot ROM size: %v", len(rom))
	}

	b := &BootROM{
		enabled: true,
	}
	copy(b.rom[:], rom)
	return b, nil
}

// Read a byte from the boot ROM.
func (b *BootROM) Read(addr uint16) uint8 {
	return b.rom[addr]
}

// Get the BOOT register.
func (b *BootROM) BOOT() uint8 {
	if b.enabled {
		return 0x0
	} else {
		return 0x1
	}
}

// Set the BOOT register.
func (b *BootROM) SetBOOT(v uint8) {
	if v == 0x1 {
		b.enabled = false
	} else {
		b.enabled = true
	}
}
