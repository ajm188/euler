package collatz

// Sequence represents a Collatz[1] sequence.
//
// The Collatz sequence is an iterative sequence with the following definiton:
//	n => (n/2) if n is even
//	n => 3n + 1 if n is odd
//
// While not yet proven (see Collatz conjecture), it is currently believed that
// the sequence always ends at 1, regardless of the starting value of n.
//
// [1]: https://en.wikipedia.org/wiki/Collatz_conjecture
type Sequence struct {
	cur   int
	terms int
}

// NewSequence returns a new Collatz sequence beginning at n.
func NewSequence(n int) *Sequence {
	return &Sequence{
		cur:   n,
		terms: 1,
	}
}

// Step advances the sequence by one iteration. It returns the next value in the
// sequence and a boolean indicating whether the sequence has terminated
// (reached 1).
func (s *Sequence) Step() (next int, done bool) {
	if s.cur == 1 {
		return s.cur, true
	}

	if s.cur%2 == 0 {
		s.cur /= 2
	} else {
		s.cur = (3 * s.cur) + 1
	}

	s.terms++
	return s.cur, false
}

// Terms returns the number of terms in the sequence so far.
// Once Step() returns (_, true), this function returns a constant value.
func (s *Sequence) Terms() int { return s.terms }
