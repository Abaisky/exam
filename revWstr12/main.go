//////////////////////////////////////////////////////////////////////////////////////
// revWstr
// allowedFunctions	[…]
// 0	"os.Args"
// 1	"len"
// 2	"--cast"
// 3	"github.com/01-edu/z01.PrintRune"
// Instructions

// Write a program that takes a string as a parameter, and prints its words in reverse.

//     A word is a sequence of alphanumerical characters.

//     If the number of parameters is different from 1, the program will display newline ('\n').

//     In the parameters that are going to be tested, there will not be any extra spaces. (meaning that there will not be additional spaces at the beginning or at the end of the string and that words will always be separated by exactly one space).

// Usage

// student@ubuntu:~/[[ROOT]]/test$ go build
// student@ubuntu:~/[[ROOT]]/test$ ./test "the time of contempt precedes that of indifference"
// indifference of that precedes contempt of time the
// student@ubuntu:~/[[ROOT]]/test$ ./test "abcdefghijklm"
// abcdefghijklm
// student@ubuntu:~/[[ROOT]]/test$ ./test "he stared at the mountain"
// mountain the at stared he
// student@ubuntu:~/[[ROOT]]/test$ ./test "" | cat-e
// $
// student@ubuntu:~/[[ROOT]]/test$
///////////////////////////////////////////////////////////////////////////////////

package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	defer z01.PrintRune('\n')
	if len(os.Args) != 2 {
		return
	}
	newWord := ""
	oldWord := os.Args[1] + " $"
	index := 0
	for i, _ := range oldWord {
		if oldWord[i] == ' ' && oldWord[i+1] != ' ' {
			newWord = oldWord[index:i] + " " + newWord
			index = i + 1
		}
	}
	for _, v := range newWord {
		z01.PrintRune(v)
	}
}
