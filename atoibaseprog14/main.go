///////////////////////////////////////////////////////////////////////////////
// atoibaseprog
// allowedFunctions	[…]
// 0	--cast
// 1	len
// WARNING! VERY IMPORTANT!

// For this exercise a function will be tested with the exam own main. However the student still needs to submit a structured program:

// This means that:

//     The package needs to be named package main.
//     The submitted code needs one declared function main(func main()) even if empty.
//     The function main declared needs to also pass the Restrictions Checker(illegal functions tester). It is advised for the student to just empty the function main after its own testings are done.
//     Every other rules are obviously the same than for a program.

// Instructions

// Write a function that takes a string number and its string base in parameters and returns its conversion as an int.

// If the base or the string number is not valid it returns 0

// Validity rules for a base :

//     A base must contain at least 2 characters.
//     Each character of a base must be unique.
//     A base should not contain + or - characters.

// Only valid string numbers will be tested.

// The function does not have to manage negative numbers.
// Expected function

// func AtoiBase(s string, base string) int {

// }

// Usage

// Here is a possible program to test your function :

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	fmt.Println(AtoiBase("125", "0123456789"))
// 	fmt.Println(AtoiBase("1111101", "01"))
// 	fmt.Println(AtoiBase("7D", "0123456789ABCDEF"))
// 	fmt.Println(AtoiBase("uoi", "choumi"))
// 	fmt.Println(AtoiBase("bbbbbab", "-ab"))
// }

// And its output :

// student@ubuntu:~/test$ go build
// student@ubuntu:~/test$ ./test
// 125
// 125
// 125
// 125
// 0
// student@ubuntu:~/test$
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
)

func main() {
	fmt.Println(AtoiBase("125", "0123456789"))
	fmt.Println(AtoiBase("1111101", "01"))
	fmt.Println(AtoiBase("7D", "0123456789ABCDEF"))
	fmt.Println(AtoiBase("uoi", "choumi"))
	fmt.Println(AtoiBase("bbbbbab", "-ab"))
}

func AtoiBase(s string, base string) int {
	if !isValid(base) {
		return 0
	}
	result := 0
	for _, vS := range s {
		for iB, vB := range base {
			if vS == vB {
				result = result*len(base) + iB
			}
		}
	}
	return result
}
func isValid(s string) bool {
	if s[0] == '-' || s[0] == '+' {
		return false
	}
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				return false
			}
		}
	}
	return true
}

// func AtoiBase(s string, base string) int {
// 	result := 0
// 	ArrBase := []rune(base)

// 	if len(ArrBase) < 2 {
// 		return 0
// 	}
// 	for _, i := range base {
// 		count := 0
// 		for _, j := range base {
// 			if i == j {
// 				count++
// 			}
// 			if count > 1 || j == '-' || j == '+' {
// 				return 0
// 			}
// 		}
// 	}
// 	for _, a := range s {
// 		for i := 0; i < len(ArrBase); i++ {
// 			if a == ArrBase[i] {
// 				result = result*len(ArrBase) + i
// 			}
// 		}
// 	}
// 	return result
// }
