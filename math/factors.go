package math

import (
	"math"
	"sort"
)

// Factors returns a sorted list of the factors of n. The factors of an integer
// n are all natural numbers i <= n such that n%i==0, including both 1 and n
// itself.
//
// The returned slice will never be nil and is guaranteed to have no duplicates.
func Factors(n int) (factors []int) {
	for i := 1; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			factors = append(factors, i)
			if div := n / i; div != i { // don't add duplicates
				factors = append(factors, div)
			}
		}
	}

	if factors == nil {
		factors = []int{}
	}
	sort.Ints(factors)
	return factors
}
