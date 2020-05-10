package gb

type GameBoy struct {
	Cpu *Cpu
}

func NewGameBoy() (*GameBoy, error) {
	gb := &GameBoy{
		Cpu: NewCpu(),
	}
	return gb, nil
}
