package main

import (
	"flag"
	"fmt"

	"github.com/ajm188/euler/math"
)

func main() {
	iters := flag.Int("iters", 50, "")
	limit := flag.Int("limit", 10000, "")
	flag.Parse()

	count := 0
	for i := 1; i < *limit; i++ {
		if math.IsLychrel(i, *iters) {
			count++
		}
	}

	fmt.Println(count)
}
