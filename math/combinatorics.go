package math

import "math/big"

// Factorial returns n!
func Factorial(n int) *big.Int {
	if n <= 1 {
		return big.NewInt(1)
	}

	prod := big.NewInt(int64(n))
	for i := n - 1; i > 1; i-- {
		prod.Mul(prod, big.NewInt(int64(i)))
	}

	return prod
}

// Choose returns the number of k-combinations of n.
//
// See https://en.wikipedia.org/wiki/Combination.
func Choose(n int, k int) *big.Int {
	if k > n {
		return big.NewInt(0)
	}

	if n == k { // optimization to avoid computing large factorials.
		return big.NewInt(1)
	}

	factN := Factorial(n)
	factK := Factorial(k)

	// (n!) / (k!(n-k)!)
	return factN.Div(factN, factK.Mul(factK, Factorial(n-k)))
}
