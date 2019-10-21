// Copyright 2019 The darwin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import "fmt"

//给定 nums = [2, 7, 11, 15], target = 9
//因为 nums[0] + nums[1] = 2 + 7 = 9
//所以返回 [0, 1]
//这种解法只适用于有序
func twoSumSorted(nums []int, target int) []int {
	low, high := 0, len(nums)-1
	for low < high {
		sum := nums[low] + nums[high]
		if sum < target {
			low++
		} else if sum > target {
			high--
		} else {
			return []int{low, high}
		}
	}

	return []int{}
}

func twoSum(nums []int, target int) []int {
	hmap := make(map[int]int, 0)
	for index1, value := range nums {
		if _, exist := hmap[value]; !exist {
			hmap[value] = index1
		}

		req := target - value
		if index2, exist := hmap[req]; exist && index1 != index2 {
			return []int{index2, index1}
		}
	}
	return []int{}
}

func main() {
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
}
