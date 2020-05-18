package gb

import (
	"log"

	"github.com/ruiqimao/go-gb-emu/cart"
	"github.com/ruiqimao/go-gb-emu/gb/cpu"
	"github.com/ruiqimao/go-gb-emu/gb/joypad"
	"github.com/ruiqimao/go-gb-emu/gb/mmu"
	"github.com/ruiqimao/go-gb-emu/gb/ppu"
)

const (
	BaseClock = 256     // Run at a base of 256Hz.
	CPUClock  = 4194304 // CPU clock is 4.194304MHz.
)

type GameBoy struct {
	mmu  *mmu.MMU
	cpu  *cpu.CPU
	ppu  *ppu.PPU
	jp   *joypad.Joypad
	cart *cart.Cartridge

	clk *Clock

	// Input events.
	events chan joypad.Input

	// Latest rendered frame.
	F chan []byte
}

func NewGameBoy() (*GameBoy, error) {
	gb := &GameBoy{
		events: make(chan joypad.Input, 16), // Allow a buffer of input events.
		F:      make(chan []uint8, 1),
	}

	// Create the components.
	gb.mmu = mmu.NewMMU()
	gb.cpu = cpu.NewCPU()
	gb.ppu = ppu.NewPPU()
	gb.jp = joypad.NewJoypad()
	gb.clk = NewClock(BaseClock)

	// Attach components together.
	gb.cpu.AttachMMU(gb.mmu.CPUBus())
	gb.ppu.AttachMMU(gb.mmu.PPUBus())
	gb.jp.AttachMMU(gb.mmu.JoypadBus())

	gb.mmu.AttachCPU(gb.cpu)
	gb.mmu.AttachPPU(gb.ppu)
	gb.mmu.AttachJoypad(gb.jp)

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
	bootrom, err := NewBootROM(rom)
	if err != nil {
		return err
	}
	gb.mmu.AttachBootROM(bootrom)
	return nil
}

// Load a cartridge.
func (gb *GameBoy) LoadCartridge(cartridge *cart.Cartridge) {
	gb.mmu.AttachCartridge(cartridge)
}

// Register input.
func (gb *GameBoy) Input(event joypad.Input) {
	gb.events <- event
}
