package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand/v2"

	"github.com/client9/inthash"

)

func ExportJSModule(m [][]float64) string {
	raw, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}
	return string(raw)
}

/* Does not work */
/*
func getRandFunc[V uint32 | uint64](intType V) func() V {
	switch any(intType).(type) {
	case uint32:
		return rand.Uint32
	case uint64:
		return rand.Uint64
	}
	return nil
}
*/
/*
func getHashBits[V uint32 | uint64](intType V) int {
	switch any(intType).(type) {
	case uint32:
		return 32
	case uint64:
		return 64
	}
	return 0
}
*/
func avalanche32(hashFunc func(uint32) uint32) [][]float64 {
	return avalanche(32, hashFunc, rand.Uint32)
}
func avalanche64(hashFunc func(uint64) uint64) [][]float64 {
	return avalanche(64, hashFunc, rand.Uint64)
}
func avalanche[V uint32 | uint64](hashBits int, hashFunc func(V) V, randFunc func() V) [][]float64 {
	const rounds = 100000

	buckets := make([][]float64, hashBits, hashBits)
	for i := 0; i < hashBits; i++ {
		buckets[i] = make([]float64, hashBits, hashBits)
	}

	for r := 0; r < rounds; r++ {
		x1 := randFunc()
		h1 := hashFunc(x1)
		for i := 0; i < hashBits; i++ {
			x2 := x1 ^ (1 << i)
			h2 := hashFunc(x2)

			diff := h1 ^ h2
			for j := 0; j < hashBits; j++ {
				if diff&(1<<j) != 0 {
					buckets[i][j]++
				}
			}
		}
	}

	// normalize

	amin := 1.0
	amax := 0.0
	demon := float64(rounds)
	for i := 0; i < hashBits; i++ {
		for j := 0; j < hashBits; j++ {
			val := buckets[i][j] / demon
			if val < amin {
				amin = val
			}
			if val > amax {
				amax = val
			}

			buckets[i][j] = val
		}
	}
	log.Printf("Amin = %f, amax = %f", amin, amax)
	if false {
		amin = 0.49
		amax = 0.51
		for i := 0; i < hashBits; i++ {
			for j := 0; j < hashBits; j++ {
				val := (buckets[i][j] - amin) / (amax - amin)
				buckets[i][j] = val
			}
		}
	}

	return buckets
}

func main() {
	buckets := avalanche32(inthash.CA32_single)
	//buckets := avalanche32(Murmur3)
	//buckets := avalanche32(murmur3_mix32)
	//buckets := avalanche32(hashint_fib32)
	//buckets := avalanche64(murmur2)
	//buckets := avalanche64(hashint_wang64)
	//buckets := avalanche32(hashint_fnv_unroll)
	out := ExportJSModule(buckets)
	fmt.Println(out)

	//out := ListDensityPlot(buckets)
	//	fn := hash_fib64
	//fn := hashint_fib64
	//fn := hashint_fnv64_unroll
	//fn := hashint_wang64
	//fn := hash_memhash64
	//fn := hash_wyhash64

}
