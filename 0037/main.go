package main

import (
	"context"
	"flag"
	"fmt"
	"math/big"

	"github.com/ajm188/euler/math"
)

func main() {
	targetCount := flag.Int("target-count", 11, "")
	flag.Parse()

	ctx, cancel := context.WithCancel(context.Background())
	ch := math.Primes(ctx, 100)

	primes := map[int64]struct{}{}
	truncatablePrimes := make([]int, 0, *targetCount)
	for p := range ch {
		if len(truncatablePrimes) >= *targetCount {
			cancel()
			break
		}

		primes[int64(p)] = struct{}{}

		if p < 10 {
			continue
		}

		digits := fmt.Sprintf("%d", p)
		truncatable := true
		for i := 0; i < len(digits); i++ {
			ltr := digits[i:]
			rtl := digits[:i]

			if ltr == "" || ltr == digits || rtl == "" || rtl == digits {
				continue
			}

			var l, r big.Int
			l.SetString(ltr, 10)
			r.SetString(rtl, 10)

			li := l.Int64()
			ri := r.Int64()
			if _, ok := primes[li]; !ok {
				truncatable = false
				break
			}

			if _, ok := primes[ri]; !ok {
				truncatable = false
				break
			}
		}

		if truncatable {
			truncatablePrimes = append(truncatablePrimes, p)
		}
	}

	fmt.Println(math.SumInts(truncatablePrimes))
}
