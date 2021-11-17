package main

import (
	"flag"
	"fmt"

	"github.com/ajm188/euler/math"
)

func main() {
	n := flag.Int("n", 1000, "")
	flag.Parse()

	var (
		primes   []int
		primeMap = map[int]bool{}
	)

	for p := range math.PrimesUntil(*n, func(N, i int) bool {
		return N >= *n
	}) {
		primes = append(primes, p)
		primeMap[p] = true
	}

	longestSum, mostTerms := 0, 0
	for i := 0; i < len(primes); i++ {
		sum := 0
		terms := 0

		for j := i; j < len(primes); j++ {
			sum += primes[j]
			terms++
			if _, ok := primeMap[sum]; ok {
				if terms > mostTerms {
					longestSum = sum
					mostTerms = terms
				}
			}
		}
	}

	fmt.Printf("%d (%d terms)\n", longestSum, mostTerms)
}
