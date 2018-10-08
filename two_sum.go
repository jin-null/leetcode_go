package main

import "fmt"

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i,j}
			}
		}
	}
	return []int{}
}

func main() {
	nums := []int{2, 5, 5, 11}
	target := 10
	output := twoSum(nums, target)
	fmt.Println(output)
}
