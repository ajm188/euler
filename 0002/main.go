package main

import (
	"flag"
	"fmt"
)

func fib(n int) chan int {
	ch := make(chan int, n)

	go func() {
		defer close(ch)

		a, b := 1, 2
		ch <- a
		ch <- b

		for b < n {
			a, b = b, (a + b)
			ch <- b
		}
	}()

	return ch
}

func main() {
	limit := flag.Int("limit", 10, "")

	flag.Parse()

	ch := fib(*limit)
	sum := 0
	for x := range ch {
		if x%2 == 0 {
			sum += x
		}
	}

	fmt.Println(sum)
}
