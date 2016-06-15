package main

import (
	"testing"
)

func TestInterruptedBubbleSort(t *testing.T) {
	type Pair struct {
		input, expected string
	}

	pairs := []Pair{{
		`36 47 78 28 20 79 87 16 8 45 72 69 81 66 60 8 3 86 90 90 | 1
40 69 52 42 24 16 66 | 2
54 46 0 34 15 48 47 53 25 18 50 5 21 76 62 48 74 1 43 74 78 29 | 6
48 51 5 61 18 | 2
59 68 55 31 73 4 1 25 26 19 60 0 | 2`,
		`36 47 28 20 78 79 16 8 45 72 69 81 66 60 8 3 86 87 90 90
40 42 24 16 52 66 69
0 15 25 18 34 5 21 46 47 48 48 1 43 50 53 29 54 62 74 74 76 78
5 48 18 51 61
55 31 59 4 1 25 26 19 60 0 68 73`}}
	for _, p := range pairs {
		result, err := InterruptedBubbleSort([]byte(p.input))
		if err != nil {
			t.Fatal(err)
		}
		if p.expected != result {
			t.Fatalf("Input: %#v\nExpected: %#v\n     Got: %#v\n",
				p.input, p.expected, result)
		}
	}
}
