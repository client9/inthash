package inthash

import "math/bits"

/*
k ^= k >> 16
k *= 0x85ebca6b
k ^= k >> 13
k *= 0xc2b2ae35
k ^= k >> 16
*/
func rule30(left, center, right uint32) uint32 {
	return left ^ (center | right)
}

func CA32_single(k uint32) uint32 {
	return bits.RotateLeft32(k, 31) ^ (k | bits.RotateLeft32(k, 1))
}
func CA32_10x(k uint32) uint32 {
	k = bits.RotateLeft32(k, 31) ^ (k | bits.RotateLeft32(k, 1))
	k = bits.RotateLeft32(k, 31) ^ (k | bits.RotateLeft32(k, 1))
	k = bits.RotateLeft32(k, 31) ^ (k | bits.RotateLeft32(k, 1))
	k = bits.RotateLeft32(k, 31) ^ (k | bits.RotateLeft32(k, 1))
	k = bits.RotateLeft32(k, 31) ^ (k | bits.RotateLeft32(k, 1))
	k = bits.RotateLeft32(k, 31) ^ (k | bits.RotateLeft32(k, 1))
	k = bits.RotateLeft32(k, 31) ^ (k | bits.RotateLeft32(k, 1))
	k = bits.RotateLeft32(k, 31) ^ (k | bits.RotateLeft32(k, 1))
	k = bits.RotateLeft32(k, 31) ^ (k | bits.RotateLeft32(k, 1))
	return bits.RotateLeft32(k, 31) ^ (k | bits.RotateLeft32(k, 1))
}
func CA32_single2(k uint32) uint32 {
	return k ^ (bits.RotateLeft32(k, 1) | bits.RotateLeft32(k, 2))
}

const delta = 7

func hash_ca32(k uint32) uint32 {
	//k = bits.RotateLeft32(k, delta) ^ (k | bits.RotateLeft32(k, 32-delta))
	k ^= k >> 16
	k *= 0x85ebca6b
	k ^= k >> 13
	//k = bits.RotateLeft32(k, delta) ^ (k | bits.RotateLeft32(k, 32-delta))
	k *= 0xc2b2ae35
	k ^= k >> 16

	//k = bits.RotateLeft32(k, delta) ^ (k | bits.RotateLeft32(k, 32-delta))
	//k ^= k >> 16
	//k *= 0xc2b2ae35
	/*
		k = bits.RotateLeft32(k, 1) ^ (k | bits.RotateLeft32(k, 31))
		k = bits.RotateLeft32(k, 1) ^ (k | bits.RotateLeft32(k, 31))
		k = bits.RotateLeft32(k, 1) ^ (k | bits.RotateLeft32(k, 31))
		k = bits.RotateLeft32(k, 1) ^ (k | bits.RotateLeft32(k, 31))
		k = bits.RotateLeft32(k, 1) ^ (k | bits.RotateLeft32(k, 31))
		k = bits.RotateLeft32(k, 1) ^ (k | bits.RotateLeft32(k, 31))

		k ^= k >> 15
		k *= 0xc2b2ae35
	*/
	return k
}
