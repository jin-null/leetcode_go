package main

import "fmt"

func mySqrt(x int) int {


	res := x

	for res*res > x {
		res = (res + x/res) / 2
	}

	return res
}

func main() {
	n := 10
	output := mySqrt(n)
	fmt.Println(output)
}
