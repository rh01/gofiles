// Copyright 2019 The darwin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ex140

// 给定一个非负整数 numRows，生成杨辉三角的前 numRows 行。
func generate(numRows int) [][]int {
	res := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		res[i] = make([]int, i+1)
	}

	res[0][0] = 1
	for i := 1; i < numRows; i++ {
		for j := 0; j <= i; j++ {
			if j == 0 {
				res[i][j] = 1
			} else if j == i {
				res[i][j] = 1
			} else {
				res[i][j] = res[i-1][j] + res[i-1][j-1]
			}
		}
	}
	return res
}
