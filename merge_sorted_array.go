package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) []int {
	temp := make([]int, m)
	copy(temp, nums1)

	j, k := 0, 0
	for i := 0; i < len(nums1); i++ {
		// nums2用完了
		if k >= n {
			nums1[i] = temp[j]
			j++
			continue
		}
		// temp 用完了
		if j >= m {
			nums1[i] = nums2[k]
			k++
			continue
		}
		// 比较后，放入
		if temp[j] < nums2[k] {
			nums1[i] = temp[j]
			j++
		} else {
			nums1[i] = nums2[k]
			k++
		}
	}
	return nums1
}

func main() {
	nums1 := []int{1,2,3,0,0,0}
	m := 3
	nums2 := []int{2,5,6}
	n := 3
	output := merge(nums1,m,nums2,n)
	fmt.Println(output)
}
