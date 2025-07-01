package inthash

import (
	"fmt"
	"log"
	"testing"
)

func TestHashDistribution(t *testing.T) {
	const hashBits = 11
	t.Skip()
	/*
		letters := []rune{}
		for i := 0; i < 0x2FF; i++ {
			r := rune(i)
			if unicode.IsLetter(r) {
				letters = append(letters, r)
			}
		}
	*/
	//chars := letters
	chars := []rune(" abcdefghijklmnopqrstuvwxyz")
	//chars := []rune(" abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	log.Printf("Got %d chars", len(chars))
	buckets := make([]int, 1<<hashBits, 1<<hashBits)
	for _, i := range chars {
		for _, j := range chars {
			for _, k := range chars {
				buckets[Runes3(i, j, k)]++
			}
		}
	}

	bmax := 0
	for _, val := range buckets {
		if val > bmax {
			bmax = val
		}
	}
	bmax++
	counts := make([]int, bmax, bmax)
	for _, val := range buckets {
		counts[val]++
	}

	sum := 0
	csum := 0
	nonzero := 0
	for i, val := range counts {
		if val == 0 {
			continue
		}

		csum += i
		nonzero += 1

		sum += i * val
		fmt.Printf("%d --> %d\n", i, val)
	}
	fmt.Printf("TOTAL: %d\n", sum)
	fmt.Printf("AVER : %f\n", float64(csum)/float64(nonzero))
	fmt.Printf("Expected: %f", float64(sum)/2048.0)
}

type hashtest32 struct {
	name string
	fn   Hash32
}

type hashtest64 struct {
	name string
	fn   Hash64
}

var tests64 = []hashtest64{
	{
		"Identity64",
		func(x uint64) uint64 { return x },
	},
	{
		"Fib64",
		Fib64,
	},
	{
		"Wyhash64",
		Wyhash64,
	},
	{
		"Wang64",
		Wang64,
	},
	{
		"Wang64b",
		Wang64b,
	},
	{
		"FNV64",
		FNV64,
	},
	{
		"murmur2",
		Murmur2_64,
	},
}

func BenchmarkIntHash64(b *testing.B) {

	for _, tt := range tests64 {
		b.Run(tt.name, func(b *testing.B) {
			count := uint64(0)
			for b.Loop() {
				for i := 0; i < 100; i++ {
					count = tt.fn(count)
				}
			}
		})
	}
}

func BenchmarkIntHash32(b *testing.B) {

	tests := []hashtest32{
		{
			"Identity32",
			func(x uint32) uint32 { return x },
		},
		{
			"CA 32",
			CA32_single,
		},
		{
			"CA 32 10x",
			CA32_10x,
		},
		{
			"Murmur3",
			Murmur3,
		},
		{
			"Murmur3 mix32",
			Murmur3_mix32,
		},
		{
			"FNV32",
			FNV32,
		},
		{
			"Wang32",
			Wang32,
		},
		{
			"Jenkins",
			Jenkins,
		},
		{
			"ShiftMul",
			Wang32_shiftmult,
		},
		{
			"ShiftMul b",
			Wang32_shiftmult_b,
		},
		{
			"Prospector",
			Prospector2,
		},
		{
			"Fib32",
			Fib32,
		},
		{
			"Fib32 2 iterations",
			Fib32_2iter,
		},
	}

	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			count := uint32(1)
			for b.Loop() {
				for i := 0; i < 100; i++ {
					count = tt.fn(count)
				}
			}
		})
	}
}
func BenchmarkWang32(b *testing.B) {
	count := uint32(1)
	for b.Loop() {
		for i := 0; i < 100; i++ {
			count = Wang32(count)
		}
	}
}
func BenchmarkFib32Inline(b *testing.B) {
	count := uint32(1)
	for b.Loop() {
		for i := 0; i < 100; i++ {
			count = uint32((uint64(count) * uint64(golden)) >> 32)
		}
	}
}
func BenchmarkHashRune3(b *testing.B) {
	var val uint32
	for b.Loop() {
		for i := 0; i < 100; i++ {
			val ^= Runes3('a', 'b', 'c')
		}
	}
}
