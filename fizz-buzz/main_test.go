package main

import (
	"testing"
	"fmt"
	"strings"
)


func TestFizzBuzzLine(t *testing.T) {
	type Pair struct {
		input, expected string
	}

	pairs := []Pair{
		{"3 5 10", "1 2 F 4 B F 7 8 F B"},
		{"2 7 15", "1 F 3 F 5 F B F 9 F 11 F 13 FB 15"},
	}
	for _, p := range pairs {
		result := fizzBuzzLine(p.input)
		if p.expected != result {
			fmt.Println("Input:", p.input)
			t.Fatalf("Expected %#v, got %#v", p.expected, result)
		}
	}
}

func TestFizzBuzz(t *testing.T) {
	type Pair struct {
		input, expected string
	}

	pairs := []Pair{{"3 5 10\n2 7 15\n", strings.Join([]string{
		"1 2 F 4 B F 7 8 F B",
		"1 F 3 F 5 F B F 9 F 11 F 13 FB 15",
	}, "\n")}}
	for _, p := range pairs {
		result, err := fizzBuzz([]byte(p.input))
		if err != nil {
			t.Fatal(err)
		}
		if p.expected != result {
			t.Fatalf("Expected %v, got %v", p.expected, result)
		}
	}
}

