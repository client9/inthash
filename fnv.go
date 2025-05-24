package inthash

import (
	"hash/fnv"
)

func FNV32(x uint32) uint32 {
	const prime32 = 16777619
	h := uint32(2166136261)
	h ^= (x >> 24) & 0xff
	h *= prime32
	h ^= (x >> 16) & 0xff
	h *= prime32
	h ^= (x >> 8) & 0xff
	h *= prime32
	h ^= (x) & 0xff
	h *= prime32
	return h
}

func FNV64(x uint64) uint64 {
	const fnv_prime64 uint64 = 1099511628211

	h := uint64(14695981039346656037)
	h ^= (x >> 56) & 0xff
	h *= fnv_prime64
	h ^= (x >> 48) & 0xff
	h *= fnv_prime64
	h ^= (x >> 40) & 0xff
	h *= fnv_prime64
	h ^= (x >> 32) & 0xff
	h *= fnv_prime64
	h ^= (x >> 24) & 0xff
	h *= fnv_prime64
	h ^= (x >> 16) & 0xff
	h *= fnv_prime64
	h ^= (x >> 8) & 0xff
	h *= fnv_prime64
	h ^= (x) & 0xff
	h *= fnv_prime64
	return h
}

//
// SLOWER
//
//

func hashint_fnv_naive(x uint32) uint32 {
	h := fnv.New32a()
	h.Write([]byte{
		byte((x >> 24) & 0xff),
		byte((x >> 16) & 0xff),
		byte((x >> 8) & 0xff),
		byte(x & 0xff),
	})
	return h.Sum32()
}

func hashint_fnv64_naive(x uint64) uint64 {
	buf := []byte{
		byte((x >> 56) & 0xff),
		byte((x >> 48) & 0xff),
		byte((x >> 40) & 0xff),
		byte((x >> 32) & 0xff),
		byte((x >> 24) & 0xff),
		byte((x >> 16) & 0xff),
		byte((x >> 8) & 0xff),
		byte(x & 0xff),
	}
	h := fnv.New64a()
	h.Write(buf)
	return h.Sum64()
}

func hashint_fnv64_naive2(x uint64) uint64 {
	buf := []byte{
		byte(x >> 56),
		byte(x >> 48),
		byte(x >> 40),
		byte(x >> 32),
		byte(x >> 24),
		byte(x >> 16),
		byte(x >> 8),
		byte(x),
	}
	h := fnv.New64a()
	h.Write(buf)
	return h.Sum64()
}
