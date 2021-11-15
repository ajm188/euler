package main

import (
	"flag"
	"fmt"
	"math/big"
	"strconv"
)

func bigExp(base, pow int) *big.Int {
	// For whatever reason, doing big.NewFloat(math.Pow(base, pow)) results in
	// different output from .String().
	//
	// Doing all of the mathematics directly within the big.Int type seems to
	// work, however. Gross, but it works!
	//
	// See: https://play.golang.org/p/Pju1fFmc9-A.
	x := big.NewInt(int64(base))
	for i := 1; i < pow; i++ {
		x = x.Mul(x, big.NewInt(int64(base)))
	}

	return x
}

func main() {
	base := flag.Int("base", 2, "")
	pow := flag.Int("pow", 15, "")

	flag.Parse()

	x := bigExp(*base, *pow)
	digits := x.String()
	sum := 0
	for i := 0; i < len(digits); i++ {
		digit := digits[i : i+1]
		n, err := strconv.ParseInt(digit, 10, 64)
		if err != nil {
			panic(err)
		}

		sum += int(n)
	}

	fmt.Println(sum)
}
