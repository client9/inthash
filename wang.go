package inthash

func Wang32(key uint32) uint32 {
	key = ^key + (key << 15) // key = (key << 15) - key - 1;
	key = key ^ (key >> 12)
	key = key + (key << 2)
	key = key ^ (key >> 4)
	key = key * 2057 // key = (key + (key << 3)) + (key << 11);
	key = key ^ (key >> 16)
	return key
}

func Wang64(key uint64) uint64 {
	key = (^key) + (key << 21) // key = (key << 21) - key - 1;
	key = key ^ (key >> 24)
	key = (key + (key << 3)) + (key << 8) // key * 265
	key = key ^ (key >> 14)
	key = (key + (key << 2)) + (key << 4) // key * 21
	key = key ^ (key >> 28)
	key = key + (key << 31)
	return key
}

// func hash32shiftmult(key uint32) uint32
func Wang32_shiftmult(key uint32) uint32 {
	const c2 = 0x27d4eb2d // a prime or an odd constant
	key = (key ^ 61) ^ (key >> 16)
	key = key + (key << 3)
	key = key ^ (key >> 4)
	key = key * c2
	key = key ^ (key >> 15)
	return key
}
