package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

import (
	"strconv"
	"errors"
)

// reverse reverses a string of single-byte runes.
func reverse(str string) string {
	strLen := len(str)
	r := make([]byte, strLen)
	for i := 0; i < strLen; i++ {
		r[i] = str[strLen-i-1]
	}
	return string(r)
}

// reverse returns a copy of a slice with the elements in the opposite order.
// This does not work for strings with multi-byte runes
//func reverse(a interface{}) interface{} {
//	val := reflect.ValueOf(a)
//	valLen := val.Len()
//	r := make([]interface{}, valLen)
//	for i := 0; i < valLen; i++ {
//		r[i] = val.Index(valLen - i - 1).Interface()
//	}
//	return r
//}

// isPrime returns true if n is a prime number.
func isPrime(n uint) bool {
	switch n {
	case 0, 1:
		return false
	default:
		for i := uint(2); i < n; i++ {
			// i is a factor if the remainder is 0
			if n % i == 0 {
				return false
			}
		}
	}
	return true
}

// isPalindrome returns true if a string is the same forwards and backwards
func isPalindrome(str string) bool {
	rev := reverse(str)
	return str == rev
}

// greatestPrimePalindrome returns the greatest prime palindrome less than n
func greatestPrimePalindrome(n uint) (uint, error) {
	for i := n - 1; i > 1; i-- {
		str := strconv.Itoa(int(i))
		if isPalindrome(str) && isPrime(i) {
			return i, nil
		}
	}
	return 0, errors.New(fmt.Sprint("No prime palindrome < ", n))
}

func gpp1000([]byte) (interface{}, error) {
	return greatestPrimePalindrome(1000)
}

func stdinToFuncToStdout(fn func([]byte) (interface{}, error)) {
	// Read in all of stdin
	stdin, err := ioutil.ReadAll(os.Stdin)
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
	stdinToFuncToStdout(gpp1000)
}
