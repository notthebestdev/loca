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
	defer func() {
		if err := tmp.Close(); err != nil {
			t.Error(err)
		}
		if err := os.Remove(tmp.Name()); err != nil {
			t.Error(err)
		}
	}()

	_, err = tmp.WriteString("line 1\nline 2\nline 3\n")
	if err != nil {
		t.Fatal(err)
	}

	lines := countLines(tmp.Name())

	if lines != 3 {
		t.Fatalf("expected 3 lines, got %d", lines)
	}
}
