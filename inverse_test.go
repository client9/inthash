package inthash

import (
	"testing"
)

func TestInverse32(t *testing.T) {
	cases := []uint32{1, 3, 5, 7, 9, 11, 13}
	for _, a := range cases {
		inv := inverse32(a)
		id := a * inv
		if id != 1 {
			t.Errorf("%d *%d != 1", a, inv)
		}
	}
}
func TestInverse64(t *testing.T) {
	cases := []uint64{1, 3, 5, 7, 9, 11, 13}
	for _, a := range cases {
		inv := inverse64(a)
		id := a * inv
		if id != 1 {
			t.Errorf("%d *%d != 1", a, inv)
		}
	}
}
