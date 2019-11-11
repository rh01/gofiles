// Copyright 2019 The darwin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ex171

// bit computation
func subsets(nums []int) [][]int {
	S := len(nums)
	N := 1 << uint(S)
	res := make([][]int, 0)
	for i := 0; i < N; i++ {
		v := make([]int, 0)
		for j := 0; j < S; j++ {
			if i&(1<<uint(j)) == 0 {
				v = append(v, nums[j])
			}
		}
		res = append(res, v)
	}
	return res
}
