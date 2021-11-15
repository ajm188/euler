package math

import (
	"fmt"
	"testing"
)

func TestChoose(t *testing.T) {
	t.Parallel()

	tests := []struct {
		n, k int
		out  string
	}{
		{4, 2, "6"},
		{40, 20, "137846528820"},
	}

	for _, tt := range tests {
		tt := tt
		name := fmt.Sprintf("C(n=%d,k=%d)", tt.n, tt.k)
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			out := Choose(tt.n, tt.k)
			if out.String() != tt.out {
				t.Errorf("%s got %v want %v", name, out.String(), tt.out)
			}
		})
	}
}
