// Copyright 2019 The darwin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ex172

// 输入: [2,3,1,1,4]
// 输出: 2
// 解释: 跳到最后一个位置的最小跳跃数是 2。
//       从下标为 0 跳到下标为 1 的位置，跳 1 步，然后跳 3 步到达数组的最后一个位置。
func canJump(nums []int) bool {
	n := 1
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] >= n {
			n = 1
		} else {
			n++
		}

		if i == 0 && nums[i] > 1 {
			return false
		}
	}

	return true
}
