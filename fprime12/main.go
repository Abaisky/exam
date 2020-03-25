////////////////////////////////////////////////////////////////////////////////////////////
// fprime
// allowedFunctions	[…]
// 0	"strconv.Atoi"
// 1	"os.Args"
// 2	"len"
// 3	"--cast"
// 4	"fmt.*"
// Instructions

// Write a program that takes a positive int and displays its prime factors, followed by a newline ('\n').

//     Factors must be displayed in ascending order and separated by *.

//     If the number of parameters is different from 1, the program displays a newline.

//     The input, when there is one, will always be valid.

// Usage

// student@ubuntu:~/[[ROOT]]/test$ go build
// student@ubuntu:~/[[ROOT]]/test$ ./test 225225
// 3*3*5*5*7*11*13
// student@ubuntu:~/[[ROOT]]/test$ ./test 8333325
// 3*3*5*5*7*11*13*37
// student@ubuntu:~/[[ROOT]]/test$ ./test 9539
// 9539
// student@ubuntu:~/[[ROOT]]/test$ ./test 804577
// 804577
// student@ubuntu:~/[[ROOT]]/test$ ./test 42
// 2*3*7
// student@ubuntu:~/[[ROOT]]/test$ ./test a

// student@ubuntu:~/[[ROOT]]/test$ ./test 0

// student@ubuntu:~/[[ROOT]]/test$ ./test 1

// student@ubuntu:~/[[ROOT]]/test$
////////////////////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	defer fmt.Println()
	if len(os.Args) != 2 {
		return
	}
	num, err := strconv.Atoi(os.Args[1])
	if err != nil {
		return
	}
	cop := num
	result := ""
	test := 1
	for i := 2; i <= num; i++ {
		if num%i == 0 {
			test = test * i
			result += "*" + myItoa(i)
			num = num / i
			i = i - 1
		}
	}
	if test == cop && len(result) > 1 {
		fmt.Print(result[1:])
	}

}
func myItoa(i int) string {
	str := ""
	if i >= 10 {
		for i > 0 {
			str = string(i%10+'0') + str
			i = i / 10
		}
	} else {
		str = string(i + '0')
	}
	return str
}

func prime(i int) bool {
	if i > 1 {
		for n := 2; n < i; n++ {
			if i%n == 0 {
				return false
			}
		}
		return true
	}
	return false
}
