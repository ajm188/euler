package math

// SumInts returns the sum of a slice of ints.
func SumInts(nums []int) (sum int) {
	for _, i := range nums {
		sum += i
	}

	return sum
}
