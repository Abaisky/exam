///////////////////////////////////////////////////////////////////////////////////////////

// alphamirror
// allowedFunctions	[…]
// 0	"github.com/01-edu/z01.PrintRune"
// 1	"os.Args"
// 2	"len"
// 3	"--cast"
// 4	"--no-lit=[b-mB-Y]"
// Instructions

// Write a program called alphamirror that takes a string as argument and displays this string after replacing each alphabetical character with the opposite alphabetical character.

// The case of the letter stays the same, for example :

// ‘a’ becomes ‘z’, ‘Z’ becomes ‘A’ ‘d’ becomes ‘w’, ‘M’ becomes ‘N’

// The final result will be followed by a newline ('\n').

// If the number of arguments is different from 1, the program displays only a newline ('\n').
// Usage

// student@ubuntu:~/[[ROOT]]/alphamirror$ go build
// student@ubuntu:~/[[ROOT]]/alphamirror$ ./alphamirror "abc"
// zyx
// student@ubuntu:~/[[ROOT]]/alphamirror$ ./alphamirror "My horse is Amazing." | cat -e
// Nb slihv rh Znzarmt.$
// student@ubuntu:~/[[ROOT]]/alphamirror$ ./alphamirror | cat -e
// $
// student@ubuntu:~/[[ROOT]]/alphamirror$
///////////////////////////////////////////////////////////////////////////////////////////////////////////////

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
	} else {
		for _, r := range arg[0] {
			if r >= 'a' && r <= 'z' {
				z01.PrintRune('a' + 'z' - r)
			} else if r >= 'A' && r <= 'Z' {
				z01.PrintRune('A' + 'Z' - r)
			} else {
				z01.PrintRune(r)
			}
		}
	}
}
