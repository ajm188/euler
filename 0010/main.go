package main

import (
	"flag"
	"fmt"

	"github.com/ajm188/euler/math"
)

func main() {
	n := flag.Int("n", 10, "")
	flag.Parse()

	ch := math.PrimesUntil(*n, func(x, i int) bool {
		return x < *n
	})
	sum := 0
	for x := range ch {
		sum += x
	}

	fmt.Println(sum)
}
