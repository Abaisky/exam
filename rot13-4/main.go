/////////////////////////////////////////////////////
// rot13

// "allowedFunctions":[
// 	"--cast",
// 	"github.com/01-edu/z01.PrintRune",
// 	"os.*",
// 	"len",
// 	"--no-lit=[b-mB-M]"
//  ]

// Instructions

// Write a program that takes a string and displays it, replacing each of its letters by the letter 13 spaces ahead in alphabetical order.

//     ‘z’ becomes ‘m’ and ‘Z’ becomes ‘M’. Case remains unaffected.

//     The output will be followed by a newline ('\n').

//     If the number of arguments is different from 1, the program displays a newline ('\n').

// Usage

// student@ubuntu:~/[[ROOT]]/test$ go build
// student@ubuntu:~/[[ROOT]]/test$ ./test "abc"
// nop
// student@ubuntu:~/[[ROOT]]/test$ ./test "hello there"
// uryyb gurer
// student@ubuntu:~/[[ROOT]]/test$ ./test "HELLO, HELP"
// URYYB, URYC
// student@ubuntu:~/[[ROOT]]/test$ ./test

// student@ubuntu:~/[[ROOT]]/test$
////////////////////////////////////////////////////////////

package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	// fmt.Printf("%T %v %q", args, args, args)
	// fmt.Println()
	// fmt.Println(args[0])
	if len(args) == 1 {
		for i := 0; i < len(args[0]); i++ {
			if (args[0][i] >= 'a' && args[0][i] < 'n') || (args[0][i] >= 'A' && args[0][i] < 'N') {
				z01.PrintRune(rune(args[0][i]) + 13)
			} else if (args[0][i] >= 'n' && args[0][i] <= 'z') || (args[0][i] >= 'N' && args[0][i] <= 'Z') {
				z01.PrintRune(rune(args[0][i]) - 13)
			} else {
				z01.PrintRune(rune(args[0][i]))
			}
		}
	}
	z01.PrintRune('\n')

	// for i := 'A'; i <= 'z'; i++ {
	// 	z01.PrintRune(rune(i))
	// }
	// fmt.Println()

	// for i := 'A'; i <= 'z'; i++ {
	// 	fmt.Print(rune(i))
	// }
	// fmt.Println()
}
