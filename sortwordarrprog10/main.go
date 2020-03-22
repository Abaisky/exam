/////////////////////////////////////////////////////////////////////////////////////
// sortwordarrprog

// allowedFunctions	[…]
// 0	"--cast"
// 1	"make"

// ##WARNING! VERY IMPORTANT!

// For this exercise a function will be tested with the exam own main. However the student still needs to submit a structured program:

// This means that:

//     The package needs to be named package main.
//     The submitted code needs one declared function main(func main()) even if empty.
//     The function main declared needs to also pass the Restrictions Checker(illegal functions tester). It is advised for the student to just empty the function main after its own testings are done.
//     Every other rules are obviously the same than for a program.

// Instructions

// Write a function SortWordArr that sorts by ascii (in ascending order) a string array.
// Expected function

// func SortWordArr(array []string) {

// }

// Usage

// Here is a possible program to test your function :

// package main

// import (
// 	"fmt"
// )

// func main() {

// 	result := []string{"a", "A", "1", "b", "B", "2", "c", "C", "3"}
// 	SortWordArr(result)

// 	fmt.Println(result)
// }

// And its output :

// student@ubuntu:~/test$ go build
// student@ubuntu:~/test$ ./test
// [1 2 3 A B C a b c]
// student@ubuntu:~/test$

/////////////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
)

func main() {

	result := []string{"a", "A", "1", "b", "B", "2", "c", "C", "3"}
	SortWordArr(result)

	fmt.Println(result)
}

func SortWordArr(array []string) {
	l := 0
	for range array {
		l++
	}
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			if array[i] > array[j] {
				array[i], array[j] = array[j], array[i]
			}
		}
	}
}
