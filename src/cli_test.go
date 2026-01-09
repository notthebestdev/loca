package main

import (
	"os"
	"testing"
)

func TestCountLines(t *testing.T) {
	tmp, err := os.CreateTemp("", "loca-test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmp.Name())

	tmp.WriteString("line 1\nline 2\nline 3\n")

	lines := countLines(tmp.Name())

	if lines != 3 {
		t.Fatalf("expected 3 lines, got %d", lines)
	}
}
