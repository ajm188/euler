package math

import (
	"fmt"
	"math/big"

	"github.com/ajm188/euler/text"
)

// IsLychrel returns true if n is a Lychrel number within maxIters.
//
// Let R be a function that reverses the digits in a number. A number X such
// that (X + R(X)) is a palindrome, or (X + R(X)) + R(X + R(X)), or so on in
// this manner, is **not** a Lychrel number.
func IsLychrel(n int, maxIters int) bool {
	x := big.NewInt(int64(n))
	for i := 1; i <= maxIters; i++ {
		s := x.String()
		rev := text.Reverse(s)
		revX, ok := big.NewInt(0).SetString(rev, 10)
		if !ok {
			panic(fmt.Sprintf("failed to reverse %s (rev=%s)", s, rev))
		}

		x = x.Add(x, revX)

		if text.IsPalindrome(x.String()) {
			return false
		}
	}

	return true
}
