/////////////////////////////////////////////////
// rostring
// allowedFunctions	[…]
// 0	os.Args
// 1	make
// 2	len
// 3	--cast
// 4	github.com/01-edu/z01.PrintRune
// Instructions

// Write a program that takes a string and displays this string after rotating it one word to the left.

// Thus, the first word becomes the last, and others stay in the same order.

// A word is a sequence of alphanumerical characters.

// Words will be separated by only one space in the output.

// If the number of arguments is different from 1, the program displays a newline.
// Usage

// student@ubuntu:~/[[ROOT]]/rostring$ go build
// student@ubuntu:~/[[ROOT]]/rostring$ ./rostring "abc   " | cat -e
// abc$
// student@ubuntu:~/[[ROOT]]/rostring$ ./rostring "Let there     be light"
// there be light Let
// student@ubuntu:~/[[ROOT]]/rostring$ ./rostring "     AkjhZ zLKIJz , 23y"
// zLKIJz , 23y AkjhZ
// student@ubuntu:~/[[ROOT]]/rostring$ ./rostring | cat -e
// $
// student@ubuntu:~/[[ROOT]]/rostring$
////////////////////////////////////////////

package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args[1:]) != 1 {
		z01.PrintRune('\n')
		return
	}
	arg := os.Args[1]
	printMyRune(RotString(string(arg)))
}

func printMyRune(s string) {
	for _, v := range s {
		z01.PrintRune(v)
	}
	z01.PrintRune('\n')
}

func RotString(s string) string {
	s = leaveOneSpace(trim(s))
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			return s[i+1:] + " " + s[:i]
		}
		if i == len(s)-1 {
			return s
		}
	}
	return ""
}

func leaveOneSpace(s string) string {
	res := ""
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' && s[i+1] == ' ' {
			continue
		}
		res += string(s[i])
	}
	return res
}

func trim(s string) string {
	if len(s) == 0 {
		return s
	}
	if s[0] == ' ' {
		return trim(s[1:])
	}
	if s[len(s)-1] == ' ' {
		return trim(s[:len(s)-1])
	}
	return s
}
