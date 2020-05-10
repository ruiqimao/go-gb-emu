package gb

type GameBoy struct {
}

func NewGameBoy() (*GameBoy, error) {
	gb := &GameBoy{}
	return gb, nil
}
