package main

import (
	"context"
	"flag"
	"fmt"
	stdmath "math"

	"github.com/ajm188/euler/math"
)

func primes(ctx context.Context, n int) chan int {
	ch := make(chan int, n)
	go func() {
		defer close(ch)

		for i := n - 1; i >= 2; i-- {
			select {
			case <-ctx.Done():
				return
			default:
			}

			if math.IsPrime(i) {
				ch <- i
			}
		}
	}()

	return ch
}

func main() {
	n := flag.Int("n", 13195, "")

	flag.Parse()

	ctx, cancel := context.WithCancel(context.Background())
	ch := primes(ctx, int(stdmath.Sqrt(float64(*n))))
	for x := range ch {
		if *n%x == 0 {
			cancel()
			fmt.Println(x)
			break
		}
	}

}
