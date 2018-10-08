package main

import "fmt"

func strStr(haystack string, needle string) int {

	for i := 0; i <= len(haystack)-len(needle); i++ {


		if haystack[i:i+len(needle)] == needle {
			return i
		}


	}
	return 0
}

func main() {
	haystack := "hello"
	needle := "ll"
	output := strStr(haystack, needle)
	fmt.Println(output)
}
