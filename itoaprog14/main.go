//////////////////////////////////////////////////////////////////////
// itoaprog
// allowedFunctions	[…]
// 0	len
// 1	make
// 2	--cast
// ##WARNING! VERY IMPORTANT!

// For this exercise a function will be tested with the exam own main. However the student still needs to submit a structured program:

// This means that:

//     The package needs to be named package main.
//     The submitted code needs one declared function main(func main()) even if empty.
//     The function main declared needs to also pass the Restrictions Checker(illegal functions tester). It is advised for the student to just empty the function main after its own testings are done.
//     Every other rules are obviously the same than for a program.

// Instructions

//     Write a function that simulates the behaviour of the Itoa function in Go. Itoa transforms a number represented as an int in a number represented as a string.

//     For this exercise the handling of the signs + or - does have to be taken into account.

// Expected function

// func Itoa(n int) string {

// }
///////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
)

func main() {
	x := -9223372036854775808
	y := Itoa(x)
	fmt.Println(y)
}

func Itoa(n int) string {
	mark := n
	str := ""
	for n != 0 {
		num := n % 10
		if num < 0 {
			num *= -1
		}
		str = string(num+'0') + str
		n = n / 10
	}
	if mark < 0 {
		return "-" + str
	}
	return str
}

// func Itoa(n int) string {
// 	nbr := n
// 	result := ""
// 	for n != 0 {
// 		index := n % 10
// 		fmt.Println(index)
// 		if index < 0 {

// 			index *= -1
// 		}
// 		result = string(index+48) + result
// 		n /= 10
// 	}
// 	if nbr < 0 {
// 		result = "-" + result
// 	}
// 	return result
// }
