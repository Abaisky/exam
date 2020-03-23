/////////////////////////////////////////////////////////////////////////////////////////////////////
// hiddenp
// allowedFunctions	[…]
// 0	"len"
// 1	"--cast"
// 2	"os.Args"
// 3	"github.com/01-edu/z01.PrintRune"
// Instructions

// Write a program named hiddenp that takes two string and that, if the first string is hidden in the second one, displays 1 followed by a newline ('\n'), otherwise it displays 0 followed by a newline.

// Let s1 and s2 be string. It is considered that s1 is hidden in s2 if it is possible to find each character from s1 in s2, in the same order as they appear in s1.

// If s1 is an empty string it is considered hidden in any string.

// If the number of parameters is different from 2, the program displays a newline.
// Usage

// student@ubuntu:~/[[ROOT]]/hiddenp$ go build
// student@ubuntu:~/[[ROOT]]/hiddenp$ ./hiddenp "fgex.;" "tyf34gdgf;'ektufjhgdgex.;.;rtjynur6" | cat -e
// 1$
// student@ubuntu:~/[[ROOT]]/hiddenp$ ./hiddenp "abc" "2altrb53c.sse" | cat -e
// 1$
// student@ubuntu:~/[[ROOT]]/hiddenp$ ./hiddenp "abc" "btarc" | cat -e
// 0$
// student@ubuntu:~/[[ROOT]]/hiddenp$ ./hiddenp "DD" "DABC" | cat -e
// 0$
// student@ubuntu:~/[[ROOT]]/hiddenp$ ./hiddenp | cat -e
// $
// student@ubuntu:~/[[ROOT]]/hiddenp$

/////////////////////////////////////////////////////////////////////////////////////////////////////

package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	defer z01.PrintRune('\n')
	counter := 0
	if len(os.Args) != 3 {
		return
	}
	for _, v := range os.Args[2] {
		if len(os.Args[1]) == counter {
			z01.PrintRune('1')
			return
		}
		if v == rune(os.Args[1][counter]) {
			counter++
		}
	}
	z01.PrintRune('0')

}
