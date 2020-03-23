///////////////////////////////////////////////////////////////////////////////////
// gcd
// allowedFunctions	[…]
// 0	"--cast"
// 1	"fmt.*"
// 2	"os.Args"
// 3	"len"
// 4	"strconv.Atoi"
// Instructions

// Write a program that takes two string representing two strictly positive integers that fit in an int.

// The program displays their greatest common divisor followed by a newline ('\n').

// If the number of parameters is different from 2, the program displays a newline.

// All arguments tested will be positive int values.
// Usage

// student@ubuntu:~/[[ROOT]]/gcd$ go build
// student@ubuntu:~/[[ROOT]]/gcd$ ./gcd 42 10 | cat -e
// 2$
// student@ubuntu:~/[[ROOT]]/gcd$ ./gcd 42 12 | cat -e
// 6$
// student@ubuntu:~/[[ROOT]]/gcd$ ./gcd 14 77 | cat -e
// 7$
// student@ubuntu:~/[[ROOT]]/gcd$ ./gcd 17 3 | cat -e
// 1$
// student@ubuntu:~/[[ROOT]]/gcd$ ./gcd | cat -e
// $
// student@ubuntu:~/[[ROOT]]/gcd$ ./gcd 50 12 4 | cat -e
// $
// student@ubuntu:~/[[ROOT]]/gcd$
///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	defer z01.PrintRune('\n')
	arg := os.Args[1:]
	if len(arg) != 2 {
		return
	}
	n1, _ := strconv.Atoi(arg[0])
	n2, _ := strconv.Atoi(arg[1])

	result := gcd(n1, n2)
	fmt.Print(result)
}
func gcd(n1, n2 int) int {
	for n2 != 0 {
		n1, n2 = n2, n1%n2
	}
	return n1
}
