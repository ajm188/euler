package main

import (
	"flag"
	"fmt"
	stdio "io"
	"log"
	"sort"
	"strings"

	"github.com/ajm188/euler/io"
)

func score(name string) int {
	tally := 0
	for _, c := range name {
		tally += (int(c) - 64) // A = 1, B = 2, ..., Z = 26
	}

	return tally
}

func main() {
	path := flag.String("path", "names.txt", "")
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

	names := strings.Split(string(data), ",")
	sort.Strings(names)

	totalScore := 0
	for i, name := range names {
		pos := i + 1
		totalScore += pos * score(strings.ReplaceAll(name, `"`, ""))
	}

	fmt.Println(totalScore)
}
