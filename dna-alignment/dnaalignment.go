package main

// Common imports for main and handling input
import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// Imports for the specific problem
import (
	"errors"
	"golang.org/x/tools/container/intsets"
	"strconv"
	"strings"
)

//
// Common logic for handling imports
//

type InputType int

const (
	// No Input
	INPUT_NONE InputType = iota
	// Input via stdin.  This is common on hackerrank.com.
	INPUT_STDIN
	// Input from the file at the path of the first command argument.  This
	// is common on codeeval.com.
	INPUT_FILEARG
	// Fake input with a constant string.  This can also be used is a
	// problem indicates using a const value, but you want to write you
	// function more generically
	INPUT_CONSTANT
)

// inputToFuncToStdout wraps a custom function to simplify input and output.
// Various websites provide programming challenges to practice using multiple
// languages and techniques.  Some of the sites (such as codewars.com) test
// submissions using language-specific unit tests.  Others, input data using
// methods common to all languages, and test the output of the program.
// Handling the output is usually as simple as printing to stdout.  However,
// the input is often more complicated, and sometime the code to handle the
// input is not provide.  Another issue is that if you develop the solution
// offline, you may want to use a different method of input while debugging.
//
// This function handles the input and output so you can focus on writing the
// custom code for the specific problem.  Pass in a function that matches the
// signature, and indicate the type of input (in the case of INPUT_CONSTANT,
// update FAKE_INPUT accordingly).  The input will be passed to the function as
// a slice of bytes.  If the function runs successfully, the "%v" form of the
// return value with be printed to stdout.
func inputToFuncToStdout(fn func([]byte) (interface{}, error), it InputType) {
	fp, input, err := (*os.File)(nil), []byte{}, error(nil)

	switch it {
	case INPUT_STDIN:
		fp = os.Stdin

	case INPUT_FILEARG:
		fp, err = os.Open(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}

	case INPUT_CONSTANT:
		input = []byte(FAKE_INPUT)

	// No Input
	default:
		//log.Println("No input")
	}

	if fp != nil {
		// Read all of input from file pointer
		input, err = ioutil.ReadAll(fp)
		if err != nil {
			fp.Close()
			log.Fatal(err)
		}
		fp.Close()
	}

	// Call fn with input
	defer func() {
		if r := recover(); r != nil {
			log.Fatalf("Panic while running fn\n"+
				"Input: %#v\n"+
				"Error: %v\n", string(input), r)
		}
	}()
	rv, err := fn(input)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print result to stdout
	fmt.Println(rv)
}

//
// Custom code for this problem
//

// Update if using INPUT_CONSTANT
const FAKE_INPUT = ""

// Just update the function name and the input type
func main() {
	inputToFuncToStdout(DnaAlignment, INPUT_FILEARG)
}

// StrsToInts converts a slice of strings to a slice of ints.
func StrsToInts(strs []string) ([]int, error) {
	ints := make([]int, len(strs))
	for i, str := range strs {
		int_, err := strconv.Atoi(str)
		if err != nil {
			return ints, err
		}
		ints[i] = int_
	}
	return ints, nil
}

// IntsToStrs converts a slice of ints to a slice of strings.
func IntsToStrs(ints []int) []string {
	strs := make([]string, len(ints))
	for i, int_ := range ints {
		str := strconv.Itoa(int_)
		strs[i] = str
	}
	return strs
}

// UIntsToStrs converts a slice of ints to a slice of strings.
func UIntsToStrs(uints []uint) []string {
	strs := make([]string, len(uints))
	for i, uint_ := range uints {
		str := strconv.FormatUint(uint64(uint_), 10)
		strs[i] = str
	}
	return strs
}

// Factorial calculates the product of (n, n-1, n-2, ..., 3, 2).
func Factorial(n uint) uint {
	var rv uint
	for rv = 1; n > 1; n-- {
		rv *= n
	}
	return rv
}

// Combination calculates the number of combos when choosing k elems from n.
// There is NO replacement, and the order of the elements doesn't matter (the
// same elements in a different order is considered the combination
func Combination(n, k uint) uint {
	return Factorial(n) / (Factorial(k) * Factorial(n-k))
}

// CombinationReplacement calculates combination with replacement.
func CombinationReplacement(n, k uint) uint {
	return Combination(n+k-1, k)
}

// Permutation calculates the number of ways k elems can be selected from n.
// There is NO replacement, but order is important (the same elements in a
// different order is a different permutation
func Permutation(n, k uint) uint {
	return Factorial(n) / (Factorial(n - k))
}

// MakeRange returns such that `CountingNumbers[s:e] == MakeRange(s:e)`.
// CountingNumbers are {0, 1, 2, ...}
func MakeRange(start, end uint) []uint {
	rvLen := end - start
	if rvLen < 1 {
		return []uint{}
	}

	rv := make([]uint, rvLen)
	for i, _ := range rv {
		rv[i] = start + uint(i)
	}

	return rv
}

