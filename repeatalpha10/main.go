/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// repeatalpha
// allowedFunctions	[…]
// 0	"--cast"
// 1	"github.com/01-edu/z01.PrintRune"
// 2	"os.*"
// 3	"len"
// Instructions

// Write a program called repeat_alpha that takes a string and displays it repeating each alphabetical character as many times as its alphabetical index.

// The result must be followed by a newline ('\n').

// 'a' becomes 'a', 'b' becomes 'bb', 'e' becomes 'eeeee', etc…

// The case remains unchanged.

// If the number of arguments is different from 1, the program displays a newline ('\n').
// Usage

// student@ubuntu:~/[[ROOT]]/repeatalpha$ go build
// student@ubuntu:~/[[ROOT]]/repeatalpha$ ./repeatalpha "abc" | cat -e
// abbccc
// student@ubuntu:~/[[ROOT]]/repeatalpha$ ./repeatalpha "Choumi." | cat -e
// CCChhhhhhhhooooooooooooooouuuuuuuuuuuuuuuuuuuuummmmmmmmmmmmmiiiiiiiii.$
// student@ubuntu:~/[[ROOT]]/repeatalpha$ ./repeatalpha "abacadaba 01!" | cat -e
// abbacccaddddabba 01!$
// student@ubuntu:~/[[ROOT]]/repeatalpha$ ./repeatalpha | cat -e
// $
// student@ubuntu:~/[[ROOT]]/repeatalpha$ ./repeatalpha "" | cat -e
// $
// student@ubuntu:~/[[ROOT]]/repeatalpha$

/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	defer z01.PrintRune('\n')
	result := ""
	arg := os.Args[1:]
	if len(arg) != 1 {
		return
	}
	for _, v := range arg[0] {
		if 'a' <= v && v <= 'z' {
			for x := 0; x <= int(v)-'a'; x++ {
				result += string(v)
			}
		} else if 'A' <= v && v <= 'Z' {
			for x := 0; x <= int(v)-'A'; x++ {
				result += string(v)
			}
		} else {
			result += string(v)
		}
	}
	printStr(result)
}

func printStr(s string) {
	for _, v := range s {
		z01.PrintRune(v)
	}
}
