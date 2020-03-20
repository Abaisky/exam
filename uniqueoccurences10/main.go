///////////////////////////////////////////////////////////////////////////////
// uniqueoccurences
// allowedFunctions	[…]
// 0	github.com/01-edu/z01.PrintRune
// 1	os.*
// 2	len
// 3	make
// Instructions

// Write a program that outputs true if the number of occurrences of each character is unique, otherwise false. \n should be at the of line.

// If number of arguments is not 1 output \n.

// Only lower case characters will be given.
// Usage

// student@ubuntu:~/[[ROOT]]/test$ go build
// student@ubuntu:~/[[ROOT]]/test$ ./main "abbaac"
// true
// student@ubuntu:~/[[ROOT]]/test$ ./main "ab"
// false
// student@ubuntu:~/[[ROOT]]/test$ ./main "abcacccazb"
// true

// In first example, character ‘a’ has 3 occurrences, ‘b’ has 2 and ‘c’ has 1. No two characters have the same number of occurrences.
////////////////////////////////////////////////////////////////////////////

package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[1:]
	if len(arg) != 1 {
		z01.PrintRune('\n')
		return
	}

	mp := map[rune]int{}

	for _, ar := range arg {
		for _, v := range ar {
			mp[v]++
		}
	}

	for fv, fi := range mp {
		for sv, si := range mp {
			if fv != sv && fi == si {
				fl := "false"
				printMyRune(fl)
				return
			}
		}
	}
	tr := "true"
	printMyRune(tr)

}

func printMyRune(data string) {
	for _, v := range data {
		z01.PrintRune(v)
	}
	z01.PrintRune('\n')
}
