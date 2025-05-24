package inthash

import (
	"math/bits"
	// "unsafe"
)

func Murmur3(data uint32) uint32 {

	const (
		c1_32 uint32 = 0xcc9e2d51
		c2_32 uint32 = 0x1b873593
	)

	h1 := seed32

	//nblocks := len(data) / 4
	//nblocks := 1
	//var p uintptr
	//if len(data) > 0 {
	//	p = uintptr(unsafe.Pointer(&data[0]))
	//}
	//p1 := p + uintptr(4*nblocks)
	//for ; p < p1; p += 4 {
	//k1 := *(*uint32)(unsafe.Pointer(p))
	k1 := data
	k1 *= c1_32
	k1 = bits.RotateLeft32(k1, 15)
	k1 *= c2_32

	h1 ^= k1
	h1 = bits.RotateLeft32(h1, 13)
	h1 = h1*4 + h1 + 0xe6546b64
	//}

	/*
		tail := data[nblocks*4:]

		var k1 uint32
		switch len(tail) & 3 {
		case 3:
			k1 ^= uint32(tail[2]) << 16
			fallthrough
		case 2:
			k1 ^= uint32(tail[1]) << 8
			fallthrough
		case 1:
			k1 ^= uint32(tail[0])
			k1 *= c1_32
			k1 = bits.RotateLeft32(k1, 15)
			k1 *= c2_32
			h1 ^= k1
		}
	*/
	//h1 ^= uint32(len(data))
	h1 ^= 4

	h1 ^= h1 >> 16
	h1 *= 0x85ebca6b
	h1 ^= h1 >> 13
	h1 *= 0xc2b2ae35
	h1 ^= h1 >> 16

	return h1
}

func Murmur3_mix32(k uint32) uint32 {
	k ^= k >> 16
	k *= 0x85ebca6b
	k ^= k >> 13
	k *= 0xc2b2ae35
	k ^= k >> 16
	return k
}
