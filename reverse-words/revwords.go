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

const FAKE_INPUT = "Hello World"

//
// Custom code for this problem
//

// Just update the function name and the input type
func main() {
	inputToFuncToStdout(ReverseWords, INPUT_FILEARG)
}

// revStrings returns a slice of strings in the reverse order.
func revStrings(strs []string) []string {
	strsLen := len(strs)
	revs := make([]string, strsLen)
	for i, str := range strs {
		revs[strsLen-1-i] = str
	}
	return revs
}

func ReverseWords(stdin []byte) (interface{}, error) {
	lines := strings.Split(string(stdin), "\n")
	revLines := make([]string, len(lines))
	for i, line := range lines {
		revLines[i] = strings.Join(revStrings(
			strings.Split(line, " ")), " ")
	}
	return strings.Join(revLines, "\n"), nil
}
