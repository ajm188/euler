package main

import (
	"fmt"

	"github.com/ajm188/euler/math/sequences/collatz"
)

func main() {
	sequences := make(map[int]*collatz.Sequence, 999_999)
	for i := 1; i < 1_000_000; i++ {
		sequences[i] = collatz.NewSequence(i)
	}

	longest, longestStart := 0, 0
	for len(sequences) > 1 {
		var toDelete []int
		for start, seq := range sequences {
			if _, done := seq.Step(); done {
				if seq.Terms() > longest {
					longestStart = start
					longest = seq.Terms()
				}

				toDelete = append(toDelete, start)
			}
		}

		for _, key := range toDelete {
			delete(sequences, key)
		}
	}

	for start, seq := range sequences {
		if terms := seq.Terms(); terms >= longest {
			longestStart = start
			longest = seq.Terms()
		}
	}

	fmt.Printf("start: %d; terms: %d\n", longestStart, longest)
}
