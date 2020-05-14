package utils

// Combine two bytes into a short.
func CombineBytes(hi uint8, lo uint8) uint16 {
	return (uint16(hi) << 8) | uint16(lo)
}

// Split a short into high and low bytes.
func SplitShort(s uint16) (uint8, uint8) {
	return uint8(s >> 8), uint8(s)
}

// Get the bit at a position.
func GetBit(v uint8, i int) bool {
	return (v>>i)&0x1 == 0x1
}

// Set the bit at a position.
func SetBit(v uint8, i int, on bool) uint8 {
	if on {
		return v | (0x1 << i)
	} else {
		return v & ^(0x1 << i)
	}
}

// Get the bit at a position for a short.
func GetBit16(v uint16, i int) bool {
	return (v>>i)&0x1 == 0x1
}

// Set the bit at a position for a short.
func SetBit16(v uint16, i int, on bool) uint16 {
	if on {
		return v | (0x1 << i)
	} else {
		return v & ^(0x1 << i)
	}
}

// Copy a bit from one byte to another.
func CopyBit(src uint8, dst uint8, i int) uint8 {
	return SetBit(src, i, GetBit(dst, i))
}
