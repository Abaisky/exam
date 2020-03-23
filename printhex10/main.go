//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// printhex
// allowedFunctions	[…]
// 0	"len"
// 1	"--cast"
// 2	"github.com/01-edu/z01.PrintRune"
// 3	"os.*"
// 4	"strconv.*"
// Instructions

// Write a program that takes a positive (or zero) number expressed in base 10, and displays it in base 16 (with lowercase letters) followed by a newline ('\n').

//     If the number of parameters is different from 1, the program displays a newline.
//     Error cases have to be handled as shown in the example below.

// Usage

// student@ubuntu:~/[[ROOT]]/test$ go build
// student@ubuntu:~/[[ROOT]]/test$ ./test "10"
// a
// student@ubuntu:~/[[ROOT]]/test$ ./test "255"
// ff
// student@ubuntu:~/[[ROOT]]/test$ ./test "5156454"
// 4eae66
// student@ubuntu:~/[[ROOT]]/test$ ./test

// student@ubuntu:~/[[ROOT]]/test$ ./test "123 132 1" | cat -e
// 0$
// student@ubuntu:~/[[ROOT]]/test$
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////

package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	defer z01.PrintRune('\n')
	arg := os.Args[1:]
	if len(arg) != 1 {
		return
	}
	num, err := strconv.Atoi(arg[0])
	if err != nil {
		z01.PrintRune('0')
		return
	}
	base := "0123456789abcdef"
	result := ""

	if num == 0 {
		result = "0"
	} else {
		for num != 0 {
			result = string(base[num%16]) + result
			num = num / 16
		}
	}

	for _, v := range result {
		z01.PrintRune(v)
	}
}
