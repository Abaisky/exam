/////////////////////////////////////////////////////////////////////////////////////////////////////
// union
// allowedFunctions	[…]
// 0	"len"
// 1	"os.Args"
// 2	"github.com/01-edu/z01.PrintRune"
// 3	"append"
// 4	"--cast"
// Instructions

// Write a program that takes two string and displays, without doubles, the characters that appear in either one of the string.

// The display will be in the order that the characters will appear on the command line and will be followed by a newline ('\n').

// If the number of arguments is different from 2, then the program displays newline ('\n').
// Usage

// student@ubuntu:~/[[ROOT]]/union$ go build
// student@ubuntu:~/[[ROOT]]/union$ ./union zpadinton paqefwtdjetyiytjneytjoeyjnejeyj | cat -e
// zpadintoqefwjy$
// student@ubuntu:~/[[ROOT]]/union$ ./union ddf6vewg64f gtwthgdwthdwfteewhrtag6h4ffdhsd | cat -e
// df6vewg4thras$
// student@ubuntu:~/[[ROOT]]/union$ ./union rien "cette phrase ne cache rien" | cat -e
// rienct phas$
// student@ubuntu:~/[[ROOT]]/union$ ./union | cat -e
// $
// student@ubuntu:~/[[ROOT]]/union$ ./union rien | cat -e
// $
// student@ubuntu:~/[[ROOT]]/union$
/////////////////////////////////////////////////////////////////////////////////////////////////////

package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[1:]
	defer z01.PrintRune('\n')
	if len(arg) != 2 {
		return
	}
	newSentence := ""
	for _, v := range arg[0] + arg[1] {
		if !contain(newSentence, v) {
			newSentence += string(v)
		}
	}
	for _, v := range newSentence {
		z01.PrintRune(v)
	}
}
func contain(newSentence string, v rune) bool {
	for _, nv := range newSentence {
		if nv == v {
			return true
		}
	}
	return false
}
