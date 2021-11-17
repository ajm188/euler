package hexagonal

// Sequence represents one instance of a sequence of hexagonal numbers.
//
// The hexagonal number sequence is defined as:
//	h_n = n(2n - 1)
type Sequence struct {
	n int
	h int
}

// NewSequence returns a Sequence at the beginning of the hexagonal numbers.
func NewSequence() *Sequence {
	return &Sequence{
		n: 1,
		h: 1,
	}
}

// Step advances the sequence by one iteration, returning the new H value.
func (s *Sequence) Step() int {
	s.n++
	s.h = s.n * (2*s.n - 1)
	return s.h
}

// H returns the current value of the sequence, that is H_n.
func (s *Sequence) H() int { return s.h }

// N returns the current iteration value of the sequence.
func (s *Sequence) N() int { return s.n }
