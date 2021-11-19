package main

import (
	"flag"
	"fmt"
	"log"
	"math/big"

	"github.com/ajm188/euler/math"
)

type factorialChain struct {
	cur   *big.Int
	terms map[string]bool
}

func (fc *factorialChain) next() *big.Int {
	s := fc.cur.String()
	if _, ok := fc.terms[s]; ok {
		return math.CloneBigInt(fc.cur)
	}

	fc.terms[fc.cur.String()] = true
	fc.cur = math.SumDigitFactorials(fc.cur)
	return math.CloneBigInt(fc.cur)
}

func main() {
	target := flag.Int("target", 60, "")
	maxStart := flag.Int("max-start", 1_000_000, "")
	flag.Parse()

	// If I wanted to be super clever and effecient, I would keep a map of
	// <start num> => <num_terms>, and then if, while following a chain we hit
	// a number already in the map, just add the number of terms without
	// repeating the computation.
	// But, you know what? Patience is a virtue my friend.

	var count int
	for i := 1; i < *maxStart; i++ {
		chain := &factorialChain{
			cur:   big.NewInt(int64(i)),
			terms: map[string]bool{},
		}

		done := false
		for !done {
			cur := chain.cur
			next := chain.next()
			done = cur.Cmp(next) == 0

			if i == 169 {
				log.Printf("%s=>%s\n", cur, next)
			}
		}

		if len(chain.terms) == *target {
			count++
		}
	}

	fmt.Println(count)
}
