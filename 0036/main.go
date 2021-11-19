package main

import (
	"flag"
	"fmt"
	"strconv"

	"github.com/ajm188/euler/text"
)

func main() {
	max := flag.Int64("max", 1_000_000, "")
	flag.Parse()

	var sum int64
	for i := int64(0); i < *max; i++ {
		s10, s2 := strconv.FormatInt(i, 10), strconv.FormatInt(i, 2)
		if text.IsPalindrome(s10) && text.IsPalindrome(s2) {
			sum += i
		}
	}

	fmt.Println(sum)
}
