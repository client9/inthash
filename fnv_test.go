package inthash

import "testing"

// testing unrolled versions match the slow-but-correct implimentations

func TestHashFNV32(t *testing.T) {
	const x = 10
	if hashint_fnv_naive(x) != FNV32(x) {
		t.Errorf("Fail")
	}
}

func TestHashFNV64(t *testing.T) {
	const x = 10
	if hashint_fnv64_naive(x) != FNV64(x) {
		t.Errorf("Fail")
	}
}
