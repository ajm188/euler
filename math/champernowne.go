package math

import (
	"fmt"
	"strings"
)

// Champernowne returns a string representation of the first n digits of the
// fractional component of the Champernowne constant [1].
//
// [1]: https://en.wikipedia.org/wiki/Champernowne_constant.
func Champernowne(n int) string {
	var buf strings.Builder
	for i := 1; buf.Len() <= n; i++ {
		fmt.Fprintf(&buf, "%d", i)
	}

	s := buf.String()
	return s[:n]
}
