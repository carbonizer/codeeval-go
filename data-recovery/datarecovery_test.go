package main

import (
	"testing"
)

func TestReverseWords(t *testing.T) {
	type Pair struct {
		input, expected string
	}

	pairs := []Pair{{
		`2000 and was not However, implemented 1998 it until;9 8 3 4 1 5 7 2
programming first The language;3 2 1
programs Manchester The written ran Mark 1952 1 in Autocode from;6 2 1 7 5 3 11 4 8 9`,
		`However, it was not implemented until 1998 and 2000
The first programming language
The Manchester Mark 1 ran programs written in Autocode from 1952`}}
	for _, p := range pairs {
		result, err := DataRecovery([]byte(p.input))
		if err != nil {
			t.Fatal(err)
		}
		if p.expected != result {
			t.Fatalf("Input: %#v\nExpected: %#v\n     Got: %#v\n",
				p.input, p.expected, result)
		}
	}
}
