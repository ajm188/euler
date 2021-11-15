package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/ajm188/euler/math"
)

func triangleNumbers(ctx context.Context, n int) chan int {
	ch := make(chan int, n)
	go func() {
		defer close(ch)

		n := 0
		for i := 1; ; i++ {
			select {
			case <-ctx.Done():
				return
			default:
			}

			n += i
			ch <- n
		}
	}()

	return ch
}

type FactoredInteger interface {
	Number() int
	Factors() []int
}

type factoredInt struct {
	n       int
	factors []int
}

func (fi *factoredInt) Number() int {
	return fi.n
}

func (fi *factoredInt) Factors() []int {
	if fi.factors != nil {
		return fi.factors
	}

	fi.factors = math.Factors(fi.n)
	return fi.factors
}

func factorize(ch chan int) chan FactoredInteger {
	outch := make(chan FactoredInteger, cap(ch))
	go func() {
		defer close(outch)

		for n := range ch {
			fi := &factoredInt{n: n}
			fi.Factors()
			outch <- fi
		}
	}()

	return outch
}

func main() {
	numFactors := flag.Int("num-factors", 5, "")
	flag.Parse()

	ctx, cancel := context.WithCancel(context.Background())
	trich := triangleNumbers(ctx, 20)

	withFactorsCh := factorize(trich)
	for tri := range withFactorsCh {
		if len(tri.Factors()) > *numFactors {
			fmt.Println(tri.Number())
			cancel()
			break
		}
	}
}
