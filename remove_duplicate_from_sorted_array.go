package main

import "fmt"

func removeDuplicates(nums []int) []int {



		if len(nums) == 0{
			return  []int{}
		}
		i:=0
		for j:=1;j<len(nums);j++{
			if nums[j]!=nums[i]{
				i++
				nums[i] = nums[j]
			}

		}
	return nums[:i+1]

}

func main() {
	nums := []int{1,1,2}
	output := removeDuplicates(nums)
	fmt.Println(output)
}
