package main

import (
	"flag"
	"fmt"
	"math/big"
)

func main() {
	n := flag.Int("n", 1000, "")
	flag.Parse()

	var sum big.Int
	for i := 1; i <= *n; i++ {
		x := big.NewInt(int64(i))
		x.Exp(x, x, nil)
		sum.Add(&sum, x)
	}

	fmt.Println(sum.String()[len(sum.String())-10:])
}
