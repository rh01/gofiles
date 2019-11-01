// Copyright 2019 readailib.com. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ex148

func singleNumber(nums []int) int {
	res := 0
	for i:=0; i < len(nums); i++ {
		res ^= nums[i]
	}
	return res
}