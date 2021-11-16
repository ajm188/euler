package text

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
