// Copyright 2019 readailib.com. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import "fmt"

/*
给定一个包括 n 个整数的数组 nums 和 一个目标值 target。
找出 nums 中的三个整数，使得它们的和与 target 最接近。
返回这三个数的和。假定每组输入只存在唯一答案。

例如，给定数组 nums = [-1，2，1，-4], 和 target = 1.

与 target 最接近的三个数的和为 2. (-1 + 2 + 1 = 2).

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/3sum-closest
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

//给定数组 nums = [-1，2，1，-4], 和 target = 1.
//与 target 最接近的三个数的和为 2. (-1 + 2 + 1 = 2).
func threeSumClosest(nums []int, target int) int {
	size := len(nums)
	ISort3(&nums)
	fmt.Println(nums)
	res := 0
	diff := 2 << 31
	for i := 0; i < size-2; i++ {
		num1 := nums[i]
		// 声明两个指针
		low, high := i+1, size-1

		
		for low < high {
			sum := nums[low] + nums[high] + num1
			cdiff := abs(sum - target)
			fmt.Println(sum)
			fmt.Println(cdiff)
			fmt.Println(diff)
			if cdiff < diff {
				diff = cdiff
				res = sum
			}
			
			if sum < target {
				low++
			} else if sum > target {
				high--
			} else {
				return sum
			}
		}
	}
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func ISort3(x *[]int) {
	for i := 0; i < len(*x); i++ {
		t := (*x)[i]
		j := i
		for ; j > 0 && (*x)[j-1] > t; j-- {
			(*x)[j] = (*x)[j-1]
		}
		(*x)[j] = t
	}
}

func main() {
	fmt.Println(threeSumClosest([]int{1, 1, 1, 0}, -100))
}
