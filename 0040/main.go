package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/ajm188/euler/math"
)

func main() {
	// prepend a decimal because it's a convenient and appropriate placeholder
	// and it makes the 1-based indexing in the problem work out of the box.
	d := fmt.Sprintf(".%s", math.Champernowne(1000000))
	product := 1

	for i := 1; i <= 1000000; i *= 10 {
		x, err := strconv.ParseInt(d[i:i+1], 10, 64)
		if err != nil {
			log.Fatalf("cannot parse d[%d] (%s): %s", i, d[i:i+1], err)
		}

		product *= int(x)
	}

	fmt.Println(product)
}
