package main

import (
	"flag"
	"fmt"
	"math"

	"github.com/ajm188/euler/text"
)

func main() {
	digits := flag.Int("digits", 2, "")
	flag.Parse()

	start := int(math.Pow(10, float64(*digits-1))) // digits==2; produce [10^1, 10^2)
	stop := int(math.Pow(10, float64(*digits)))

	largest := 0
	for i := start; i < stop; i++ {
		for j := start; j < stop; j++ {
			product := i * j
			if text.IsPalindrome(fmt.Sprintf("%d", product)) {
				if product > largest {
					largest = product
				}
			}
		}
	}

	fmt.Println(largest)
}
