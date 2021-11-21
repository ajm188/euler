package main

import (
	"flag"
	"fmt"
	stdio "io"
	"log"
	"strings"

	"github.com/ajm188/euler/io"
	"github.com/ajm188/euler/math/roman"
)

func main() {
	path := flag.String("path", "roman.txt", "")
	flag.Parse()

	f, err := io.Open(*path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	data, err := stdio.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	var saved int
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}

		numeral, err := roman.Parse(line)
		if err != nil {
			log.Fatal(err)
		}

		saved += len(line) - len(numeral.String())
	}

	fmt.Println(saved)
}
