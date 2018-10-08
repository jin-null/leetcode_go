package main

import (
	"fmt"
	"strconv"
)

func countAndSay(n int) string {
	var nums = []int{1}
	if n == 1 {
		return "1"
	}
	for i := 1; i < n; i ++ {
		var newNums []int
		var t int = 1
		var j int
		for j = 1; j < len(nums); j++ {
			if nums[j] == nums[j-1] {
				t++
			} else {
				newNums = append(newNums, t, nums[j-1])
				t = 1
			}
		}
		newNums = append(newNums, t, nums[j-1])
		nums = newNums
	}
	var str string = ""
	for i := 0; i < len(nums); i++ {
		str += strconv.FormatInt(int64(nums[i]), 10)
	}
	return str

}

func main() {
	n := 10
	output := countAndSay(n)
	fmt.Println(output)
}
