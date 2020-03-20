//////////////////////////////////////////////////////////////////////////////////////////////
// balancedstring
// allowedFunctions	[…]
// 0	fmt.*
// 1	os.*
// 2	len
// Instructions

// Balanced string is a string that has equal quantity of ‘C’ and ‘D’ characters.

// Write a program that takes a string and outputs maximum amount of balanced strings without ignoring any letters. Display output with \n at the end of line.

// If the number of arguments is not 1, display \n.

// It will only be tested strings containing the characters ‘C’ and ‘D’.
// Usage

// student@ubuntu:~/[[ROOT]]/balancedstring$ go build
// student@ubuntu:~/[[ROOT]]/balancedstring$ ./balancedstring "CDCCDDCDCD"
// 4
// student@ubuntu:~/[[ROOT]]/balancedstring$ ./balancedstring "CDDDDCCCDC"
// 3
// student@ubuntu:~/[[ROOT]]/balancedstring$ ./balancedstring "DDDDCCCC"
// 1
// student@ubuntu:~/[[ROOT]]/balancedstring$ ./balancedstring "CDCCCDDCDD"
// 2

// In first example “CDCCDDCDCD” can be split into “CD”, “CCDD”, “CD”, “CD”, each substring contains same number of ‘C’ and ‘D’.

// Second, “CDDDDCCCDC” can be split into: “CD”, “DDDCCC”, “DC”.

// “DDDDCCCC” can be split into “DDDDCCCC”.

// “CDCCCDDCDD” can be split into: “CD”, “CCCDDCDD”, since each substring contains an equal number of ‘C’ and ‘D’
////////////////////////////////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"os"
)

func main() {
	arg := os.Args[1:]
	if len(arg) != 1 {
		fmt.Println()
		return
	}
	counter := 0
	balance := 0
	for _, v := range arg {
		for _, v2 := range v {
			if v2 == 'C' {
				counter++
			}
			if v2 == 'D' {
				counter--
			}
			if counter == 0 {
				balance++
			}
		}
	}
	fmt.Println(balance)
}
