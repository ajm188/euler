/*
Package roman provides types and functions for dealing with Roman numerals.

All Roman numerals must follow three basic rules to be considered valid:

	1. Numerals must be arranged in descending order of size.
	2. M, C, and X cannot be equalled or exceeded by smaller denominations.
	3. D, L, and V can each only appear once.

Roman numerals may also be written using subtractive combinations. For example,
four may be written as IIII or as IV, because it is one (I) before five (V).
Generally, the Romans preferred to use as few numerals as possible when writing
numbers, so the preferred form of four would be IV.

By medieval times, it had become standard practice to avoid more than three
consecutive identical numerals, opting to replace them with a more compact
subtractive combination.

In addition to the three basic rules, if subtractive combinations are used to
display a number, the following four rules must also be followed:

	1. Only one I, X, and C can be used as the leading numeral in part of a subtractive pair.
	2. I can only be placed before V and X.
	3. X can only be placed before L and C.
	4. C can only be placed before D and M.

For example, IL would be an invalid form of forty-nine (49). XXXXIIIIIIIII,
XXXXVIIII, XXXXIX, XLIIIIIIIII, XLVIIII, and XLIX are all valid forms, with the
last being the preferred form due to its minimization of numerals used.

It is also expected, but not required, that higher denominations be used
wherever possible.
*/
package roman

import (
	"errors"
	"fmt"
	"strings"
)

// ErrInvalidFormat is returned when Parse is called on an invalid numeral
// representation.
var ErrInvalidFormat = errors.New("invalid numeral format")

// Numeral is number that can be expressed in Roman numerals.
// Negative numbers and zero have undefined behavior.
type Numeral uint

// String returns the preferred (minimal) representation of the number in Roman
// numerals.
func (n Numeral) String() string {
	return format(uint(n))
}

func format(n uint) string {
	var buf strings.Builder
	for n >= 1000 {
		buf.WriteByte('M')
		n -= 1000
	}

	if n >= 900 {
		buf.WriteString("CM")
		n -= 900
	}

	for n >= 500 {
		buf.WriteByte('D')
		n -= 500
	}

	if n >= 400 {
		buf.WriteString("CD")
		n -= 400
	}

	for n >= 100 {
		buf.WriteByte('C')
		n -= 100
	}

	if n >= 90 {
		buf.WriteString("XC")
		n -= 90
	}

	for n >= 50 {
		buf.WriteByte('L')
		n -= 50
	}

	if n >= 40 {
		buf.WriteString("XL")
		n -= 40
	}

	for n >= 10 {
		buf.WriteByte('X')
		n -= 10
	}

	if n >= 9 {
		buf.WriteString("IX")
		n -= 9
	}

	for n >= 5 {
		buf.WriteByte('V')
		n -= 5
	}

	if n >= 4 {
		buf.WriteString("IV")
		n -= 4
	}

	for n > 0 {
		buf.WriteByte('I')
		n--
	}

	return buf.String()
}

// IsValid returns true if the given string is a valid Roman numeral.
func IsValid(s string) bool {
	return validate(s) == nil
}

// Parse converts a string representation of a Roman numeral into a Numeral.
// It returns an error if the numeral is in an invalid format.
func Parse(s string) (Numeral, error) {
	if err := validate(s); err != nil {
		return 0, err
	}

	n := uint(0)
	i := 0
	for i < len(s)-1 {
		if numeralValues[s[i]] < numeralValues[s[i+1]] {
			n += numeralValues[s[i+1]] - numeralValues[s[i]]
			i += 2
			continue
		}

		n += numeralValues[s[i]]
		i++
	}

	if i < len(s) {
		n += numeralValues[s[i]]
	}

	return Numeral(n), nil
}

var (
	orderedNumerals = []byte{'I', 'V', 'X', 'L', 'C', 'D', 'M'}
	numeralValues   = map[byte]uint{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
)

func validate(s string) error {
	if len(s) == 0 {
		return nil
	}

	s = strings.ToUpper(s)

	globalCounts := map[byte]int{}
	for _, c := range orderedNumerals {
		globalCounts[c] = 0
	}

	checkBasic := func(current byte, count int) error {
		// Basic rule #2: M, C, and X cannot be equalled or exceeded by
		// smaller denominations.
		for _, c := range []byte{'M', 'C', 'X'} {
			if numeralValues[current] >= numeralValues[c] {
				// Not a smaller denomination
				continue
			}

			if count*int(numeralValues[current]) >= int(numeralValues[c]) {
				return fmt.Errorf("%w: %d %c's should be (partially) replaced with equivalent %c", ErrInvalidFormat, count, current, c)
			}
		}

		// Basic rule #3: D, L, and V can each appear only once.
		for _, c := range []byte{'D', 'L', 'V'} {
			if globalCounts[c] > 1 {
				return fmt.Errorf("%w: at most 1 %c may appear", ErrInvalidFormat, c)
			}
		}

		return nil
	}

	current := s[0]
	count := 1
	globalCounts[current]++

	for i := 1; i < len(s); i++ {
		globalCounts[s[i]]++

		if numeralValues[s[i]] > numeralValues[current] { // Subtractive pair
			// Subtractive rule #1: Only one I, X, and C can be used as the
			// leading numeral in part of a subtractive pair.
			switch current {
			case 'I', 'X', 'C':
			default:
				return fmt.Errorf("%w: %c may not be used as the leading numeral in a subtractive pair", ErrInvalidFormat, current)
			}
			if count > 1 {
				return fmt.Errorf("%w: only one (have %d) %c allowed as leading numeral in a subtractive pair", ErrInvalidFormat, count, current)
			}

			switch current {
			case 'I':
				// Subtractive rule #2: I can only be placed before V and X.
				switch s[i] {
				case 'V', 'X':
				default:
					return fmt.Errorf("%w: I can only be placed before V and X (not %c)", ErrInvalidFormat, s[i])
				}
			case 'X':
				// Subtractive rule #3: X can only be placed before L and C.
				switch s[i] {
				case 'L', 'C':
				default:
					return fmt.Errorf("%w: X can only be placed before L and C (not %c)", ErrInvalidFormat, s[i])
				}
			case 'C':
				// Subtractive rule #4: C can only be placed before D and M.
				switch s[i] {
				case 'D', 'M':
				default:
					return fmt.Errorf("%w: C can only be placed before D and M (not %c)", ErrInvalidFormat, s[i])
				}
			}
		} else if numeralValues[s[i]] < numeralValues[current] {
			if err := checkBasic(current, count); err != nil {
				return err
			}

			current = s[i]
			count = 1
		} else {
			count++
		}
	}

	if err := checkBasic(current, count); err != nil {
		return err
	}

	return nil
}
