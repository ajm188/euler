package main

import (
	"flag"
	"fmt"
	"math/big"
)

func main() {
	denomLimit := flag.Int("denom-limit", 100, "")
	trivial := flag.String("trivial", "0", "do not consider these characters for reduction")
	debug := flag.Bool("debug", false, "")
	flag.Parse()

	product := big.NewRat(1, 1)

	for denom := 2; denom < *denomLimit; denom++ {
		for numer := 1; numer < denom; numer++ {
			r1 := big.NewRat(int64(numer), int64(denom))
			n, d := fmt.Sprintf("%d", numer), fmt.Sprintf("%d", denom)

			for i := 0; i < len(n); i++ {
				for j := 0; j < len(d); j++ {
					if string(n[i]) == *trivial {
						continue
					}
					if n[i] == d[j] {
						n2 := n[:i] + n[i+1:]
						d2 := d[:j] + d[j+1:]

						var r2 big.Rat
						r2.SetString(fmt.Sprintf("%s/%s", n2, d2))

						if r1.Cmp(&r2) == 0 {
							if *debug {
								fmt.Printf("found\t%d/%d\n", numer, denom)
							}
							product.Mul(product, r1)
						}
					}
				}
			}
		}
	}

	fmt.Println(product)
}
