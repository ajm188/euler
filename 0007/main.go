package main

import (
	"flag"
	"fmt"

	"github.com/ajm188/euler/math"
)

func main() {
	n := flag.Int("n", 6, "")
	flag.Parse()

	ch := math.FirstNPrimes(*n)
	var last int
	for x := range ch {
		last = x
	}

	fmt.Println(last)
}
