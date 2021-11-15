package main

import (
	"fmt"
	"strings"
)

var (
	numberWords = map[int]string{
		1:    "one",
		2:    "two",
		3:    "three",
		4:    "four",
		5:    "five",
		6:    "six",
		7:    "seven",
		8:    "eight",
		9:    "nine",
		10:   "ten",
		11:   "eleven",
		12:   "twelve",
		13:   "thirteen",
		14:   "fourteen",
		15:   "fifteen",
		16:   "sixteen",
		17:   "seventeen",
		18:   "eighteen",
		19:   "nineteen",
		20:   "twenty",
		30:   "thiity",
		40:   "forty",
		50:   "fifty",
		60:   "sixty",
		70:   "seventy",
		80:   "eighty",
		90:   "ninety",
		100:  "hundred",
		1000: "thousand",
	}
	stringNumberWords map[string]string
)

func init() {
	stringNumberWords = make(map[string]string, len(numberWords))
	for x, word := range numberWords {
		stringNumberWords[fmt.Sprintf("%d", x)] = word
	}
}

func toEnglish(x int) string {
	return digitsToWords(fmt.Sprintf("%d", x))
}

func digitsToWords(s string) string {
	switch len(s) {
	case 4: // cheating; according to the problem 1000 is the largest number, so we can get away with hardcoding here.
		return "onethousand"
	case 3:
		word := &strings.Builder{}
		word.WriteString(stringNumberWords[s[0:1]])
		word.WriteString("hundred")
		if s[1:] == "00" {
			return word.String()
		}

		word.WriteString("and")
		word.WriteString(digitsToWords(strings.TrimLeft(s[1:], "0")))
		return word.String()
	case 2:
		switch s[0] {
		case '1': // "ten", "eleven", ..., "nineteen"
			return stringNumberWords[s]
		default:
			switch s[1] {
			case '0': // "twenty", "thirty", ..., "ninety"
				return stringNumberWords[s]
			default:
				word := &strings.Builder{}
				word.WriteString(stringNumberWords[s[0:1]+"0"])
				word.WriteString(stringNumberWords[s[1:]])
				return word.String()
			}
		}
	case 1: // "one", "two", ..., "nine"
		return stringNumberWords[s]
	default:
		return ""
	}
}

func main() {
	count := 0
	for i := 1; i <= 1000; i++ {
		word := toEnglish(i)
		count += len(word)
	}

	fmt.Println(count)
}
