package main

import (
	"fmt"
	"sort"

	"github.com/ajm188/euler/math"
)

func arePerumtations(a, b, c int) bool {
	as := []byte(fmt.Sprintf("%d", a))
	bs := []byte(fmt.Sprintf("%d", b))
	cs := []byte(fmt.Sprintf("%d", c))

	sort.Slice(as, func(i, j int) bool {
		return as[i] < as[j]
	})
	sort.Slice(bs, func(i, j int) bool {
		return bs[i] < bs[j]
	})
	sort.Slice(cs, func(i, j int) bool {
		return cs[i] < cs[j]
	})

	return string(as) == string(bs) && string(bs) == string(cs)
}

func main() {
	for i := 1000; i < 10_000-3330; i++ {
		a, b, c := i, i+3330, i+3330+3330

		if !arePerumtations(a, b, c) {
			continue
		}

		if math.IsPrime(a) && math.IsPrime(b) && math.IsPrime(c) {
			fmt.Printf("(%d, %d, %d): %s\n", a, b, c, fmt.Sprintf("%d%d%d", a, b, c))
		}
	}
}
