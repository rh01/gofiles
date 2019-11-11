// Copyright 2019 The darwin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ex170

func subsetsWithDup(nums []int) [][]int {
	res := make([][]int, 0)
	size := len(nums)
	res = append(res, []int{nums[0]})

	walk(nums, res, size, 0, 0)
	return res
}

func walk(nums []int, res [][]int, size int, curi, curj int) {
	if curi < 0 || curi >= size || curj < curi || curj >= size {
		return
	}

	res = append(res, nums[curj])

	walk(nums, res, size, cur+1)
	return
}

// 输入: [1,2,2]
// 输出:
// [
//   [2],
//   [1],
//   [1,2,2],
//   [2,2],
//   [1,2],
//   []
// ]
