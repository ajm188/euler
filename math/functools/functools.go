package functools

// MapInts returns a new slice containing f(x) for each x in the input slice.
//
// Haskell equivalent is:
//	map f xs
func MapInts(xs []int, f func(x int) int) []int {
	ys := make([]int, len(xs))
	for i, x := range xs {
		ys[i] = f(x)
	}

	return ys
}

// ReduceInts returns the result of applying acc=f(x, acc) from left to right
// for each x in te input slice.
//
// Haskell equivalent is:
//	foldl f init xs
func ReduceInts(xs []int, init int, f func(x int, acc int) int) int {
	acc := init
	for _, x := range xs {
		acc = f(x, acc)
	}

	return acc
}
