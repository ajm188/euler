package main

import (
	"fmt"
	"math"
)

/*
Pre:
	a^2 + b^2 = c^2
	a + b + c = 1000

Therefore:
	1. c = 1000 - (a + b)
	2. a^2 + b^2 = (1000 - (a + b))^2
	(substitute 1 into pythogorean theorem)
	3. a^2 + b^2 = 1000^2 - 2(1000(a + b)) + (a + b)^2
	4. 			 = 1_000_000 - 2000a - 2000b + (a^2 + 2ab + b^2)
	5. 			 = 1_000_000 - 2000(a+b) + a^2 + 2ab + b^2
	6. 0 = 1_000_000 - 2000(a + b) +2ab
	7. 2000a + 2000b = 1_000_000 + 2ab
	8. 2000a + 2000b - 2ab = 1_000_000
	9. 1000a + 1000b - ab = 500_000

Furthermore, a < b < c, so if a=1, then the largest b can be is (1000-1)/2 = 499
*/
func solve() (a, b, c int) {
	for b := 2; b <= 499; b++ {
		for a := 1; a < b; a++ {
			lhs := (1000 * a) + (1000 * b) - (a * b)
			if lhs == 500_000 {
				c := int(math.Sqrt(float64((a * a) + (b * b))))
				return a, b, c
			}
		}
	}

	panic("no solution found")
}

func main() {
	a, b, c := solve()
	fmt.Println(a * b * c)
}
