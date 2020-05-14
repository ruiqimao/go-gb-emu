package main

import (
	"bufio"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/ruiqimao/go-gb-emu/gb"
	"github.com/ruiqimao/go-gfx/gfx"
)

func (e *Emulator) debugLoop() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		scanner.Scan()
		input := scanner.Text()
		if len(input) == 0 {
			break
		}
		e.debugExec(strings.Fields(input))
	}

	// Kill the emulator when debug loop ends.
	gfx.Halt()
}

func (e *Emulator) debugExec(input []string) {
	cpu := e.gb.CPU()
	ppu := e.gb.PPU()
	mem := e.gb.Memory()

	var err error
	cmd := strings.ToLower(input[0])
	switch cmd {

	// Dump CPU.
	case "dump", "d":
		// Print registers.
		fmt.Printf("B  C   D  E   H  L   A  F\n")
		fmt.Printf("%02x %02x  %02x %02x  %02x %02x  %02x %02x\n",
			cpu.B(), cpu.C(), cpu.D(), cpu.E(), cpu.H(), cpu.L(), cpu.A(), cpu.F())
		fmt.Printf("\n")

		// Print flags.
		fmt.Printf("Z N H C\n")
		fmt.Printf("%d %d %d %d\n",
			boolToUint8(cpu.FlagZ()),
			boolToUint8(cpu.FlagN()),
			boolToUint8(cpu.FlagH()),
			boolToUint8(cpu.FlagC()))
		fmt.Printf("\n")

		// Print stack pointer and program counter.
		fmt.Printf("SP: %04x (%04x)\n", cpu.SP(), mem.Read16(cpu.SP()))
		fmt.Printf("PC: %04x (%s)\n", cpu.PC(), e.gb.InstructionName())

	// Read memory.
	case "print", "p":
		if len(input) < 2 {
			fmt.Printf("Usage: print <address>\n")
			break
		}
		addr, err := hexToUint16(input[1])
		if err != nil {
			break
		}

		// Read both a byte and a short at the address.
		fmt.Printf("%02x %04x\n", mem.Read(addr), mem.Read16(addr))

	// Dump the background.
	case "background", "bg":
		vram := ppu.VRAM()
		bgMap := ppu.BgTileMap()
		fmt.Printf("%04x\n", bgMap)
		for i := 0; i < 32; i++ {
			for j := 0; j < 32; j++ {
				id := vram[bgMap+uint16(i)*32+uint16(j)]
				fmt.Printf("%02x ", id)
			}
			fmt.Printf("\n")
		}

	// Dump the window.
	case "window", "wd":
		vram := ppu.VRAM()
		wdMap := ppu.WindowTileMap()
		fmt.Printf("%04x\n", wdMap)
		for i := 0; i < 32; i++ {
			for j := 0; j < 32; j++ {
				id := vram[wdMap+uint16(i)*32+uint16(j)]
				fmt.Printf("%02x ", id)
			}
			fmt.Printf("\n")
		}

	// Display a tile.
	case "tile", "t":
		if len(input) < 2 {
			fmt.Printf("Usage: tile <id>\n")
			break
		}
		id, err := hexToUint8(input[1])
		if err != nil {
			break
		}

		// Get the pixel data.
		for i := 0; i < 16; i += 2 {
			line0 := ppu.Tile(id, uint8(i))
			line1 := ppu.Tile(id, uint8(i+1))

			for j := 7; j >= 0; j-- {
				lo := (line0 >> j) & 0x1
				hi := (line1 >> j) & 0x1
				data := lo | hi<<1

				// Convert to a Pixel and resolve the color
				px := gb.NewPixel(data, gb.PixelSrcBG)
				color := ppu.Resolve(px)

				// Print the color as a block.
				char := ""
				switch color {
				case 0:
					char = "\u2588"
				case 1:
					char = "\u2593"
				case 2:
					char = "\u2592"
				case 3:
					char = "\u2591"
				}
				fmt.Printf("%s%s", char, char)
			}
			fmt.Printf("\n")
		}

	// Step forward.
	case "step", "s":
		steps := 1
		if len(input) == 2 {
			steps, err = strconv.Atoi(input[1])
			if err != nil {
				break
			}
		}

		cycles := 0
		for i := 0; i < steps; i++ {
			cycles += e.gb.Step()
		}
		fmt.Printf("%d cycles\n", cycles)

	// Run until PC reaches memory address.
	case "break", "b":
		if len(input) < 2 {
			fmt.Printf("Usage: break <address>\n")
			break
		}
		addr, err := hexToUint16(input[1])
		if err != nil {
			break
		}

		cycles := e.gb.Step()
		for {
			if cpu.PC() == addr {
				break
			}

			cycles += e.gb.Step()
		}
		fmt.Printf("%d cycles\n", cycles)

	// Run.
	case "run", "r":
		e.gb.Resume()

	// Halt.
	case "halt", "h":
		e.gb.Pause()

	default:
		fmt.Fprintf(os.Stderr, "Unknown command %s\n", cmd)

	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}
}

func boolToUint8(v bool) uint8 {
	if v {
		return 1
	}
	return 0
}

func hexToUint8(s string) (uint8, error) {
	bytes, err := hex.DecodeString(fmt.Sprintf("%02s", s))
	if err != nil {
		return 0, err
	}
	return bytes[0], nil
}

func hexToUint16(s string) (uint16, error) {
	bytes, err := hex.DecodeString(fmt.Sprintf("%04s", s))
	if err != nil {
		return 0, err
	}
	return binary.BigEndian.Uint16(bytes), nil
}
