package ppu

// Fetcher states.
const (
	FetcherTile  = 0
	FetcherData0 = 1
	FetcherData1 = 2
	FetcherIdle  = 3
)

// A Pixel contains color data and its source.
type Pixel struct {
	data uint8

	bg       bool
	palette  uint8
	priority bool
}

// The Fetcher retrieves data from VRAM and stores it in the FIFO.
// This implementation uses the Ultimate Game Boy Talk as a reference. Although the talk is
// slightly wrong with some regards to the PPU pipeline, it is close enough for an acceptable
// approximation.
// TODO: Sprites.
type Fetcher struct {
	ppu *PPU

	// Pixel FIFO.
	fifo []Pixel

	// Fetcher state.
	state uint8
	bgMap uint16

	tileX       uint8
	tileY       uint8
	tileDiscard uint8
	tileOffset  uint8

	tileN uint8
	data0 uint8
	data1 uint8
}

func NewPixel(data uint8, bg bool, palette uint8, priority bool) Pixel {
	return Pixel{data, bg, palette, priority}
}

func NewFetcher(ppu *PPU) *Fetcher {
	f := &Fetcher{
		ppu: ppu,
	}
	return f
}

// Reset the fetcher.
func (f *Fetcher) Reset(x uint8, y uint8, mapAddr uint16) {
	f.fifo = nil

	// Keep the upper 5 bits to get tiles from the map.
	f.tileX = x / 8
	f.tileY = y / 8

	// Use the lower 3 bits to determine how many pixels to discard.
	f.tileDiscard = x & 0x03

	// Use the y position modulo 8 to determine the vertical offset for the pixel.
	f.tileOffset = (y % 8) * 2

	// Set the map address.
	f.bgMap = mapAddr

	// Reset the state.
	f.state = FetcherTile
}

// Pop an element off the FIFO. Returns the color and whether the pop was successful.
func (f *Fetcher) Pop() (uint8, bool) {
	// FIFO must have at least 8 elements.
	if len(f.fifo) <= 8 {
		return 0, false
	}

	// Pop a pixel off.
	pixel := f.fifo[0]
	f.fifo = f.fifo[1:]

	// Discard the pixel if needed.
	if f.tileDiscard > 0 {
		f.tileDiscard--
		return 0, false
	}

	// Return the evaluated color.
	return f.ppu.resolve(pixel), true
}

// Do one step of the fetcher.
// The fetcher runs at half speed, so this should be run every other clock.
func (f *Fetcher) Step() {
	switch f.state {

	case FetcherTile:
		// Fetch the tile number.
		f.tileN = f.ppu.vram[f.bgMap+uint16(f.tileY)*32+uint16(f.tileX)]
		f.state = FetcherData0

	case FetcherData0:
		// Fetch the first byte of data.
		f.data0 = f.ppu.tileData(f.tileN, f.tileOffset, false)
		f.state = FetcherData1

	case FetcherData1:
		// Fetch the second byte of data.
		f.data1 = f.ppu.tileData(f.tileN, f.tileOffset+1, false)

		// Try to load the data into the FIFO.
		if f.load() {
			f.state = FetcherTile
		} else {
			f.state = FetcherIdle
		}

	case FetcherIdle:
		// Try to load the data into the FIFO.
		if f.load() {
			f.state = FetcherTile
		}

	}
}

// Try to load data into the FIFO. Returns if successful.
func (f *Fetcher) load() bool {
	// FIFO can have at most 8 elements to load.
	if len(f.fifo) > 8 {
		return false
	}

	// Load each pixel. Load from MSB to LSB, since MSB is to the left.
	for i := 7; i >= 0; i-- {
		lo := (f.data0 >> i) & 0x1
		hi := (f.data1 >> i) & 0x1

		data := lo | hi<<1
		// TODO: Handle sprites properly.
		bg := true
		palette := uint8(0)
		priority := false

		f.fifo = append(f.fifo, NewPixel(data, bg, palette, priority))
	}

	// Increment the tile X position.
	f.tileX = (f.tileX + 1) % 32

	return true
}