// IndexCombinationsReplacement returns all combinations of indices.
// k elements are selected from {0, 1, ..., n-1}.  Theses can then be used to
// to create combinations of other slices.
func IndexCombinations(n, k uint) [][]uint {
	combos := make([][]uint, Combination(n, k))
	switch k {
	case 0:
		// Nothing.  Init value will work
	case 1:
		// [[0] [1] ... [n-1]]
		for i, _ := range combos {
			combos[i] = []uint{uint(i)}
		}
	default:
		// Get the preliminary k-1 element groups.  Because of the way
		// the combos are constructed, they are already sorted. Also,
		// because there is NO replacement, the first k-1 elements
		// can never include the greatest index.  As such, the first
		// argument to the recursive call is n-1
		missingLasts := IndexCombinations(n-1, k-1)

		// For each k-1 group, append possible k'th values, each
		// creating a separate, new combo
		i := 0
		for _, missingLast := range missingLasts {
			// Since there is NO replacement, the last element must
			// have a value greater than the (k-1)'th element
			start := missingLast[len(missingLast)-1] + 1

			for _, last := range MakeRange(start, n) {
				combos[i] = append(missingLast, last)
				// Inner iteration varies in size so we need to
				// increment the index manually
				i++
			}
		}
	}
	return combos
}

// IndexCombinationsReplacement returns all combos with replacement of indices.
func IndexCombinationsReplacement(n, k uint) [][]uint {
	combos := make([][]uint, CombinationReplacement(n, k))
	switch k {
	case 0:
		// Nothing.  Init value will work
	case 1:
		// [[0] [1] ... [n-1]]
		for i, _ := range combos {
			combos[i] = []uint{uint(i)}
		}
	default:
		// Get the preliminary k-1 element groups
		missingLasts := IndexCombinationsReplacement(n, k-1)

		// For each k-1 group, append possible k'th values, each
		// creating a separate, new combo
		i := 0
		for _, missingLast := range missingLasts {
			// Due to the nature of how these combos are built up,
			// they are already sorted.  As such, the last value
			// must be greater than or equal the last element.
			// Equal to is included because replacement is allowed.
			start := missingLast[len(missingLast)-1]

			for _, last := range MakeRange(start, n) {
				combos[i] = append(missingLast, last)
				// Inner iteration varies in size so we need to
				// increment the index manually
				i++
			}
		}
	}
	return combos
}

// Concat concatenates each slice of a slice of slices into a single slice.
func Concat(a interface{}) (interface{}, error) {
	switch outer := a.(type) {
	case [][]int:
		outs := []int{}
		for _, inner := range outer {
			outs = append(outs, inner...)
		}
		return outs, nil
	case [][]uint:
		outs := []uint{}
		for _, inner := range outer {
			outs = append(outs, inner...)
		}
		return outs, nil
	case [][]byte:
		outs := []byte{}
		for _, inner := range outer {
			outs = append(outs, inner...)
		}
		return outs, nil
	case [][]rune:
		outs := []rune{}
		for _, inner := range outer {
			outs = append(outs, inner...)
		}
		return outs, nil
	case [][]string:
		outs := []string{}
		for _, inner := range outer {
			outs = append(outs, inner...)
		}
		return outs, nil
	default:
		return nil, errors.New("Input type must be one of: " +
			"[][]int, [][]uint, [][]byte, [][]rune, [][]string")
	}
}

// How to score attempt by comparing corresponding runes
var Score = struct {
	match, mismatch, indelStart, indelExt int
}{
	match:      3,
	mismatch:   -3,
	indelStart: -8,
	indelExt:   -1,
}

func ScoreAttempt(full, a []rune) int {
	var total int
	isStart := true
	for i, c := range a {
		if c == '-' {
			if isStart {
				total += Score.indelStart
				isStart = false
			} else {
				total += Score.indelExt
			}
		} else {
			isStart = true
			if c == full[i] {
				total += Score.match
			} else {
				total += Score.mismatch
			}
		}
	}
	return total
}

func DnaAlignment(input []byte) (interface{}, error) {
	lines := strings.Split(strings.TrimRight(string(input), "\n"), "\n")
	outLines := make([]string, len(lines))

	type Attempt struct {
		seq   []rune
		score int
	}

	for i, line := range lines {
		// " | " splits full and partial sequence
		argStrs := strings.Split(line, " | ")
		full := []rune(argStrs[0])
		partial := []rune(argStrs[1])

		// Determine possible combination of indices for where to stick
		// the inner runes of partial.  Inner means we are ignoring the
		// first and last runes which are anchored, hence the use of -2
		// twice. Note that because we are ignoring the first rune,
		// these indices will be shifted from what they need to be.
		// This is compensated for later, when they are used.
		replaceCombos := IndexCombinations(
			uint(len(full)-2), uint(len(partial)-2))

		possibles := make([]Attempt, len(replaceCombos))
		best := Attempt{[]rune{}, intsets.MinInt}
		for j, combo := range replaceCombos {
			// Prep output with default values
			out := make([]rune, len(full))
			for k, _ := range out {
				switch k {
				// Set the first rune to the
				// first rune of partial
				case 0:
					out[k] = partial[0]

				// Set the last rune to the
				// last rune of partial
				case len(out) - 1:
					out[k] = partial[len(partial)-1]

				// Set all other runes to gaps
				default:
					out[k] = '-'
				}
			}

			// Override some of the default gaps with the inner
			// runes of partial according to the combination of
			// indices.
			for k, letter := range partial[1 : len(partial)-1] {
				// Since we are skipping the first rune, the
				// indices from combo must be shifted up by one
				out[combo[k]+1] = letter
			}

			// Score possibility and replace best if better
			poss := Attempt{out, ScoreAttempt(full, out)}
			if poss.score > best.score {
				best = poss
			}
			possibles[j] = poss
		}

		outLines[i] = strconv.Itoa(best.score)
	}

	return strings.Join(outLines, "\n"), nil
}
