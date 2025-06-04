package main

import (
	"fmt"
	"math"

	"github.com/client9/inthash"

	"gonum.org/v1/gonum/mathext"
)

// ChiSquaredP computes the p-value from a chi-squared statistic and degrees of freedom.
func ChiSquaredP(chi2 float64, df int) float64 {
	if chi2 < 0 || df <= 0 {
		return math.NaN()
	}
	return mathext.GammaIncReg(0.5*float64(df), 0.5*float64(chi2))
}
func main() {

	const rounds = 1 << 32
	fmt.Printf("Rounds = %d\n", rounds)
	const hashBits = 12
	const mask = (1 << hashBits) - 1
	//randFunc := rand.Uint32
	hashFunc := inthash.Jenkins //inthash.hashint_fib32_reverse //hashint_jenkins //hashint_fnv_unroll //lashint2 //hashint_jenkins //murmur3_mix32 //

	fmt.Printf("Rounds: %d\n", rounds)
	fmt.Printf("Expected: %d\n", rounds/(1<<hashBits))
	buckets := make([]float64, 1<<hashBits, 1<<hashBits)

	for r := 0; r < rounds; r++ {
		//x1 := randFunc()
		h1 := hashFunc(uint32(r))
		b := h1 & mask
		buckets[b]++
	}

	var amin float64 = rounds
	var amax float64 = 0
	var expected = float64(rounds / (1 << hashBits))
	var sum float64
	for i := 0; i < 1<<hashBits; i++ {
		delta := float64(buckets[i] - expected)
		if delta < amin {
			amin = delta
		}
		if delta > amax {
			amax = delta
		}
		//fmt.Printf("%03d: %g %g %d\n", i, buckets[i], delta, int(math.Round(100*(float64(delta))/expected)))
		sum += (delta * delta) / expected
	}
	fmt.Printf("min,max error: %g %g\n", amin, amax)
	fmt.Printf("Chi Squared: %g\n", sum)
	fmt.Printf("P-Value 1: %g\n", ChiSquaredP(sum, (1<<hashBits)-1))
	fmt.Printf("P-Value 2: %g\n", mathext.GammaIncRegComp(0.5*float64((1<<hashBits)-1), 0.5*sum))
}
