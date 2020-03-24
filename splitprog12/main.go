///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// splitprog
// allowedFunctions	[…]
// 0	"fmt.*"
// 1	"len"
// 2	"os.Args"
// 3	"--cast"
// 4	"make"
// WARNING! VERY IMPORTANT!

// For this exercise a function will be tested with the exam own main. However the student still needs to submit a structured program:

// This means that:

//     The package needs to be named package main.
//     The submitted code needs one declared function main(func main()) even if empty.
//     The function main declared needs to also pass the Restrictions Checker(illegal functions tester). It is advised for the student to just empty the function main after its own testings are done.
//     Every other rules are obviously the same than for a program.

// Instructions

// Write a function that separates the words of a string and puts them in a string array.

// The separators are the characters of the charset string given in parameter.
// Expected function

// func Split(str, charset string) []string {

// }

// Usage

// Here is a possible program to test your function :

// package main

// import (
// 	"fmt"
// 	piscine ".."
// )

// func main() {
// 	str := "HelloHAhowHAareHAyou?"
// 	fmt.Println(piscine.Split(str, "HA"))
// }

// And its output :

// student@ubuntu:~/[[ROOT]]/test$ go build
// student@ubuntu:~/[[ROOT]]/test$ ./test
// [Hello how are you?]
// student@ubuntu:~/[[ROOT]]/test$
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
)

func main() {
	str := "HelloHAhowHAareHAyou?"
	fmt.Println(Split(str, "HA"))
}

// func Split(str, charset string) []string {
// 	if charset == "" {
// 		ar := make([]string, len(str))
// 		i := 0
// 		for _, v := range str {
// 			ar[i] = string(v)
// 			i++
// 		}
// 		return ar
// 	}

// 	nw := 0
// 	a := 0
// 	for i := 0; i <= len(str)-len(charset); i++ {
// 		if str[i:i+len(charset)] == charset {
// 			a++
// 			i = i + len(charset)
// 		}
// 	}
// 	ar := make([]string, a+1)
// 	j := 0
// 	for i := 0; i <= len(str)-len(charset); i++ {
// 		if str[i:i+len(charset)] == charset {
// 			ar[j] = str[nw:i]
// 			j++
// 			nw = i + len(charset)
// 			i = i + len(charset)
// 		}
// 	}
// 	ar[j] = str[nw:]
// 	return ar
// }

func Split(str, charset string) []string {
	strCharset := str + charset
	counter := 0
	for i := 0; i <= len(strCharset)-len(charset); i++ {
		if strCharset[i:i+len(charset)] == charset {
			counter++
			i = i + len(charset)
		}
	}
	arr := make([]string, counter)
	arrIndex := 0
	strIndex := 0
	for i := 0; i <= len(strCharset)-len(charset); i++ {
		if strCharset[i:i+len(charset)] == charset {
			arr[arrIndex] = strCharset[strIndex:i]
			strIndex = i + len(charset)
			arrIndex++
			i = i + len(charset)
		}
	}
	return arr

}
