package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[1:]
	if len(arg) == 1 {
		ans := ""
		ff := ""
		tmp := ""
		arg[0] += " "
		for _, c := range arg[0] {
			if c == ' ' && tmp != "" {
				if ff == "" {
					ff = tmp

					tmp = ""
					continue
				}
				if ans == "" {
					ans = tmp
				} else {
					ans = ans + string(" ") + tmp
				}
				tmp = ""
			} else if c != ' ' {
				tmp = tmp + string(c)
			}
		}
		if ans == "" {
			ans = ff
		} else {
			ans = ans + string(" ") + ff
		}
		for _, c := range ans {
			z01.PrintRune(c)
		}
	}
	z01.PrintRune('\n')
}
