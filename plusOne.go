package main

import "fmt"

func plusOne(digits []int) []int {

	l := len(digits)
	if l == 0 {
		return []int{1}
	}

	digits[l-1]++

	for i := l - 1; i > 0; i-- {
		if digits[i] < 10 {
			break
		}
		digits[i] -= 10
		digits[i-1]++

		if digits[0] > 9 {
			digits[0] -= 10

			digits = append([]int{1}, digits ...)
		}
	}
	return digits
}

func main() {
	digits := []int{9,9}
	output := plusOne(digits)
	fmt.Println(output)
}
