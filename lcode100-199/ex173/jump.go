// Copyright 2019 The darwin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ex173

// [2,3,1,1,4]
// 2 -> 3
// 2 -> 1
// 3 > 1
func jump(nums []int) int {
	cur := 0
	size := len(nums)
	count := 0
	for cur < size-1 {
		step := nums[cur]
		max := 0
		temp := 0
		// 检查如果当前的值已经到达或者超过size-1，那么就直接返回
		if cur+step >= size-1 {
			return count + 1
		}
		// 遍历，选择 nums[i] + i 最大的
		for i := cur + 1; i <= cur+step; i++ {
			curval := nums[i] + i
			if curval > max {
				max = curval
				temp = i
			}
		}
		cur = temp
		count++
	}
	return count
}
