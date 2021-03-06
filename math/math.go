package math

import (
	"fmt"
	"math/big"
	"sort"
	"strconv"
	"strings"
)

// SumInts returns the sum of a slice of ints.
func SumInts(nums []int) (sum int) {
	for _, i := range nums {
		sum += i
	}

	return sum
}

// SumDigits returns the sum of the digits in a big integer.
func SumDigits(x *big.Int) (sum int) {
	for _, ds := range x.String() {
		d, err := strconv.ParseInt(string(ds), 10, 64)
		if err != nil {
			panic(err)
		}

		sum += int(d)
	}

	return sum
}

// CloneBigInt returns a copy of n.
//
// big.Int types are not safe for shallow copies; this is the recommended way
// to safely copy a big.Int.
func CloneBigInt(n *big.Int) *big.Int {
	var ret big.Int
	ret.Set(n)
	return &ret
}

// used for comparisons to 1.
// do not call any methods that mutate the receiver (such as Add, Mul, or any
// of the Set* methods)
var _one = big.NewInt(1)

func BigFactorial(n *big.Int) *big.Int {
	product := big.NewInt(1)
	for i := big.NewInt(2); i.Cmp(n) <= 0; i.Add(i, _one) {
		product.Mul(product, i)
	}

	return product
}

// SumDigitFactorials returns the sum of the factorials of the digits of n.
func SumDigitFactorials(n *big.Int) *big.Int {
	var sum big.Int
	for i, c := range n.String() {
		var d big.Int
		if _, ok := d.SetString(string(c), 10); !ok {
			panic(fmt.Sprintf("could not parse digit %d (=%s) of %s", i, string(c), n.String()))
		}

		sum.Add(&sum, BigFactorial(&d))
	}

	return &sum
}

// Digits returns a slice of the digits of n.
func Digits(n int) []int {
	var bi big.Int
	bi.SetInt64(int64(n))
	s := bi.String()

	digits := make([]int, len(s))
	for i, c := range s {
		d, err := strconv.ParseInt(string(c), 10, 64)
		if err != nil {
			panic(fmt.Sprintf("could not parse digit %d of %s: %s", i, s, err))
		}
		digits[i] = int(d)
	}

	return digits
}

// OrderedDigitsString returns the digits of n in increasing order as a string.
func OrderedDigitsString(n int) string {
	digits := Digits(n)
	sort.Ints(digits)

	buf := strings.Builder{}
	for _, d := range digits {
		buf.WriteString(fmt.Sprintf("%d", d))
	}

	return buf.String()
}
