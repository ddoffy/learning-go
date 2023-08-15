package main

import (
	"strings"
	"testing"
)

func TestAdd(t *testing.T) {

	counters, err := countLetters(strings.NewReader("Hello, World!"))
	want := map[string]int{"H": 1, "e": 1, "l": 3, "o": 2, ",": 1, " ": 1, "W": 1, "r": 1, "d": 1, "!": 1}
	if err != nil {
		t.Errorf("got %q, wanted %q", counters, want)
	}

	for k, v := range counters {
		if want[k] != v {
			t.Errorf("got %q, wanted %q", counters, want)
		}
	}

	// got := Add(4, 6)
	// want := 10
	//
	//	if got != want {
	//		t.Errorf("got %q, wanted %q", got, want)
	//	}
}
