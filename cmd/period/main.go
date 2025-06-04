package main

import (
	"fmt"

	"github.com/client9/inthash"
)

//var seed32 uint32 = 0

func main() {

	const rounds = 1 << 32
	fmt.Printf("Rounds = %d\n", rounds)

	hashFunc := inthash.CA32_single2


	fmt.Printf("Rounds: %d\n", rounds)

	kinit := uint32(0x1)
	k := kinit
	fmt.Printf("Seed: %d\n", k)
	for r := 0; r < rounds; r++ {
		k = hashFunc(k)
		if k == kinit {
			fmt.Printf("Cycle found at round %d\n", r)
			break
		}
	}

	for r := 0; r < rounds; r++ {
		k = hashFunc(uint32(r))
		if k == uint32(r) {
			fmt.Printf("Fixed point: %d\n", k)
		}
	}
	/*
		knext := hashFunc(k)
		if k == kinit || knext == kinit {
			fmt.Printf("Complete 1<<32 period\n")
			return
		}

		fmt.Printf("Didn't end with initial value, looking for cycles\n")
		buckets := map[uint32]struct{}{}
		buckets[kinit] = struct{}{}
		k = kinit
		var count uint64 = 1
		for {
			k = hashFunc(k)
			if _, ok := buckets[k]; ok {
				fmt.Printf("Found cycle of length %d with seed %d", count, kinit)
				return
			}
			buckets[k] = struct{}{}
			count++
			if count == 0xFFFFFFFF {
				fmt.Printf("Failed\n")
				return
			}
		}
	*/
}
