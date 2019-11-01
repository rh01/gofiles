// Copyright 2019 readailib.com. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ex150

func findDuplicate(nums []int) int {
	slow, fast := 0, 0
	for {
		slow, fast = nums[slow], nums[nums[fast]]
		if slow == fast {
			break
		}
	}
	fast = 0
	for nums[slow] != nums[fast] {
		slow = nums[slow]
		fast = nums[fast]
	}
	return nums[slow]
}
