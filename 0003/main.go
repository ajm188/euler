package main

import (
	"context"
	"flag"
	"fmt"
	"math"
)

func isPrime(n int) bool {
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

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

			if isPrime(i) {
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
	ch := primes(ctx, int(math.Sqrt(float64(*n))))
	for x := range ch {
		if *n%x == 0 {
			cancel()
			fmt.Println(x)
			break
		}
	}

}
