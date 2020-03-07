/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// firstruneprog

//     WARNING! VERY IMPORTANT!

//     For this exercise a function will be tested with the exam own main. However the student still needs to submit a structured program

// This means that:

//     The package needs to be named package main.
//     The submitted code needs one declared function main(func main()) even if empty.
//     The function main declared needs to also pass the Restrictions Checker(illegal functions tester). It is advised for the student to just empty the function main after its own testings are done.
// 	Every other rules are obviously the same than for a program.

// 	"allowedFunctions":[
// 		"--cast"
// 	 ]

// Instructions

// Write a function that returns the first rune of a string.
// Expected function

// func FirstRune(s string) rune {

// }
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

package main

import "fmt"

func main() {
	str := "Hello"
	nb := FirstRune(str)
	fmt.Printf("%T %v", nb, nb)
	fmt.Println()
}

func FirstRune(s string) rune {
	ln := 0

	my_run := []rune(s)

	for _, x := range my_run {
		x = x
		ln += 1
	}

	if ln > 0 {
		return my_run[0]
	}

	return 0

}
