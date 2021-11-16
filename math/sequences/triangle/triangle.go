package triangle

// Sequence represents one instance of a sequence of triangle numbers.
//
// The triangle number sequence is defined as:
//	t_n = (n(n+1)) / 2
type Sequence struct {
	t int
	n int
}

// NewSequence returns a Sequence at the beginning of the triangle numbers.
func NewSequence() *Sequence {
	return &Sequence{
		t: 1,
		n: 1,
	}
}

// Step advances the sequence by one iteration, returning the new T value.
func (s *Sequence) Step() int {
	s.n++
	t_n := (s.n * (s.n + 1)) / 2
	s.t = t_n

	return s.t
}

// T returns the current value of the sequence, that is T_n.
func (s *Sequence) T() int { return s.t }

// N returns the current iteration value of the sequence.
func (s *Sequence) N() int { return s.n }
