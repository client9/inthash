package inthash

import (
	"hash/maphash"
	"unsafe"
)

type stringStruct struct {
	str unsafe.Pointer
	len int
}

//go:noescape
//go:linkname memhash runtime.memhash
func memhash(p unsafe.Pointer, h, s uintptr) uintptr

// MemHash is the hash function used by go map, it utilizes available hardware instructions(behaves
// as aeshash if aes instruction is available).
// NOTE: The hash seed changes for every process. So, this cannot be used as a persistent hash.
func MemHash(data []byte) uint64 {
	ss := (*stringStruct)(unsafe.Pointer(&data))
	return uint64(memhash(ss.str, 0, uintptr(ss.len)))
}

func uint32ToBytesUnsafe(n uint32) []byte {
	slice := (*[4]byte)(unsafe.Pointer(&n))[:]
	return slice
}
func uint64ToBytesUnsafe(n uint64) []byte {
	slice := (*[8]byte)(unsafe.Pointer(&n))[:]
	return slice
}
func Memhash32(x uint32) uint32 {
	return uint32(MemHash(uint32ToBytesUnsafe(x)))
}
func Memhash64(n uint64) uint64 {
	return uint64(memhash(unsafe.Pointer(&n), 0, uintptr(64)))
	// return MemHash(uint64ToBytesUnsafe(x))
}

func Maphash32(x uint32) uint32 {
	hm.Reset()
	buf := []byte{
		byte((x >> 24) & 0xff),
		byte((x >> 16) & 0xff),
		byte((x >> 8) & 0xff),
		byte(x & 0xff),
	}
	hm.Write(buf)
	return uint32(hm.Sum64() >> 32)
}

var hm = new(maphash.Hash)

func Maphash64(x uint64) uint64 {
	hm.Reset()
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
	hm.Write(buf)
	return hm.Sum64()
}
