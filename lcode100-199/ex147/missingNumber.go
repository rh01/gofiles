// Copyright 2019 readailib.com. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ex147

func missingNumber(nums []int) int {
	ret := 0
	for i := 1; i <= len(nums); i++ {
		ret += i ^ nums[i-1]
	}
	return ret
}
