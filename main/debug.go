package main

import (
	"bufio"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"os"
	"strconv"
	"strings"

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
	var err error
	cmd := strings.ToLower(input[0])
	switch cmd {

	// Dump CPU.
	case "dump", "d":
		ss := e.gb.Snapshot()

		// Print registers.
		fmt.Printf("B  C  D  E  H  L  A  F\n")
		fmt.Printf("%02x %02x %02x %02x %02x %02x %02x %02x\n",
			ss.B, ss.C, ss.D, ss.E, ss.H, ss.L, ss.A, ss.F)
		fmt.Printf("\n")
		fmt.Printf("BC   DE   HL   AF\n")
		fmt.Printf("%04x %04x %04x %04x\n", ss.BC, ss.DE, ss.HL, ss.AF)
		fmt.Printf("\n")

		// Print flags.
		fmt.Printf("Z N H C\n")
		fmt.Printf("%d %d %d %d\n",
			boolToUint8(ss.FlagZ), boolToUint8(ss.FlagN), boolToUint8(ss.FlagH), boolToUint8(ss.FlagC))
		fmt.Printf("\n")

		// Print stack pointer and program counter.
		fmt.Printf("SP: %04x (%04x)\n", ss.SP, binary.LittleEndian.Uint16(ss.Memory[ss.SP:]))
		fmt.Printf("PC: %04x (%02x) (%s)\n", ss.PC, ss.Memory[ss.PC], ss.InstructionName)

	// Read memory.
	case "read", "r":
		ss := e.gb.Snapshot()

		addrBytes, err := hex.DecodeString(fmt.Sprintf("%04s", input[1]))
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			break
		}
		addr := binary.BigEndian.Uint16(addrBytes)

		// Read both a byte and a short at the address.
		b := ss.Memory[addr]
		s := binary.LittleEndian.Uint16(ss.Memory[addr:])
		fmt.Printf("%02x %04x\n", b, s)

	// Step forward.
	case "step", "s":
		steps := 1
		if len(input) == 2 {
			steps, err = strconv.Atoi(input[1])
			if err != nil {
				fmt.Fprintf(os.Stderr, "%v\n", err)
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
		addrBytes, err := hex.DecodeString(fmt.Sprintf("%04s", input[1]))
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			break
		}
		addr := binary.BigEndian.Uint16(addrBytes)

		cycles := e.gb.Step()
		for {
			if e.gb.PC() == addr {
				break
			}

			cycles += e.gb.Step()
		}
		fmt.Printf("%d cycles\n", cycles)

	// Run.
	case "run":
		e.gb.Resume()

	// Pause.
	case "pause", "p":
		e.gb.Pause()

	default:
		fmt.Fprintf(os.Stderr, "Unknown command %s\n", cmd)

	}
}

func boolToUint8(v bool) uint8 {
	if v {
		return 1
	}
	return 0
}
