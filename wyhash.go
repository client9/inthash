package inthash

import (
	"math/bits"
)

func Wyhash64(x uint64) uint64 {
	const m5 = 0x1d8e4e27c47d124f

	var seed = seed64
	//seed ^= hashkey[0]
	//seed = 0
	a := x
	b := x<<32 | x>>32
	return mix(m5, mix(a, b^seed))
	//return mix(m5^8, mix(a^hashkey[1], b^seed))
}

/*
var hashkey = [2]uint64{
	0, 0,
	// 0x13792739ab7387,
	// 0xDEADBEEF934343,
}
*/

func mix(a, b uint64) uint64 {
	hi, lo := bits.Mul64(a, b)
	return hi ^ lo
}
