package main

import (
	"fmt"
	"math/big"

	"github.com/ajm188/euler/math"
)

func main() {
	sum := 0
	for a := int64(0); a < 100; a++ {
		for b := int64(0); b < 100; b++ {
			x := big.NewInt(a)
			x.Exp(x, big.NewInt(b), nil)
			if s := math.SumDigits(x); s > sum {
				sum = s
			}
		}
	}

	fmt.Println(sum)
}
