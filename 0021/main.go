package main

import (
	"flag"
	"fmt"

	"github.com/ajm188/euler/math"
)

var (
	d = map[int]int{}
)

// Let d(n) be defined as the sum of proper divisors of n (numbers less than n which divide evenly into n).
// If d(a) = b and d(b) = a, where a ≠ b, then a and b are an amicable pair and each of a and b are called amicable numbers.
func amicable(a, b int) bool {
	if a == b {
		return false
	}

	d_a, ok := d[a]
	if !ok {
		d_a = math.SumInts(math.ProperDivisors(a))
		d[a] = d_a
	}

	if d_a != b {
		// skip extra computation of d(b)
		return false
	}

	d_b, ok := d[b]
	if !ok {
		d_b = math.SumInts(math.ProperDivisors(b))
		d[b] = d_b
	}

	return d_b == a
}

func main() {
	limit := flag.Int("limit", 10000, "")
	flag.Parse()

	sum := 0
	for a := 1; a < *limit; a++ {
		for b := 1; b < a; b++ {
			if amicable(a, b) {
				sum += a + b
			}
		}
	}

	fmt.Println(sum)
}
