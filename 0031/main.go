package main

import "fmt"

var coinValues = []int{
	200,
	100,
	50,
	20,
	10,
	5,
	2,
	1,
}

func incr(v []int, i int) []int {
	r := make([]int, len(v))
	copy(r, v)
	if i < len(r) {
		r[i]++
	}
	return r
}

func coinsum(coins []int) (sum int) {
	if len(coins) != len(coinValues) {
		panic(fmt.Sprintf("impossible: incorrect number of coins: %d", len(coins)))
	}

	for i, count := range coins {
		sum += coinValues[i] * count
	}

	return sum
}

func value(coins []int, target int) int {
	var f func([]int, int) int
	f = func(coins []int, i int) int {
		sum := coinsum(coins)
		if sum == target {
			return 1
		} else if sum > target || i >= len(coins) {
			return 0
		}

		if len(coins) == i-1 {
			return f(incr(coins, i), i)
		}

		return f(incr(coins, i), i) + f(incr(coins, len(coins)+1 /* just copy, no incr */), i+1)
	}

	return f(coins, 0)

}

func main() {
	coins := make([]int, len(coinValues))
	fmt.Println(value(coins, 200))
}
