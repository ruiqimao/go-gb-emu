package utils

// Combine two bytes into a short.
func CombineBytes(a uint8, b uint8) uint16 {
	return (uint16(a) << 8) | uint16(b)
}

// Split a short into two bytes.
func SplitShort(s uint16) (uint8, uint8) {
	return uint8(s >> 8), uint8(s)
}

// Get the bit at a position.
func GetBit(v uint8, i int) uint8 {
	return (v >> i) & 0x1
}

// Enable the bit at a position.
func EnableBit(v uint8, i int) uint8 {
	return v | (0x1 << i)
}

// Disable the bit at a position.
func DisableBit(v uint8, i int) uint8 {
	return v & ^(0x1 << i)
}
