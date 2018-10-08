package main

import (
	"fmt"
	"strings"
)

func toGoatLatin(S string) string {
	var sp []string
	sp = strings.Split(S, " ")

	for i := range sp {
		sp[i] = handleWord(sp[i], i)
	}
	return strings.Join(sp, " ")
}

func handleWord(s string, i int) string {

	postfix := "ma" + strings.Repeat("a", i+1)

	if isBeginWithVowel(s) {
		return s + postfix
	}

	return s[1:] + s[0:1] + postfix

}

func isBeginWithVowel(s string) bool {

	switch s[0] {
	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
		return true
	default:
		return false
	}

}
func main() {
	s := "I speak Goat Latin"
	output := toGoatLatin(s)
	fmt.Println(output)
}
