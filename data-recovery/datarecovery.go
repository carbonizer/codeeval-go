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
	inputToFuncToStdout(DataRecovery, INPUT_FILEARG)
}

func DataRecovery(input []byte) (interface{}, error) {
	lines := strings.Split(strings.TrimRight(string(input), "\n"), "\n")
	recoveredLines := make([]string, len(lines))
	err := error(nil)

	for i, line := range lines {
		// Semicolon splits mixed words and positions
		semicolonSplits := strings.Split(line, ";")
		words := strings.Split(semicolonSplits[0], " ")
		positionStrs := strings.Split(semicolonSplits[1], " ")

		// Convert positions from strings to ints
		positions := make([]int, len(positionStrs))
		for j, posStr := range positionStrs {
			positions[j], err = strconv.Atoi(posStr)
			if err != nil {
				log.Fatal("Couldn't convert %#v", posStr)
			}
		}

		// pos is the 1-based position where words[j] needs to be moved
		recoveredWords := make([]string, len(words))
		for j, pos := range positions {
			recoveredWords[pos-1] = words[j]
		}

		// There are n words and n - 1 positions, meaning the last
		// mixed word needs to replace the one item in recoveredWords
		// that is blank
		for j, word := range recoveredWords {
			if len(word) == 0 {
				recoveredWords[j] = words[len(words)-1]
				break
			}
		}

		recoveredLines[i] = strings.Join(recoveredWords, " ")
	}

	return strings.Join(recoveredLines, "\n"), nil
}
