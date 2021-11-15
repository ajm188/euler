package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strings"
)

func open(path string) (io.Reader, func() error, error) {
	if path == "" {
		return os.Stdin, func() error { return nil }, nil
	}

	f, err := os.Open(path)
	if err != nil {
		return nil, nil, err
	}

	return f, f.Close, nil
}

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

	r, closer, err := open(*path)
	if err != nil {
		log.Fatal(err)
	}
	defer closer()

	data, err := io.ReadAll(r)
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
