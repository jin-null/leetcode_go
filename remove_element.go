package main

import "fmt"

func removeElement(nums []int, val int) int {

	if len(nums) == 0 {
		return 0
	}
	i := 0
	for j := 1; j < len(nums); j++ {
		if nums[j] != val {
			nums[i] = nums[j]

			i++
		}
	}

	return i

}

func main() {
	nums := []int{3, 2, 2, 3}
	val := 3
	output := removeElement(nums, val)
	fmt.Println(output)
}
