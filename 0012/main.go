package main

import (
	"context"
	"flag"
	"fmt"
	"math"
	"sort"
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

	for i := 1; i <= int(math.Sqrt(float64(fi.n))); i++ {
		if fi.n%i == 0 {
			fi.factors = append(fi.factors, i)
			if div := fi.n / i; div != i { // if i == sqrt(fi.n), don't add it twice
				fi.factors = append(fi.factors, div)
			}
		}
	}

	sort.Ints(fi.factors)
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
