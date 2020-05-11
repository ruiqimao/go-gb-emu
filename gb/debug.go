package gb

// A snapshot of the Game Boy state.
type Snapshot struct {
	// Registers.
	B  uint8
	C  uint8
	D  uint8
	E  uint8
	H  uint8
	L  uint8
	A  uint8
	F  uint8
	BC uint16
	DE uint16
	HL uint16
	AF uint16

	// Flags.
	FlagZ bool
	FlagN bool
	FlagH bool
	FlagC bool

	// Stack pointer.
	SP uint16

	// Program counter.
	PC uint16
	InstructionName string

	// Memory.
	BootROM [0x100]uint8
	VRAM    [0x4000]uint8
	WRAM    [0x2000]uint8
	OAM     [0x100]uint8
	HRAM    [0xff]uint8
	IE      uint8
	Memory  [0x10000]uint8 // Full map.

	// Misc.
	Halt    bool
	HaltBug bool
	IME     bool
}

// Run the Game Boy clock.
func (gb *GameBoy) Resume() {
	gb.clk.Resume()
}

// Pause the Game Boy clock.
func (gb *GameBoy) Pause() {
	gb.clk.Pause()
}

// Step forward by one instruction. Returns how many cycles were taken.
func (gb *GameBoy) Step() int {
	return gb.RunCycles(1) + 1
}

// Take a snapshot of the Game Boy state.
func (gb *GameBoy) Snapshot() Snapshot {
	ss := Snapshot{}

	// Registers.
	ss.B = gb.cpu.B()
	ss.C = gb.cpu.C()
	ss.D = gb.cpu.D()
	ss.E = gb.cpu.E()
	ss.H = gb.cpu.H()
	ss.L = gb.cpu.L()
	ss.A = gb.cpu.A()
	ss.F = gb.cpu.F()
	ss.BC = gb.cpu.BC()
	ss.DE = gb.cpu.DE()
	ss.HL = gb.cpu.HL()
	ss.AF = gb.cpu.AF()

	// Stack pointer.
	ss.SP = gb.cpu.SP()

	// Program counter.
	ss.PC = gb.cpu.PC()
	inst := uint16(gb.mem.Read(ss.PC))
	if inst == 0xcb {
		inst = 0x100 + uint16(gb.mem.Read(ss.PC+1))
	}
	ss.InstructionName = InstructionNames[inst]

	// Flags.
	ss.FlagZ = gb.cpu.FlagZ()
	ss.FlagN = gb.cpu.FlagN()
	ss.FlagH = gb.cpu.FlagH()
	ss.FlagC = gb.cpu.FlagC()

	// Memory.
	copy(ss.BootROM[:], gb.boot[:])
	copy(ss.VRAM[:], gb.mem.vram[:])
	copy(ss.WRAM[:], gb.mem.wram[:])
	copy(ss.OAM[:], gb.mem.oam[:])
	copy(ss.HRAM[:], gb.mem.hram[:])
	ss.IE = gb.mem.ie
	for i := 0; i < 0x10000; i++ {
		ss.Memory[i] = gb.mem.Read(uint16(i))
	}

	// Misc.
	ss.Halt = gb.cpu.Halted()
	ss.HaltBug = gb.cpu.haltBug
	ss.IME = gb.cpu.IME()

	return ss
}
