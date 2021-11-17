package text

// IsPalindrome returns true if s is a palindrome; that is, if s == reverse(s).
func IsPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		j := (len(s) - 1) - i
		if s[i] != s[j] {
			return false
		}
	}

	return true
}

// Reverse returns s backwards.
func Reverse(s string) string {
	buf := make([]byte, 0, len(s))
	for i := len(s) - 1; i >= 0; i-- {
		buf = append(buf, s[i])
	}

	return string(buf)
}

// WordScore returns a scoring of the fully-uppercased string s.
//
// A score is defined as the sum of each letter in s's position in the alphabet,
// so, A=1, B=2, and so on.
//
// WordScore is undefined for strings that do not match the regexp /^[A-Z].*$/.
func WordScore(s string) (score int) {
	for _, c := range s {
		score += int(c) - 64 // A = 1, B = 2, ..., Z = 26
	}

	return score
}
