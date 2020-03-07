//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// strlenprog
// WARNING! VERY IMPORTANT!

// For this exercise a function will be tested with the exam own main. However the student still needs to submit a structured program:

// This means that:

//     The package needs to be named package main.
//     The submitted code needs one declared function main(func main()) even if empty.
//     The function main declared needs to also pass the Restrictions Checker(illegal functions tester). It is advised for the student to just empty the function main after its own testings are done.
// 	Every other rules are obviously the same than for a program.

// 	"allowedFunctions":[
// 		"os.Args",
// 		"--cast"
// 	 ]

// Instructions

//     Write a function that counts the runes of a string and that returns that count.

// Expected function

// func StrLen(str string) int {

// }

// Usage

// Here is a possible program to test your function :

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	str := "Hello World!"
// 	nb := StrLen(str)
// 	fmt.Println(nb)
// }

// And its output :

// student@ubuntu:~/[[ROOT]]/test$ go build
// student@ubuntu:~/[[ROOT]]/test$ ./test
// 12
// student@ubuntu:~/[[ROOT]]/test$
///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
)

func main() {
	str := "Hello World!"
	nb := StrLen(str)
	fmt.Println(nb)
}

func StrLen(str string) int {
	ln := 0
	for _, v := range str {
		v = v
		ln += 1
	}
	return ln
}
