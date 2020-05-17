package gb

import (
	"fmt"
	"log"

	"github.com/ruiqimao/go-gb-emu/cart"
	"github.com/ruiqimao/go-gb-emu/gb/cpu"
	"github.com/ruiqimao/go-gb-emu/gb/ppu"
)

const (
	BaseClock = 256     // Run at a base of 256Hz.
	CPUClock  = 4194304 // CPU clock is 4.194304MHz.
)

type GameBoy struct {
	cpu  *cpu.CPU
	ppu  *ppu.PPU
	mem  *Memory
	jp   *Joypad
	cart *cart.Cartridge

	clk *Clock

	// Boot ROM.
	boot [0x100]uint8

	// Input events.
	events chan Input

	// Latest rendered frame.
	F chan []byte

	// Debug.
	dbgRom []uint8
}

func NewGameBoy() (*GameBoy, error) {
	gb := &GameBoy{
		events: make(chan Input, 16), // Allow a buffer of input events.
		F:      make(chan []uint8, 1),
	}

	// Create the components.
	gb.cpu = cpu.NewCPU()
	gb.ppu = ppu.NewPPU()
	gb.mem = NewMemory(gb)
	gb.jp = NewJoypad(gb)
	gb.clk = NewClock(BaseClock)

	// Attach components together.
	gb.cpu.AttachMMU(gb.mem)
	gb.ppu.AttachMMU(gb.mem)

	go gb.Run()

	return gb, nil
}

// Run the Game Boy loop.
func (gb *GameBoy) Run() {
	// Number of extra clocks consumed in the last tick.
	clockDebt := 0

	for {
		select {

		case <-gb.clk.C:
			clockDebt = gb.RunClocks(CPUClock/BaseClock - clockDebt)

		case event := <-gb.events:
			gb.jp.Handle(event)

		case frame := <-gb.ppu.F:
			select {
			case gb.F <- frame:
			default:
			}

		}
	}
}

// Run a number of clocks. Returns how many extra clocks above the given limit were taken.
func (gb *GameBoy) RunClocks(limit int) int {
	for limit > 0 {

		// Process an instruction.
		clocks, err := gb.cpu.Step()
		if err != nil {
			log.Fatal(err)
		}
		limit -= clocks

		// Catch the PPU up to the CPU.
		for clocks > 0 {
			clocks -= gb.ppu.Step()
		}
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

// Load a cartridge.
func (gb *GameBoy) LoadCartridge(cartridge *cart.Cartridge) {
	gb.cart = cartridge
}

// Get the CPU.
func (gb *GameBoy) CPU() *cpu.CPU {
	return gb.cpu
}

// Get the PPU.
func (gb *GameBoy) PPU() *ppu.PPU {
	return gb.ppu
}

// Get the memory controller.
func (gb *GameBoy) Memory() *Memory {
	return gb.mem
}

// Get the clock.
func (gb *GameBoy) Clock() *Clock {
	return gb.clk
}

// Register input.
func (gb *GameBoy) Input(event Input) {
	gb.events <- event
}
