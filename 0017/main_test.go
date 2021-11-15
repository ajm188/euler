package main

import (
	"strings"
	"testing"
)

func TestToEnglish(t *testing.T) {
	t.Parallel()

	tests := []struct {
		in  int
		out string
	}{
		{100, "one hundred"},
		{101, "one hundred and one"},
		{23, "twentythree"},
		{1000, "onethousand"},
		{11, "eleven"},
		{914, "nine hundred and fourteen"},
		{420 /* sick */, "four hundred and twenty"},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.out, func(t *testing.T) {
			t.Parallel()

			out := toEnglish(tt.in)
			expected := strings.ReplaceAll(tt.out, " ", "")
			if out != expected {
				t.Errorf("toEnglish(%v) got %v want %v", tt.in, out, expected)
			}
		})
	}
}
