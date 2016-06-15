package main

import (
	"fmt"
	"testing"
)

func TestReverseWords(t *testing.T) {
	type Pair struct {
		input, expected string
	}

	pairs := []Pair{{"Hello World", "World Hello"}}
	for _, p := range pairs {
		result, err := ReverseWords([]byte(p.input))
		if err != nil {
			t.Fatal(err)
		}
		if p.expected != result {
			fmt.Println("Input: %#v", p.input)
			t.Fatalf("Expected %#v, got %#v", p.expected, result)
		}
	}
}
