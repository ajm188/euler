package math

import (
	"math/big"
	"testing"
)

func TestSumDigitFactorials(t *testing.T) {
	t.Parallel()

	tests := []struct {
		in  string
		out string
	}{
		{"169", "363601"},
		{"363601", "1454"},
		{"1454", "169"},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.in, func(t *testing.T) {
			t.Parallel()

			var n big.Int
			n.SetString(tt.in, 10)
			out := SumDigitFactorials(&n)
			if out.String() != tt.out {
				t.Errorf("SumDigitFactorials(%s) got %s want %s", tt.in, out.String(), tt.out)
			}
		})
	}
}
