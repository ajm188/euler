package main

import (
	"flag"
	"fmt"
	stdio "io"
	"strconv"
	"strings"

	"github.com/ajm188/euler/io"
)

var cache = map[string]int64{}

func search(
	matrix [][]int64, row, col int,
	step func(matrix [][]int64, row, col int) [][2]int,
) int64 {
	if x, ok := cache[fmt.Sprintf("%d/%d", row, col)]; ok {
		return x
	}

	val := matrix[row][col]
	var min *int64
	for _, next := range step(matrix, row, col) {
		if v := search(matrix, next[0], next[1], step); min == nil || v < *min {
			min = &v
		}
	}

	if min != nil {
		val += *min
	}

	cache[fmt.Sprintf("%d/%d", row, col)] = val

	return val
}

func main() {
	path := flag.String("path", "matrix.txt", "")
	flag.Parse()

	f, err := io.Open(*path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	data, err := stdio.ReadAll(f)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(data), "\n")
	lines = lines[:len(lines)-1]
	matrix := make([][]int64, len(lines))
	for i, line := range lines {
		numbers := strings.Split(line, ",")
		matrix[i] = make([]int64, len(numbers))

		for j, number := range numbers {
			x, err := strconv.ParseInt(number, 10, 64)
			if err != nil {
				panic(fmt.Sprintf("cannot parse %s (row %d, pos %d) as int64: %s", number, i, j, err))
			}

			matrix[i][j] = x
		}
	}

	minPath := search(matrix, 0, 0, func(matrix [][]int64, row, col int) [][2]int {
		next := make([][2]int, 0, 2)
		if row < len(matrix)-1 {
			next = append(next, [2]int{row + 1, col})
		}

		if col < len(matrix[row])-1 {
			next = append(next, [2]int{row, col + 1})
		}

		return next
	})
	fmt.Println(minPath)
}
