///////////////////////////////////////////////////////////////////////////////////////////////////////
// tabmult
// allowedFunctions	[…]
// 0	--cast
// 1	github.com/01-edu/z01.PrintRune
// 2	os.*
// 3	len
// 4	strconv.Atoi
// Instructions
// Write a program that displays a number’s multiplication table.

//     The parameter will always be a strictly positive number that fits in an int. Said number multiplied by 9 will also fit in an int.

// Usage
// student@ubuntu:~/[[ROOT]]/test$ go build
// student@ubuntu:~/[[ROOT]]/test$ ./test 9
// 1 x 9 = 9
// 2 x 9 = 18
// 3 x 9 = 27
// 4 x 9 = 36
// 5 x 9 = 45
// 6 x 9 = 54
// 7 x 9 = 63
// 8 x 9 = 72
// 9 x 9 = 81
// student@ubuntu:~/[[ROOT]]/test$ ./test 19
// 1 x 19 = 19
// 2 x 19 = 38
// 3 x 19 = 57
// 4 x 19 = 76
// 5 x 19 = 95
// 6 x 19 = 114
// 7 x 19 = 133
// 8 x 19 = 152
// 9 x 19 = 171
// student@ubuntu:~/[[ROOT]]/test$ ./test

// student@ubuntu:~/[[ROOT]]/test$
///////////////////////////////////////////////////////////////////////////////////////////////////////
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	fmt.Printf("%T %v", args[0], args[0])
	fmt.Println()
	intargs, err := strconv.Atoi(args[0])
	fmt.Println(err)
	fmt.Printf("%T %v", intargs, intargs)
	fmt.Println()
}
