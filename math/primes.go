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

// PrimesUntil generates primes until the condition func returns false.
//
// Params:
//	n: capacity for the channel
//	f(n, i): n is the current number being consider for primeness. It may or may
//			 not be prime.
//			 i is the number of primes that have been generated _before_
//			 considering n for primeness.
func PrimesUntil(n int, f func(n, i int) bool) chan int {
	ch := make(chan int, n)
	go func() {
		defer close(ch)

		count := 0
		done := false
		for i := 2; !done; i++ {
			done = !f(i, count)

			if IsPrime(i) {
				ch <- i
				count++
			}
		}
	}()

	return ch
}
