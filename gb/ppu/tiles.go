package ppu

type TileMap bool

const (
	TileMap0 TileMap = false
	TileMap1         = true
)

type Tileset bool

const (
	Tileset0 Tileset = false
	Tileset1         = true
)

func (p *PPU) BGP() uint8 {
	return p.bgp
}

func (p *PPU) SetBGP(v uint8) {
	p.bgp = v
}

func (p *PPU) OBP0() uint8 {
	return p.obp0
}

func (p *PPU) SetOBP0(v uint8) {
	p.obp0 = v
}

func (p *PPU) OBP1() uint8 {
	return p.obp1
}

func (p *PPU) SetOBP1(v uint8) {
	p.obp1 = v
}
