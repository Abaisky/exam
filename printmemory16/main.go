///////////////////////////////////////////////////////////////////////////////////////////////////////////
// printmemory
// allowedFunctions	[…]
// 0	"len"
// 1	"--cast"
// 2	"append"
// 3	"github.com/01-edu/z01"
// WARNING! VERY IMPORTANT!

// For this exercise a function will be tested with the exam own main. However the student still needs to submit a structured program:

// This means that:

//     The package needs to be named package main.
//     The submitted code needs one declared function main(func main()) even if empty.
//     The function main declared needs to also pass the Restrictions Checker(illegal functions tester). It is advised for the student to just empty the function main after its own testings are done.
//     Every other rules are obviously the same than for a program.

// Instructions

// Write a function that takes (arr [10]int), and displays the memory as in the example.
// Expected function

// func PrintMemory(arr [10]int) {

// }

// Usage

// Here is a possible program to test your function :

// package main

// func main() {
// 	arr := [10]int{104, 101, 108, 108, 111, 16, 21, 42}
// 	PrintMemory(arr)
// }

// And its output :

// student@ubuntu:~/[[ROOT]]/test$ go build
// student@ubuntu:~/[[ROOT]]/test$ ./test | cat -e
// 6800 0000 6500 0000 6c00 0000 6c00 0000 $
// 6f00 0000 1000 0000 1500 0000 2a00 0000 $
// 0000 0000 0000 0000 $
// hello..*..$
// student@ubuntu:~/[[ROOT]]/test$

///////////////////////////////////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"

	"github.com/01-edu/z01"
)

func main() {
	arr := [10]int{104, 101, 108, 108, 111, 16, 21, 42}
	PrintMemory(arr)
}

func PrintMemory(arr [10]int) {
	for i, nm := range arr {
		hex := nbrBase(nm, "0123456789abcdef")

		fmt.Print(hex)
		n := len(hex)
		for n <= 8 {

			if n == 4 {

				fmt.Print(" ")
			} else {
				fmt.Print("0")
			}

			n++
		}
		fmt.Print(" ")
		if (i+1)%4 == 0 {
			fmt.Println()
		}
	}
	fmt.Println()
	for _, nm := range arr {
		if nm >= 31 {
			z01.PrintRune(rune(nm))
		} else {
			z01.PrintRune('.')
		}
	}
	fmt.Println()
}

func nbrBase(nm int, base string) string {
	result := ""
	isNegative := false
	if nm < 0 {
		isNegative = true
		nm *= -1
	}
	for nm != 0 {
		result = string(base[nm%len(base)]) + result
		nm /= len(base)
	}

	if isNegative {
		result = "-" + result
	}

	return result
}
