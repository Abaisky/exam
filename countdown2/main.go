/////////////////////////////////////////////////////////////////////////////////
// countdown
// "allowedFunctions":[
// 	"github.com/01-edu/z01.PrintRune",
// 	"--cast",
// 	"--no-lit=[1-8]"
//  ]
// Instructions
// Write a program that displays all digits in descending order, followed by a newline ('\n').
// Usage
// student@ubuntu:~/[[ROOT]]/test$ go build
// student@ubuntu:~/[[ROOT]]/test$ ./test
// 9876543210
// student@ubuntu:~/[[ROOT]]/test$
///////////////////////////////////////////////////////////////////////////

package main

import "github.com/01-edu/z01"

func main() {
	for i := '9'; i >= '0'; i-- {
		z01.PrintRune(rune(i))
	}
	z01.PrintRune('\n')
}
