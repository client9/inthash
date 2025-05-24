package inthash

func Prospector1(x uint32) uint32 {
	x ^= x >> 16
	x *= 0x21f0aaad
	x ^= x >> 15
	x *= 0x735a2d97
	x ^= x >> 15
	return x
}

// best available?
func Prospector2(x uint32) uint32 {
	x ^= x >> 16
	x *= 0x7feb352d
	x ^= x >> 15
	x *= 0x846ca86b
	x ^= x >> 16
	return x
}
