package main

import (
	"bufio"
	"fmt"
	"os"
)

func countLines(file string) int {
	f, err := os.Open(file)
	if err != nil {
		return 0
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Fprintf(os.Stderr, "Error closing file: %v\n", err)
		}
	}()

	scanner := bufio.NewScanner(f)
	lines := 0
	for scanner.Scan() {
		lines++
	}
	return lines
}
