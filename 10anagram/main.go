package main

import (
	"fmt"
)

func main() {

	test1 := IsAnagram("listen", "silent")
	fmt.Println(test1)

	test2 := IsAnagram("alem", "school")
	fmt.Println(test2)

	test3 := IsAnagram("neat", "a net")
	fmt.Println(test3)

	test4 := IsAnagram("anna madrigal", "a man and a girl")
	fmt.Println(test4)
}

func IsAnagram(s1, s2 string) bool {

	k := map[rune]int{}
	fmt.Println(k)

	for _, v := range s1 {
		if v != ' ' {
			k[v]++
		}
	}
	// fmt.Println(k)
	for _, g := range s2 {
		if g != ' ' {
			// fmt.Println(g)
			k[g]--
		}
	}
	// fmt.Println(k)
	for _, p := range k {
		if p != 0 {
			// fmt.Println(p)
			return false
		}

	}
	return true
}

// func IsAnagram(str1, str2 string) bool {
// 	flag := false
// 	newstr1 := ""
// 	// balance := 0
// 	j := 0
// 	for i := 0; i < len(str1); i++ {
// 		if str1[i] != ' ' {
// 			newstr1 = newstr1[:j] + str1[i:]
// 			j++
// 		}
// 	}
// 	return flag
// }

// func IsAnagram(str1, str2 string) bool {

// 	for k := 0; k < len(str1); k++ {
// 		c1 := str1[k]

// 		found := false

// 		for i := 0; i < len(str2); i++ {
// 			c2 := str2[i]
// 			if c1 == c2 {
// 				found = true
// 				// fmt.Println(str2)
// 				str2 = str2[:i] + str2[i+1:]
// 				i--
// 				// fmt.Println(str2)

// 				break
// 			}
// 		}

// 		if !found {
// 			return false
// 		}
// 	}

// 	return true
// }

// func IsAnagram(str1, str2 string) bool {
// 	str := str1 + str2
// 	sum := byte(0)
// 	for _, v := range str {
// 		if v == ' ' {
// 			continue
// 		}
// 		sum ^= byte(v)
// 	}
// 	//fmt.Println(string(sum))
// 	return sum == 0
// }
