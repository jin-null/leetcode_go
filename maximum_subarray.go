package main

import (
	"fmt"
	"math"
)

func maxSubArray(nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	sum := 0

	max := math.MinInt32
	for i := 0; i < len(nums); i++ {
		sum =maxNum(nums[i],nums[i]+sum)
		max =maxNum(max, sum);
	}
	return max
}


func maxNum(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func main() {
	nums := []int{ -1}
	output := maxSubArray(nums)
	fmt.Println(output)
}
