package math

import (
	"math/big"
	"strconv"
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
