package main

import (
	"flag"
	"fmt"
	"sort"
)

func smallestMultiple(factors []int) int {
	sort.Ints(factors)

	n := factors[len(factors)-1]
	for {
		// Eventually we will hit the full-product of `factors` which is
		// guaranteed to be a multiple and terminate the loop.
		if isMultiple(n, factors) {
			return n
		}

		n++
	}
}

func isMultiple(n int, factors []int) bool {
	for _, i := range factors {
		if n%i != 0 {
			return false
		}
	}

	return true
}

func main() {
	n := flag.Int("n", 10, "")
	flag.Parse()

	factors := make([]int, *n)
	for i := 1; i <= *n; i++ {
		factors[i-1] = i
	}

	fmt.Println(smallestMultiple(factors))
}
