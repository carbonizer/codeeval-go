package main

import (
	"testing"
	"fmt"
)


func TestReverse(t *testing.T) {
	type Pair struct {
		input, expected string
	}

	pairs := []Pair{{"hello", "ollehh"}}
	for _, p := range pairs {
		result := reverse(p.input)
		if p.expected == result {
			fmt.Println("Input:", p.input)
			t.Fatalf("Expected %v, got %v", p.expected, result)
		}
	}
}


func TestIsPrime(t *testing.T) {
	type Pair struct {
		input uint
		expected bool
	}

	pairs := []Pair{{0, false}, {1, false}, {2, true}, {3, true},
			{4, false}, {7, true}}
	for _, p := range pairs {
		result := isPrime(p.input)
		if p.expected != result {
			fmt.Println("Input:", p.input)
			t.Fatalf("Expected %v, got %v", p.expected, result)
		}
	}
}

func TestIsPalindrome(t *testing.T) {
	type Pair struct {
		input string
		expected bool
	}

	pairs := []Pair{{"mom", true}, {"hello", false}}
	for _, p := range pairs {
		result := isPalindrome(p.input)
		if p.expected != result {
			fmt.Println("Input:", p.input)
			t.Fatalf("Expected %v, got %v", p.expected, result)
		}
	}
}

func TestGreatestPrimePalindrome(t *testing.T) {
	type Pair struct {
		input, expected uint
	}

	pairs := []Pair{{1e5, 98689}, {10000, 929}, {1000, 929}, {100, 11},
			{10, 7}}
	for _, p := range pairs {
		result, err := greatestPrimePalindrome(p.input)
		if err != nil {
			t.Fatal(err)
		}
		if p.expected != result {
			t.Fatalf("Expected %v, got %v", p.expected, result)
		}
	}
}
