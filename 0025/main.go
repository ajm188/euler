package main

import (
	"flag"
	"fmt"
	"math/big"

	"github.com/ajm188/euler/math/sequences/fibonacci"
)

func main() {
	digits := flag.Int("digits", 1000, "")
	flag.Parse()

	index := 1
	for range fibonacci.BigFibonacciUntil(10, func(n *big.Int, i int) bool { return len(n.String()) == *digits }) {
		index++
	}

	fmt.Println(index)
}
