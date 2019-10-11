// Copyright 2019 readailib.com. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

示例:

输入: [-2,1,-3,4,-1,2,1,-5,4],
输出: 6
解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。
进阶:

如果你已经实现复杂度为 O(n) 的解法，尝试使用更为精妙的分治法求解。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/maximum-subarray
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package main

import "fmt"

func maxSubArray(nums []int) int {
	length := len(nums)
	for i := 1; i < length; i++ {
		nums[i] = nums[i] + max(nums[i-1], 0)

	}
	max := -2<<32 - 1
	for i := 0; i < length; i++ {
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	narr := []int{-2,1,-3,4,-1,2,1,-5,4}
	fmt.Println(maxSubArray(narr))
}
