package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

import (
	"strconv"
	"strings"
	"log"
)

// Aliasing the type and implementing another interface wasn't a good idea in
// this case, but I wanted to get some practice with it
type int_ int

type Dividend interface {
	IsDivisible(int) bool
}

// IsDivisible returns true if the dividend is evenly divisible by the divisor.
func (dividend int_) IsDivisible(divisor int_) bool {
	return dividend % divisor == 0
}

// Input parameters for each Fizz Buzz
type params struct {
	X, Y, Last int_
}

//
func strToInts(input string) []int {
	argStrs := strings.Split(input, " ")
	args := make([]int, len(argStrs))
	for j, str := range argStrs {
		args[j], _ = strconv.Atoi(str)
	}
	return args
}

// fizzBuzzNum performs Fizz Buzz on one number.
func fizzBuzzNum(p *params, num int_) string {
	str := ""
	if num.IsDivisible(p.X) {
		str += "F"
	}
	if num.IsDivisible(p.Y) {
		str += "B"
	}
	if len(str) == 0 {
		str += strconv.Itoa(int(num))
	}
	return str
}

// fizzBuzzLine performs Fizz Buzz for one line of input
func fizzBuzzLine(line string) string {
	args := strToInts(line)
	p := params{int_(args[0]), int_(args[1]), int_(args[2])}
	numStrs := make([]string, p.Last)

	// For each num
	for i, num := 0, int_(1); num <= p.Last; i, num = i + 1, num + 1 {
		numStrs[i] = fizzBuzzNum(&p, num)
	}
	return strings.Join(numStrs, " ")
}

func fizzBuzz(stdin []byte) (interface{}, error) {
	lines := strings.Split(strings.Trim(string(stdin), "\n"), "\n")
	outs := make([]string, len(lines))
	for i, line := range lines {
		outs[i] = fizzBuzzLine(line)
	}
	rv := strings.Join(outs, "\n")
	return rv, nil
}


func stdinToFuncToStdout(fn func([]byte) (interface{}, error)) {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	// Read in all of stdin
	stdin, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Get call fn with stdin
	rv, err := fn(stdin)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print result to stdout
	fmt.Println(rv)
}

func main() {
	stdinToFuncToStdout(fizzBuzz)
}
