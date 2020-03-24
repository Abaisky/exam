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
	defer z01.PrintRune('\n')
	arg := os.Args[1:]
	if len(arg) != 1 {
		return
	}
	printStr(rostring(spacesBegin(spacesEnd(arg[0]))))
}

func printStr(s string) {
	for _, v := range s {
		z01.PrintRune(v)
	}
}

func spacesBegin(s string) string {
	for i, _ := range s {
		if s[0] == ' ' {
			if s[i] == ' ' && s[i+1] != ' ' {
				return s[i+1:]
			}
		}
	}
	return s
}

func spacesEnd(s string) string {
	for i := len(s) - 1; i >= 0; i-- {
		if s[len(s)-1] == ' ' {
			if s[i] == ' ' && s[i-1] != ' ' {
				return s[:i]
			}
		}
	}
	return s
}

func rostring(s string) string {
	for i, _ := range s {
		if s[i] == ' ' && s[i+1] != ' ' {
			return s[i+1:] + " " + s[:i]
		}
	}
	return s
}
