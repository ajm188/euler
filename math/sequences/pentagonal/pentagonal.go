package pentagonal

// Sequence represents one instance of a sequence of pentagonal numbers.
//
// The pentagonal number sequence is defined as:
//	h_n = n(3n - 1)/2
type Sequence struct {
	p int
	n int
}

// NewSequence returns a Sequence at the beginning of the pentagonal numbers.
func NewSequence() *Sequence {
	return &Sequence{
		n: 1,
		p: 1,
	}
}

// Step advances the sequence by one iteration, returning the new P value.
func (s *Sequence) Step() int {
	s.n++
	s.p = (s.n * (3*s.n - 1)) / 2
	return s.p
}

// P returns the current value of the sequence, that is P_n.
func (s *Sequence) P() int { return s.p }

// N returns the current iteration value of the sequence.
func (s *Sequence) N() int { return s.n }
