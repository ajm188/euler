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
