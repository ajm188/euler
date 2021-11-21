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

import "strings"

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

func IsValid(s string) bool {
	panic("unimplemented!")
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
	numeralIndexes map[byte]int
)

func init() {
	numeralIndexes = make(map[byte]int, len(orderedNumerals))
	for i, n := range orderedNumerals {
		numeralIndexes[n] = i
	}
}

// Parse converts a string representation of a Roman numeral into a Numeral.
// It returns an error if the numeral is in an invalid format.
func Parse(s string) (Numeral, error) {
	// TODO: implement IsValid and return error
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
