package main

import (
	"flag"
	"fmt"
	stdio "io"
	"log"
	"strings"

	"github.com/ajm188/euler/io"
	"github.com/ajm188/euler/math/sequences/triangle"
	"github.com/ajm188/euler/text"
)

var triangleNumbers = map[int]bool{}

func main() {
	path := flag.String("path", "words.txt", "")
	flag.Parse()

	r, err := io.Open(*path)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	data, err := stdio.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}

	numTriangles := 0
	triangles := triangle.NewSequence()
	triangleNumbers[triangles.T()] = true

	words := strings.Split(string(data), ",")
	for _, word := range words {
		word := strings.ReplaceAll(word, `"`, "")
		score := text.WordScore(word)

		// generate triangle numbers <= score
		for triangles.T() < score {
			x := triangles.Step()
			triangleNumbers[x] = true
		}

		if _, ok := triangleNumbers[score]; ok {
			numTriangles++
		}
	}

	fmt.Println(numTriangles)
}
