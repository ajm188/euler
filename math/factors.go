package math

import (
	stdmath "math"
	"sort"
)

// Factors returns a sorted list of the factors of n. The factors of an integer
// n are all natural numbers i <= n such that n%i==0, including both 1 and n
// itself.
//
// The returned slice will never be nil and is guaranteed to have no duplicates.
func Factors(n int) []int {
	return factors(n, nil)
}

// ProperDivisors returns a sorted list of the proper divisors of n. The proper
// divisors of n are all natural numbers i < n such that n%i==0, including 1 but
// not n itself.
//
// The returned slice will never be nil and is guaranteed to have no duplicates.
func ProperDivisors(n int) (divisors []int) {
	if n == 1 {
		// Special case: int(math.Sqrt(1)) == 1, meaning our common `factors`
		// loop will return a slice of []int{1}, which we would have to check
		// and remove. Rather than wasting (very small amounts of) time
		// computing the square root, and allocating and appending to a slice,
		// only to then remove all elements and return the empty slice, let's
		// just return the empty slice now.
		return []int{}
	}

	return factors(n, func(n, i, div int) bool {
		return div == n // don't add `n` itself to the list of divisors
	})
}

func factors(n int, skipCond func(n, i, div int) bool) (factors []int) {
	for i := 1; i <= int(stdmath.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			factors = append(factors, i)
			if div := n / i; div != i /* don't add duplicates */ && (skipCond == nil || !skipCond(n, i, div)) {
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
