package main

import (
	"fmt"
	"math/rand"
)

func getRandomPhrase(lines int) string {
	phrases := []struct {
		format string
		key    string
	}{
		{"That's like reading the entire Lord of the Rings trilogy %d times over!", ""},
		{"You could read War and Peace %d times with that many lines!", ""},
		{"That's equivalent to %d copies of the Harry Potter series!", ""},
		{"You've written enough code to fill %d Python textbooks!", ""},
		{"That's like a small novel repeated %d times!", ""},
		{"You could teach %d coding bootcamps with this!", ""},
		{"That's %d times the length of the average Wikipedia article!", ""},
		{"You've written more than %d copies of Don Quixote!", ""},
		{"That's like %d complete video game scripts!", ""},
		{"You could fill %d technical manuals with this code!", ""},
	}
	times := lines / 150
	if times == 0 {
		return ""
	}
	return fmt.Sprintf(phrases[rand.Intn(len(phrases))].format, times)
}
