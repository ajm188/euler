package main

import (
	"flag"
	"fmt"

	"github.com/ajm188/euler/math"
)

func main() {
	n := flag.Int("n", 2, "")
	flag.Parse()

	// In an NxN grid, there are 2N decision points, half of which (that is, N)
	// must be "down".
	paths := math.Choose(2*(*n), *n)
	fmt.Println(paths)
}
