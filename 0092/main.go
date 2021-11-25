package main

import (
	"flag"
	"fmt"

	"github.com/ajm188/euler/math"
	"github.com/ajm188/euler/math/functools"
)

func main() {
	limit := flag.Int("limit", 10_000_000, "")
	target := flag.Int("target", 89, "")
	flag.Parse()

	var count int
	chains := make(map[string]int, *limit)
	chains["1"] = 1
	chains["89"] = 89

	squareFn := func(x int) int { return x * x }
	sumFn := func(x, y int) int { return x + y }

	for i := 1; i < *limit; i++ {
		terms := []string{}
		j := i
		for j != 1 && j != 89 {
			s := math.OrderedDigitsString(j)
			terms = append(terms, s)

			if end, ok := chains[s]; ok {
				j = end
			} else {
				j = functools.ReduceInts(
					functools.MapInts(math.Digits(j), squareFn),
					0, sumFn,
				)
			}
		}

		for _, term := range terms {
			if _, ok := chains[term]; !ok {
				chains[term] = j
			}
		}

		if j == *target {
			count++
		}
	}

	fmt.Println(count)
}
