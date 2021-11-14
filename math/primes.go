package math

import stdmath "math"

func IsPrime(n int) bool {
	for i := 2; i <= int(stdmath.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func FirstNPrimes(n int) chan int {
	ch := make(chan int, n)
	go func() {
		defer close(ch)

		count := 0
		for i := 2; count < n; i++ {
			if IsPrime(i) {
				ch <- i
				count++
			}
		}
	}()

	return ch
}
