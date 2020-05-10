package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ruiqimao/go-gb-emu/gb"
	"github.com/ruiqimao/go-gfx/gfx"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %v <rom>\n", os.Args[0])
		os.Exit(1)
	}

	// Create the gameboy.
	gameboy, err := gb.NewGameBoy()
	if err != nil {
		log.Fatal(err)
	}

	// Create a display for the gameboy.
	display, err := NewDisplay()
	if err != nil {
		log.Fatal(err)
	}

	// Run the main loop.
	go mainLoop(gameboy, display)

	// Run the graphics loop. This must be done on the main thread.
	gfx.Run()
}

func mainLoop(gameboy *gb.GameBoy, display *Display) {
	// TODO.
}
