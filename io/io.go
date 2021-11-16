package io

import (
	stdio "io"
	"os"
)

// Open opens the file at path for reading, or returns stdin if path is the
// empty string.
func Open(path string) (stdio.ReadCloser, error) {
	if path == "" {
		return os.Stdin, nil
	}

	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	return f, nil
}
