package inthash

import "math/bits"

const golden uint64 = 11400714819323198485

func Fib32(x uint32) uint32 {
	return uint32((uint64(x) * golden) >> 32)
}

// yes it's just input * constant
func Fib64(x uint64) uint64 {
	return x * golden
}

// Reverse output bits for testing performance of high bits in the testing
// frameworks
func Fib64_reverse(x uint64) uint64 {
	return bits.Reverse64(x * golden)
}

// Inlines two rounds into one
func Fib32_2iter(x uint32) uint32 {
	k := uint64(x)
	return uint32((((k * golden) >> 32) * golden) >> 32)
}

func Runes3(a, b, c rune) uint32 {
	i := uint64(a) << 42
	j := uint64(b) << 21
	k := uint64(c)
	return uint32(((i | j | k) * uint64(golden)) >> (53))

	// return uint32(((uint64(a)<<42 | uint64(b)<<21 | uint64(c)) * uint64(golden)) >> 53)
}
