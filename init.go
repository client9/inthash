package inthash

import (
	"math/rand"
)

var seed64 uint64
var seed32 uint32

func init() {
	seed64 = rand.Uint64()
	seed32 = rand.Uint32()
}
