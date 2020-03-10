package main

import "fmt"

func main() {
	str := "HelloHAhowHAareHAyou?"
	fmt.Println(Split(str, "HA"))
}

func Split(str, charset string) []string {
	end := str + charset
	nbword := 0
	for i := 0; i <= len(end)-len(charset); i++ {
		if end[i:i+len(charset)] == charset {
			nbword++
		}
	}
	fmt.Println(nbword)
	arr := make([]string, nbword)
	index := 0
	arrIndex := 0
	for i := index; i <= len(end)-len(charset); i++ {
		if end[i:i+len(charset)] == charset {
			arr[arrIndex] = end[index:i]
			index = i + len(charset)
			arrIndex++
		}
	}
	return arr
}
