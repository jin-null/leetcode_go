package main

import "fmt"

func containsDuplicate(nums []int) string {
	for i := 0; i < len(nums); i++ {
		for j := i+1; j < len(nums); j++ {
			if nums[i] ==nums[j]{
				return  "true"
			}
		}
	}
	return "false"
}

func main() {
	nums := []int{1,1,1,3,3,4,3,2,4,2}
	output := containsDuplicate(nums)
	fmt.Println(output)
}
