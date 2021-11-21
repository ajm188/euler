package roman

import (
	"fmt"
	"testing"
)

func TestString(t *testing.T) {
	t.Parallel()

	tests := []struct {
		in  uint
		out string
	}{
		{0, ""},
		{1, "I"},
		{2, "II"},
		{3, "III"},
		{4, "IV"},
		{5, "V"},
		{8, "VIII"},
		{9, "IX"},
		{29, "XXIX"},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("%d", tt.in), func(t *testing.T) {
			t.Parallel()

			out := Numeral(tt.in).String()
			if out != tt.out {
				t.Errorf("Numeral.String(%d) got %s want %s", tt.in, out, tt.out)
			}
		})
	}
}
