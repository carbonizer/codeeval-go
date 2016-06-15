package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestFactorial(t *testing.T) {
	type Pair struct {
		input, expected uint
	}

	pairs := []Pair{{0, 1}, {1, 1}, {2, 2}, {3, 6}, {7, 5040}}
	for _, p := range pairs {
		result := Factorial(p.input)
		if p.expected != result {
			t.Fatalf("Input: %#v\nExpected: %#v\n     Got: %#v\n",
				p.input, p.expected, result)
		}
	}
}

func TestMakeRange(t *testing.T) {
	type Args struct {
		start, end uint
	}
	type Pair struct {
		input    Args
		expected []uint
	}

	pairs := []Pair{
		{Args{0, 1}, []uint{0}},
		{Args{0, 5}, []uint{0, 1, 2, 3, 4}},
		{Args{1, 5}, []uint{1, 2, 3, 4}},
	}
	for _, p := range pairs {
		result := MakeRange(p.input.start, p.input.end)
		if fmt.Sprintf("%#v", p.expected) !=
			fmt.Sprintf("%#v", result) {
			t.Fatalf("Input: %#v\nExpected: %#v\n     Got: %#v\n",
				p.input, p.expected, result)
		}
	}
}

func TestIndexCombinations(t *testing.T) {
	type Args struct {
		n, k uint
	}
	type Pair struct {
		input    Args
		expected [][]uint
	}

	pairs := []Pair{
		{Args{3, 1},
			[][]uint{
				{0},
				{1},
				{2},
			},
		},
		{Args{3, 2},
			[][]uint{
				{0, 1},
				{0, 2},
				{1, 2},
			},
		},
	}
	for _, p := range pairs {
		result := IndexCombinations(p.input.n, p.input.k)
		if fmt.Sprintf("%#v", p.expected) !=
			fmt.Sprintf("%#v", result) {
			t.Fatalf("Input: %#v\nExpected: %#v\n     Got: %#v\n",
				p.input, p.expected, result)
		}
	}
}

func TestIndexCombinations2(t *testing.T) {
	type Args struct {
		n, k uint
	}

	type Params struct {
		args     Args
		expected string
	}

	// General form of the Python code to generate a strings to test against
	// "".join("".join(str(n) for n in ns)
	//     for ns in itertools.combinations(range(n), k))
	pairs := []Params{
		{Args{5, 3}, "012013014023024034123124134234"},
	}
	for _, p := range pairs {
		rv := IndexCombinations(p.args.n, p.args.k)
		rv2, _ := Concat(rv)
		result := strings.Join(UIntsToStrs(rv2.([]uint)), "")
		if fmt.Sprintf("%#v", p.expected) !=
			fmt.Sprintf("%#v", result) {
			t.Fatalf("Input: %#v\nExpected: %#v\n     Got: %#v\n",
				p.args, p.expected, result)
		}
	}
}

func TestIndexCombinationsReplacement(t *testing.T) {
	type Args struct {
		n, k uint
	}

	type Params struct {
		args           Args
		expected       [][]uint
		expectedString string
	}

	// General form of the Python code to generate a strings to test against
	// "".join("".join(str(n) for n in ns)
	//     for ns in itertools.combinations_with_replacement(range(n), k))
	pairs := []Params{
		{args: Args{3, 2}, expected: [][]uint{
			{0, 0}, {0, 1}, {0, 2}, {1, 1}, {1, 2}, {2, 2},
		}},
		{Args{5, 3}, nil,
			"0000010020030040110120130140220230240330340441111121" +
				"13114122123124133134144222223224233234244333334344444",
		},
	}
	for _, p := range pairs {
		rv := IndexCombinationsReplacement(p.args.n, p.args.k)
		expected := ""
		result := ""
		if p.expected != nil {
			expected = fmt.Sprintf("%#v", p.expected)
			result = fmt.Sprintf("%#v", rv)
		} else if p.expectedString != "" {
			expected = fmt.Sprintf("%#v", p.expectedString)
			rv2, _ := Concat(rv)
			result = fmt.Sprintf("%#v", strings.Join(
				UIntsToStrs(rv2.([]uint)), ""))
		}

		if expected != result {
			t.Fatalf("Input: %#v\nExpected: %#v\n     Got: %#v\n",
				p.args, expected, result)
		}
	}
}

func TestScoreAttempt(t *testing.T) {
	type Args struct {
		full, a []rune
	}
	type Pair struct {
		input    Args
		expected int
	}

	pairs := []Pair{
		{Args{[]rune("GAAAAAAT"), []rune("G--A-A-T")}, -13},
		{Args{[]rune("GAAAAAAT"), []rune("GAA----T")}, 1},
	}
	for _, p := range pairs {
		result := ScoreAttempt(p.input.full, p.input.a)
		if fmt.Sprintf("%#v", p.expected) !=
			fmt.Sprintf("%#v", result) {
			t.Fatalf("Input: %#v\nExpected: %#v\n     Got: %#v\n",
				p.input, p.expected, result)
		}
	}
}

func TestDnaAlignment(t *testing.T) {
	type Pair struct {
		input, expected string
	}

	pairs := []Pair{{
		`GAAAAAAT | GAAT
GCATGCT | GATTACA`,
		`1
-3`}}
	for _, p := range pairs {
		result, err := DnaAlignment([]byte(p.input))
		if err != nil {
			t.Fatal(err)
		}
		if p.expected != result {
			t.Fatalf("Input: %#v\nExpected: %#v\n     Got: %#v\n",
				p.input, p.expected, result)
		}
	}
}
