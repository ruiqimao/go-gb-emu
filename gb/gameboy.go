package gb

import (
	"fmt"
	"log"
)

const (
	BaseClock = 256     // Run at a base of 256Hz.
	CpuClock  = 4194304 // CPU clock is 4.194304MHz.
)

type GameBoy struct {
	cpu *Cpu
	ppu *Ppu
	mem *Memory

	clk *Clock

	// Boot ROM.
	boot [0x100]uint8

	// Debug.
	dbgRom []uint8
}

func NewGameBoy() (*GameBoy, error) {
	gb := &GameBoy{}

	// Create the components.
	gb.cpu = NewCpu(gb)
	gb.ppu = NewPpu(gb)
	gb.mem = NewMemory(gb)
	gb.clk = NewClock(BaseClock)

	// Initialize the instruction set.
	gb.cpu.CreateInstructionSet()

	go gb.Run()

	return gb, nil
}

// Run the Game Boy loop.
func (gb *GameBoy) Run() {
	// Number of extra cycles done in the last tick.
	cycleDebt := 0

	for {
		select {

		case <-gb.clk.C:
			cycleDebt = gb.RunCycles(CpuClock/BaseClock - cycleDebt)

		}
	}
}

// Run a number of cycles. Returns how many extra cycles above the given limit were taken.
func (gb *GameBoy) RunCycles(limit int) int {
	for limit > 0 {

		// Process an instruction.
		cycles, err := gb.cpu.Step()
		if err != nil {
			log.Fatal(err)
		}

		// Catch the PPU up to the CPU.
		gb.ppu.Update(cycles)

		limit -= cycles
	}
	return -limit
}

// Load the Boot ROM.
func (gb *GameBoy) LoadBootRom(rom []byte) error {
	if len(rom) != 0x100 {
		return fmt.Errorf("Improper Boot ROM size: %v", len(rom))
	}
	copy(gb.boot[:], rom)
	return nil
}
