package main

import (
	"flag"
	"fmt"
	stdio "io"
	"log"
	"sort"
	"strings"

	"github.com/ajm188/euler/io"
	"github.com/ajm188/euler/text"
)

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
		totalScore += pos * text.WordScore(strings.ReplaceAll(name, `"`, ""))
	}

	fmt.Println(totalScore)
}
