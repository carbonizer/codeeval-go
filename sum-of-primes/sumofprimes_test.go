package main

import (
	"fmt"
	"testing"
)

func TestSumOfPrimes(t *testing.T) {
	type Pair struct {
		input    string
		expected int
	}

	pairs := []Pair{{"1000", 3682913}}
	for _, p := range pairs {
		result, err := SumOfPrimes([]byte(p.input))
		if err != nil {
			t.Fatal(err)
		}
		if p.expected != result {
			fmt.Println("Input: %#v", p.input)
			t.Fatalf("Expected %#v, got %#v", p.expected, result)
		}
	}
}
