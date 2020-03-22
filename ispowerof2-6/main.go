////////////////////////////////////////////////////////////////////////////
// ispowerof2
// allowedFunctions	[…]
// 0	"github.com/01-edu/z01.PrintRune"
// 1	"--cast"
// 2	"strconv.Atoi"
// 3	"os.Args"
// Instructions

// Write a program that determines if a given number is a power of 2.

// This program must print true if the given number is a power of 2, otherwise it prints false.

//     If there is more than one or no argument the program should print a newline (“\n”).

//     When there is only 1 argument, it will always be a positive valid int.

// Usage :

// student@ubuntu:~/ispowerof2$ go build
// student@ubuntu:~/ispowerof2$ ./ispowerof2 2 | cat -e
// true$
// student@ubuntu:~/ispowerof2$ ./ispowerof2 64 | cat -e
// true$
// student@ubuntu:~/ispowerof2$ ./ispowerof2 513 | cat -e
// false$
// student@ubuntu:~/ispowerof2$ ./ispowerof2

// student@ubuntu:~/ispowerof2$ ./ispowerof2 64 1024

// student@ubuntu:~/ispowerof2$
//////////////////////////////////////////////////////////////////////////

package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[1:]
	result := "false"
	defer z01.PrintRune('\n')
	// fmt.Println(arrLen(arg))
	if arrLen(arg) != 1 {
		return
	}
	num, err := strconv.Atoi(arg[0])
	if err == nil {
		for x := 2; x <= num; x *= 2 {
			if x == num {
				result = "true"
			}
		}
	}
	printStr(result)

}
func arrLen(s []string) int {
	count := 0
	for i, _ := range s {
		count = i
		count++
	}
	return count
}
func printStr(s string) {
	for _, v := range s {
		z01.PrintRune(v)
	}
}
