package main

import (
	"flag"
	"fmt"

	"github.com/ajm188/euler/math"
)

func main() {
	limit := flag.Int("limit", 28123, "")
	flag.Parse()

	var (
		abundants  []int
		abundanceM = map[int]bool{}
		sum        int
	)

	for i := 1; i <= *limit; i++ {
		abundanceM[i] = math.IsAbundant(i)
		if abundanceM[i] {
			abundants = append(abundants, i)
		}
	}

	for i := 1; i <= *limit; i++ {
		found := false // found=true if i can be written as the sum of two abundants.
		for j := 0; j < len(abundants) && abundants[j] < i; j++ {
			lhs := abundants[j]
			rhs := i - lhs
			if isAbundant, ok := abundanceM[rhs]; ok && isAbundant {
				found = true
				break
			}
		}

		if !found {
			sum += i
		}
	}

	fmt.Println(sum)
}
