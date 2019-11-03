// Copyright 2019 The darwin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ex154

func productExceptSelf(nums []int) []int {
	l, r := 1, 1
	n := len(nums)
	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = 1
	}

	for i := 0; i < n; i++ {
		res[i] *= l
		l *= nums[i]

		res[n-i-1] *= r
		r *= nums[n-1-i]
	}
	return res
}
