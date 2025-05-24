package inthash

import (
	"testing"
)

// test to make sure we aren't doing an off by one error
func TestShiftRight(t *testing.T) {
	// 11 bits
	size := 11
	mask := (1 << size) - 1

	// mask should be 11 one bits
	if mask != 0x7ff {
		t.Errorf("Mask error")
	}
	if uint64(0xFFE0000000000000)>>(64-size) != uint64(mask) {
		t.Errorf("Shift failed")
	}
}

// dumb test to make sure we can pack 3 runes into 1 uint64
func TestShift(t *testing.T) {
	const maxrune = uint64(0x10FFFF)
	rune2 := maxrune << 21
	rune3 := maxrune << 42
	if rune2 != uint64(0x0000021FFFE00000) {
		t.Errorf("Shift 2 Failed")
	}
	if rune3 != uint64(0x43FFFC0000000000) {
		t.Errorf("Shift 3 failed")
	}
	packed := maxrune | rune2 | rune3
	if packed != uint64(0x43FFFE1FFFF0FFFF) {
		t.Errorf("super rune failed")
	}
	if packed>>42 != maxrune {
		t.Errorf("Shift right failed")
	}
}
