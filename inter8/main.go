//////////////////////////////////////////////////////////////////////////////////////////////////////
// inter
// allowedFunctions	[…]
// 0	"--cast"
// 1	"github.com/01-edu/z01.PrintRune"
// 2	"os.Args"
// 3	"len"
// Instructions

// Write a program that takes two string and displays, without doubles, the characters that appear in both string, in the order they appear in the first one.

//     The display will be followed by a newline ('\n').

//     If the number of arguments is different from 2, the program displays a newline ('\n').

// Usage

// student@ubuntu:~/[[ROOT]]/test$ go build
// student@ubuntu:~/[[ROOT]]/test$ ./test "padinton" "paqefwtdjetyiytjneytjoeyjnejeyj"
// padinto
// student@ubuntu:~/[[ROOT]]/test$ ./test ddf6vewg64f  twthgdwthdwfteewhrtag6h4ffdhsd
// df6ewg4
// student@ubuntu:~/[[ROOT]]/test$

//////////////////////////////////////////////////////////////////////////////////////////////////////

package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	defer z01.PrintRune('\n')
	preResult := ""
	result := ""
	arg := os.Args[1:]
	if len(arg) != 2 {
		return
	}
	for _, v := range arg[0] {
		if content(arg[1], v) {
			preResult += string(v)
		}
	}
	for _, v := range preResult {
		if !content(result, v) {
			result += string(v)
		}
	}
	for _, v := range result {
		z01.PrintRune(v)
	}

}

func content(s string, r rune) bool {
	for _, v := range s {
		if v == r {
			return true
		}
	}
	return false
}
