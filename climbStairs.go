package main

import "fmt"

func climbStairs(n int) int {

	if n < 2 {
		return 1
	}

	res := make([]int, n+1)
	res[0], res[1] = 1, 1

	for i := 2; i <= n; i++ {
		res[i] = res[i-1] + res[i-2]
	}
	return res[n]
}

func main() {
	nums := 2
	output := climbStairs(nums)
	fmt.Println(output)
}
