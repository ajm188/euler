package main

import (
	"flag"
	"fmt"

	"github.com/ajm188/euler/math/sequences/hexagonal"
	"github.com/ajm188/euler/math/sequences/pentagonal"
	"github.com/ajm188/euler/math/sequences/triangle"
)

var (
	triangles   = map[int]bool{}
	pentagonals = map[int]bool{}
	hexagonals  = map[int]bool{}
)

func main() {
	n := flag.Int("n", 40755, "")
	flag.Parse()

	t := triangle.NewSequence()
	p := pentagonal.NewSequence()
	h := hexagonal.NewSequence()

	triangles[t.T()] = true
	pentagonals[p.P()] = true
	hexagonals[h.H()] = true

	for t.T() <= *n {
		t_n := t.Step()
		triangles[t_n] = true
	}

	for p.P() <= *n {
		p_n := p.Step()
		pentagonals[p_n] = true
	}

	for h.H() <= *n {
		h_n := h.Step()
		hexagonals[h_n] = true
	}

	for {
		t_n := t.T()
		for t_n >= p.P() {
			// advance pentagonals until at least t_n
			p_n := p.Step()
			pentagonals[p_n] = true
		}

		if _, ok := pentagonals[t_n]; !ok {
			goto deferred
		}

		for t_n >= h.H() {
			h_n := h.Step()
			hexagonals[h_n] = true
		}

		if _, ok := hexagonals[t_n]; !ok {
			goto deferred
		}

		// t_n is both hexagonal and pentagonal
		break

	deferred:
		t_n = t.Step()
		triangles[t_n] = true
	}

	fmt.Printf("T(n=%d)\tP(n=%d)\tH(n=%d)\t%d\n", t.N(), p.N(), h.N(), t.T())
}
