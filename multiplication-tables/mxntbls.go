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
	INPUT_NONE InputType = iota
	INPUT_STDIN
	INPUT_FILEARG
	INPUT_CONSTANT
)

func inputToFuncToStdout(fn func([]byte) (interface{}, error), it InputType) {
	fp, input, err := (*os.File)(nil), []byte{}, error(nil)

	switch it {
	// Input via stdin
	case INPUT_STDIN:
		fp = os.Stdin
	// Input from the file at the part of the first command argument
	case INPUT_FILEARG:
		fp, err = os.Open(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
	// Fake input with a constant string
	case INPUT_CONSTANT:
		//log.Println("Fake input")
		input = []byte(FAKE_INPUT)
	// No Input
	default:
		log.Println("No input")
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
	rv, err := fn(input)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print result to stdout
	fmt.Println(rv)
}

const FAKE_INPUT = "12"

//
// Custom code for this problem
//

// Just update the function name and the input type
func main() {
	inputToFuncToStdout(MultiplicationTables, INPUT_CONSTANT)
}

func MultiplicationTables(stdin []byte) (interface{}, error) {
	n, err := strconv.Atoi(string(stdin))

	if err != nil {
		fmt.Println()
		return "", fmt.Errorf(
			"Cannot convert %#v to number\n", string(stdin))
	}

	lines := make([]string, n)
	chunks := make([]string, n)
	for i, _ := range lines {
		for j, _ := range chunks {
			chunks[j] = fmt.Sprintf("%4d", (i+1)*(j+1))
		}
		// Note that because of trimming off the front, rows starting
		// with a two digit number will be shifted off.  However, that
		// is what the description of the problem said should be done,
		// so I did it.
		lines[i] = strings.TrimLeft(strings.Join(chunks, ""), " ")
	}

	return strings.Join(lines, "\n"), nil
}
