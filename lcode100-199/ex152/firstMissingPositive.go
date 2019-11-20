// Copyright 2019 readailib.com. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import "fmt"

// 构建一个长度为n（nums的长度）的切片，然后遍历nums
// 将大于0并且小于或等于数组长度的值放到指定位置，
// 接下来遍历该数组，取出即可。
func firstMissingPositive(nums []int) int {
	// mx := len(nums)
	// if mx  ==0 {
	// 	return 1
	// }

	// while (0 < nums[i] && nums[i] <= nums.size() && nums[i] != nums[nums[i]-1]) {
	// 	swap(nums[i], nums[nums[i]-1]);
	// }

	//helpArray := make([]int, mx)
	for i := 0; i < len(nums); i++ {
		for 0 < nums[i] && nums[i] <= len(nums) && nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}

	//fmt.Println(helpArray)

	for i := 0; i < len(nums); i++ {
		if nums[i] != i + 1 {
			return i + 1
		}
	}

	return len(nums) + 1
}

func main() {
	fmt.Println(firstMissingPositive([]int{2147483647}))
}
