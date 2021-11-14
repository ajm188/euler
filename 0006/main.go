package main

import (
	"flag"
	"fmt"
)

func main() {
	n := flag.Int("n", 10, "")
	flag.Parse()

	sum := 0
	sumOfSquares := 0
	for i := 1; i <= *n; i++ {
		sum += i
		sumOfSquares += (i * i)
	}

	fmt.Println((sum * sum) - sumOfSquares)
}
