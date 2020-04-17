////////////////////////////////////////////////////////////////////////////////////////////////
// lcm
// WARNING! VERY IMPORTANT!

// For this exercise a function will be tested with the exam own main. However the student still needs to submit a structured program:

// This means that:

//     The package needs to be named package main.
//     The submitted code needs one declared function main(func main()) even if empty.
//     The function main declared needs to also pass the Restrictions Checker(illegal functions tester). It is advised for the student to just empty the function main after its own testings are done.
//     Every other rules are obviously the same than for a program.

// Instructions

// Write a function, lcm, that returns least common multiple.

// It will be tested with positive int values and 0.
// Expected function

// func Lcm(first, second int) int {

// }

// Usage

// Here is a possible program to test your function :

// package main

// func main() {
// 	fmt.Println(Lcm(2, 7))
// 	fmt.Println(Lcm(0, 4))
// }

// Output

// student@ubuntu:~/[[ROOT]]/test$ go build
// student@ubuntu:~/[[ROOT]]/test$ ./test
// 14
// 0
// student@ubuntu:~/[[ROOT]]/test$
///////////////////////////////////////////////////////////////////////////////////////////

package main

import "fmt"

func main() {
	fmt.Println(Lcm(2, 7))
	fmt.Println(Lcm(0, 4))
}

func Lcm(first, second int) int {
	if first*second == 0 {
		return 0
	} else {
		for x := 2; x <= first*second; x++ {
			if x%first == 0 && x%second == 0 {
				return x
			}
		}
	}
	return 0
}
