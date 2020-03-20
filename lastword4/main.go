//////////////////////////////////////////////////////////////////////////////////////////////////////////////
// lastword
// allowedFunctions	[…]
// 0	--cast
// 1	github.com/01-edu/z01.PrintRune
// 2	len
// 3	os.Args
// Instructions
// Write a program that takes a string and displays its last word, followed by a newline ('\n').

//     A word is a section of string delimited by spaces or by the start/end of the string.

//     The output will be followed by a newline ('\n').

//     If the number of parameters is different from 1, or if there are no word, the program displays a newline ('\n').

// Usage
// student@ubuntu:~/[[ROOT]]/lastword$ go build
// student@ubuntu:~/[[ROOT]]/lastword$ ./lastword "FOR PONY" | cat -e
// PONY$
// student@ubuntu:~/[[ROOT]]/lastword$ ./lastword "this        ...       is sparta, then again, maybe    not" | cat -e
// not$
// student@ubuntu:~/[[ROOT]]/lastword$ ./lastword "  " | cat -e
// $
// student@ubuntu:~/[[ROOT]]/lastword$ ./lastword "a" "b" | cat -e
// $
// student@ubuntu:~/[[ROOT]]/lastword$ ./lastword "  lorem,ipsum  " | cat -e
// lorem,ipsum$
// student@ubuntu:~/[[ROOT]]/lastword$
/////////////////////////////////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("error")
		return
	}

	fl1 := false
	newarg := ""
	fl2 := false
	newarg2 := ""

	for _, arg2 := range args {
		for i := len(arg2); i > 0; i-- {
			if arg2[i-1] != ' ' && fl1 == false {

				// z01.PrintRune(rune(arg2[i-1]))
				newarg = arg2[:i]
				fl1 = true
			}
		}
	}

	for i := len(newarg); i > 0; i-- {
		if newarg[i-1] == ' ' && fl2 == false {
			newarg2 = newarg[i:]
			fl2 = true
		}
	}
	if fl2 == true {
		for i := 0; i < len(newarg2); i++ {
			z01.PrintRune(rune(newarg2[i]))

		}
	} else {
		for i := 0; i < len(newarg); i++ {
			z01.PrintRune(rune(newarg[i]))

		}
	}
	fmt.Println()

}
