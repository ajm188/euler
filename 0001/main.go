package main

import (
	"flag"
	"fmt"
	"log"
)

func loop(limit int) (sum int) {
	for i := 3; i < limit; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}

	return sum
}

func main() {
	limit := flag.Int("limit", 100, "how many natural numbers to process")

	flag.Parse()

	if *limit < 1 {
		log.Fatalf("-limit must be positive, not %d", *limit)
	}

	fmt.Println(loop(*limit))
}
