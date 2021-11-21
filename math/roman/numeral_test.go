package roman

import (
	"fmt"
	"testing"
)

func TestIsValid(t *testing.T) {
	t.Parallel()

	tests := []struct {
		in    string
		valid bool
	}{
		{"XVI", true},
		{"XIIIIII", true},
		{"IIIIIIIIIIIIIIII", false},
		{"VVVI", false},

		{"XIX", true},
		{"XVIV", false},
		{"XIIIIIIIII", true},
		{"XVIIII", true},

		{"IIII", true},
		{"IV", true},

		{"IL", false},
		{"XXXXIIIIIIIII", true},
		{"XXXXVIIII", true},
		{"XXXXIX", true},
		{"XLIIIIIIIII", true},
		{"XLVIIII", true},
		{"XLIX", true},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.in, func(t *testing.T) {
			t.Parallel()

			valid := IsValid(tt.in)
			if valid != tt.valid {
				t.Errorf("IsValid(%s) got %v want %v", tt.in, valid, tt.valid)
			}
		})
	}
}

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
