package main

import (
	"fmt"
)

func lengthOfLastWord(s string) int {
	l := len(s)
	if l == 0 {
		return 0
	}

	res := 0

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			if res != 0 {
				return res
			}

			continue
		}

		res++

	}
	return res
}
func main() {
	s := "a "
	output := lengthOfLastWord(s)
	fmt.Println(output)
}
