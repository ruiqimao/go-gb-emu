package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/ruiqimao/go-gb-emu/cart"
	"github.com/ruiqimao/go-gb-emu/gb"
	"github.com/ruiqimao/go-gfx/gfx"
)

type Emulator struct {
	gb *gb.GameBoy
	dp *Display
}

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: %v <boot.bin> <rom.gb>\n", os.Args[0])
		os.Exit(1)
	}

	// Create and run the emulator.
	_, err := NewEmulator(os.Args[1], os.Args[2])
	if err != nil {
		log.Fatal(err)
	}

	// Run the graphics loop. This must be done on the main thread.
	gfx.Run()
}

func NewEmulator(bootPath string, cartPath string) (*Emulator, error) {
	e := &Emulator{}
	var err error

	// Create the gameboy.
	e.gb, err = gb.NewGameBoy()
	if err != nil {
		return nil, err
	}

	// Create a display for the gameboy.
	e.dp, err = NewDisplay()
	if err != nil {
		return nil, err
	}

	// Load the boot ROM.
	boot, err := ioutil.ReadFile(bootPath)
	if err != nil {
		return nil, err
	}
	err = e.gb.LoadBootRom(boot)
	if err != nil {
		return nil, err
	}

	// Load the cartridge.
	cartData, err := ioutil.ReadFile(cartPath)
	if err != nil {
		return nil, err
	}
	cart, err := cart.NewCartridge(cartData)
	if err != nil {
		return nil, err
	}
	e.gb.LoadCartridge(cart)

	// Run the main loop.
	go e.mainLoop()

	// Run the debug loop.
	go e.debugLoop()

	return e, nil
}

func (e *Emulator) mainLoop() {
	// Start the gameboy.
	e.gb.Resume()

	// Handle communication between the display and the gameboy.
	for {
		select {

		// Receive frame from gameboy.
		case frame := <-e.gb.F:
			e.dp.Draw(frame)

		// Receive input from display.
		case event := <-e.dp.I:
			e.gb.Input(event)

		}
	}
}
