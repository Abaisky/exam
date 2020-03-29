///////////////////////////////////////////////////////////////////////////////////////////////////////////
// addprimesum
// llowedFunctions	[…]
// 0	"--cast"
// 1	"github.com/01-edu/z01.PrintRune"
// 2	"os.*"
// 3	"len"
// Instructions

// Write a program that takes a positive integer as argument and displays the sum of all prime numbers inferior or equal to it followed by a newline ('\n').

//     If the number of arguments is different from 1, or if the argument is not a positive number, the program displays 0 followed by a newline.

// Usage

// student@ubuntu:~/[[ROOT]]/test$ go build
// student@ubuntu:~/[[ROOT]]/test$ ./test 5
// 10
// student@ubuntu:~/[[ROOT]]/test$ ./test 7
// 17
// student@ubuntu:~/[[ROOT]]/test$ ./test -2
// 0
// student@ubuntu:~/[[ROOT]]/test$ ./test 0
// 0
// student@ubuntu:~/[[ROOT]]/test$ ./test
// 0
// student@ubuntu:~/[[ROOT]]/test$ ./test 5 7
// 0
// student@ubuntu:~/[[ROOT]]/test$
///////////////////////////////////////////////////////////////////////////////////////////////////////////

package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[1:]
	defer z01.PrintRune('\n')

	if len(arg) != 1 {
		z01.PrintRune('0')
		return
	}
	num := myAtoi(arg[0])
	sum := 0
	for x := 2; x <= num; x++ {
		if isPrime(x) {
			sum += x
		}
	}
	// fmt.Println(sum)
	result := ""
	n := 0
	for sum > 0 {

		n = sum % 10
		for _, v := range string(n) {
			result = string(v+'0') + result
		}
		sum = sum / 10

	}
	for _, v := range result {
		z01.PrintRune(v)
	}

}

func isPrime(n int) bool {
	for x := 2; x < n; x++ {
		if n%x == 0 {
			return false
		}
	}
	return true
}
func myAtoi(s string) int {
	n := 0
	for _, v := range s {
		if v >= '0' && v <= '9' {
			n = n*10 + int(v-'0')
		} else {
			return 0
		}
	}
	return n
}
