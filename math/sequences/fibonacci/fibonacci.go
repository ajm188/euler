package fibonacci

import (
	"math/big"

	"github.com/ajm188/euler/math"
)

// BigFibonacciUntil generates big.Int values of the fibonacci sequence until
// the condition function is satisfied.
//
// The condition function takes the next value and index of the fibonnaci
// sequence, which will be sent on the returned channel if the condition
// evaluates to false.
func BigFibonacciUntil(bufsize uint, f func(n *big.Int, i int) bool) <-chan *big.Int {
	ch := make(chan *big.Int, bufsize)
	go func() {
		defer close(ch)

		a := big.NewInt(0)
		b := big.NewInt(1)
		i := 1

		for !f(math.CloneBigInt(b), i) {
			ch <- math.CloneBigInt(b)
			var next big.Int
			next.Add(a, b)
			a, b = b, &next
			i++
		}
	}()

	return ch
}
