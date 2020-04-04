package main

import "fmt"

func merge(nums1 []int, nums2 []int) []int {
	var ret = []int{}
	i, j := 0, 0
	m, n := len(nums1), len(nums2)
	for i < m && j < n {
		if nums1[i] < nums2[j] {
			ret = append(ret, nums1[i])
			i++
		} else {
			ret = append(ret, nums2[j])
			j++	
		}
	}
	if i < m {
		ret = append(ret, nums1[i:]...)
	}
	if j < n {
		ret = append(ret, nums2[j:]...)
	}
	return ret
}

func sortArray(nums []int) []int {
	sz := len(nums)
	if sz < 2 {
		return nums
	}
	mid := sz / 2
	l := sortArray(nums[:mid])
	r := sortArray(nums[mid:])
	return merge(l, r)
}

func main() {
	sArr := []int{-2, 3, -5}
	fmt.Println(sortArray(sArr))
}
