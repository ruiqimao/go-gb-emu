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
	cpu := e.gb.CPU()
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
		addrBytes, err := hex.DecodeString(fmt.Sprintf("%04s", input[1]))
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			break
		}
		addr := binary.BigEndian.Uint16(addrBytes)

		// Read both a byte and a short at the address.
		fmt.Printf("%02x %04x\n", mem.Read(addr), mem.Read16(addr))

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
}

func boolToUint8(v bool) uint8 {
	if v {
		return 1
	}
	return 0
}
