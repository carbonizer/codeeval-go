package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestMultiplicationTables(t *testing.T) {
	type Pair struct {
		input, expected string
	}

	pairs := []Pair{{"4", strings.Join([]string{
		"1   2   3   4",
		"2   4   6   8",
		"3   6   9  12",
		"4   8  12  16",
	}, "\n")}}
	for _, p := range pairs {
		result, err := MultiplicationTables([]byte(p.input))
		if err != nil {
			t.Fatal(err)
		}
		if p.expected != result {
			fmt.Println("Input: %#v", p.input)
			t.Fatalf("Expected %#v, got %#v", p.expected, result)
		}
	}
}
