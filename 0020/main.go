package main

import (
	"flag"
	"fmt"
	"math/big"
	"strconv"
)

func factorial(n int) *big.Int {
	f := big.NewInt(int64(n))
	for i := n - 1; i > 1; i-- {
		f = f.Mul(f, big.NewInt(int64(i)))
	}

	return f
}

func main() {
	n := flag.Int("n", 10, "")
	flag.Parse()

	nfact := factorial(*n)
	digits := nfact.String()
	sum := 0
	for _, d := range digits {
		x, err := strconv.ParseInt(string(d), 10, 64)
		if err != nil {
			panic(err)
		}

		sum += int(x)
	}

	fmt.Println(sum)
}
