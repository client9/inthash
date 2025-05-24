package inthash

func Murmur2_64(k uint64) uint64 {

	// golang complains
	//  m << 3
	// overflows if m is const
	// need to precompute this.
	var m uint64 = 0xc6a4a7935bd1e995
	const r = 47
	const keyLen uint64 = 8

	//h := seed64 ^ (keyLen * m)
	h := seed64 ^ (m << 3)
	k *= m
	k ^= k >> r
	k *= m

	h ^= k
	h *= m

	h ^= h >> r
	h *= m
	h ^= h >> r

	return h
}
