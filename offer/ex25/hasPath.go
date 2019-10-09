// Copyright 2019 readailib.com. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ex25

/*
# -*- coding:utf-8 -*-
class Solution:
    def hasPath(self, matrix, rows, cols, path):
        # write code here
*/

// 矩阵中的路径
func hasPath(matrix []string, rows int, cols int, path string) bool {
	// 这个是标记数组，用来标记当前是否访问过
	flag := make([]bool, len(matrix))

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if judge(matrix, rows, cols, flag, i, j, path, 0) {
				return true
			}
		}
	}
	return false
}

// level as the current charactor
func judge(matrix []string, rows int, cols int, flag []bool, i int, j int, path string, level int) bool {

	// 通过i，j计算当前的 索引位置
	index := i*cols + j

	// 判断是否满足递归的终止条件
	if i < 0 || j < 0 || i >= rows || j >= cols || matrix[index] != string(path[level]) || flag[index] == true {
		return false
	}

	// 标记当前的已被访问
	flag[index] = true

	if judge(matrix, rows, cols, flag, i+1, j, path, level+1) ||
		judge(matrix, rows, cols, flag, i-1, j, path, level+1) ||
		judge(matrix, rows, cols, flag, i, j+1, path, level+1) ||
		judge(matrix, rows, cols, flag, i, j-1, path, level+1) {
		return true
	}

	// 如果不满足上面的条件，将上面的标记为false
	flag[index] = false
	return false
}
