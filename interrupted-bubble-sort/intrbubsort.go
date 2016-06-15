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
	inputToFuncToStdout(InterruptedBubbleSort, INPUT_FILEARG)
}

// strsToInts converts a slice of strings to a slice of ints
func strsToInts(strs []string) ([]int, error) {
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

// intsToStrs converts a slice of ints to a slice of strings
func intsToStrs(ints []int) []string {
	strs := make([]string, len(ints))
	for i, int_ := range ints {
		str := strconv.Itoa(int_)
		strs[i] = str
	}
	return strs
}

// bubbleSortMax applies a bubble sort in place, limited to max iterations.
// It returns the actual number of iterations used.
func bubbleSortMax(nums []int, max int) int {
	if len(nums) < 2 {
		return 0
	}

	count := 0
	for ; count < max; count++ {
		isChanged := false
		// i is the index of first and i+1 is the index of second
		for i, second := range nums[1:] {
			if nums[i] > second {
				nums[i], nums[i+1] = second, nums[i]
				isChanged = true
			}
		}
		// Already sorted
		if !isChanged {
			// Since we are breaking, count this loop manually
			count++
			break
		}
	}
	return count
}

func InterruptedBubbleSort(input []byte) (interface{}, error) {
	lines := strings.Split(strings.TrimRight(string(input), "\n"), "\n")
	outLines := make([]string, len(lines))
	for i, line := range lines {
		// " | " splits list and number of sorts
		argStrs := strings.Split(line, " | ")

		nums, err := strsToInts(strings.Split(argStrs[0], " "))
		if err != nil {
			return outLines, err
		}

		max, err := strconv.Atoi(argStrs[1])
		if err != nil {
			return outLines, err
		}

		bubbleSortMax(nums, max)

		outLines[i] = strings.Join(intsToStrs(nums), " ")
	}

	return strings.Join(outLines, "\n"), nil
}
