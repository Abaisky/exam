////////////////////////////////////////////////////////////////////////////////////////////////

// compareprog
// allowedFunctions	[…]
// 0	"os.Args"
// 1	"github.com/01-edu/z01.PrintRune"
// 2	"len"
// 3	"--cast"
// Instructions

// Write a program that behaves like the Compare function from the Go package strings.

// This program prints a number after comparing two string lexicographically.
// Usage

// student@ubuntu:~/compareprog$ go build
// student@ubuntu:~/compareprog$ ./compareprog a b | cat -e
// -1$
// student@ubuntu:~/compareprog$ ./compareprog a a | cat -e
// 0$
// student@ubuntu:~/compareprog$ ./compareprog b a | cat -e
// 1$
// student@ubuntu:~/compareprog$ ./compareprog b a d | cat -e
// $
// student@ubuntu:~/compareprog$ ./compareprog | cat -e
// $
// student@ubuntu:~/compareprog$
///////////////////////////////////////////////////////////////////////////////////////////////

package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	// fmt.Println(strings.Compare("bb", "ba")) //for real compare

	defer z01.PrintRune('\n')
	zero := "0"
	plus := "1"
	minus := "-1"
	arg := os.Args[1:]
	if len(arg) != 2 {
		return
	}
	if arg[0] == arg[1] {
		printStr(zero)
	} else if arg[0] < arg[1] {
		printStr(minus)
	} else if arg[0] > arg[1] {
		printStr(plus)
	}
}

func printStr(s string) {
	for _, v := range s {
		z01.PrintRune(v)
	}
}
