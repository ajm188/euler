package main

import "testing"

func TestIsPalindrome(t *testing.T) {
	t.Parallel()

	tests := []struct {
		in  string
		out bool
	}{
		{"1001", true},
		{"100", false},
		{"1234321", true},
		{"9409", false},
		{"1", true},
		{"11", true},
		{"", true},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.in, func(t *testing.T) {
			out := isPalindrome(tt.in)
			if out != tt.out {
				t.Errorf("isPalindrome(%v) got %v want %v", tt.in, out, tt.out)
			}
		})
	}
}
