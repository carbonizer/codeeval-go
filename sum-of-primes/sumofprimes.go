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

const FAKE_INPUT = "1000"

//
// Custom code for this problem
//

// Just update the function name and the input type
func main() {
	inputToFuncToStdout(SumOfPrimes, INPUT_CONSTANT)
}

// isPrime returns true if n is a prime number.
func isPrime(n int) bool {
	switch n {
	case 0, 1:
		return false
	default:
		for i := 2; i < n; i++ {
			// i is a factor if the remainder is 0
			if n%i == 0 {
				return false
			}
		}
	}
	return true
}

func SumOfPrimes(stdin []byte) (interface{}, error) {
	n, err := strconv.Atoi(string(stdin))

	if err != nil {
		fmt.Println()
		return "", fmt.Errorf(
			"Cannot convert %#v to number\n", string(stdin))
	}

	sum := 0
	for i, numPrimes := 0, 0; numPrimes < n; i++ {
		if isPrime(i) {
			sum += i
			numPrimes++
		}
	}

	return sum, nil
}
