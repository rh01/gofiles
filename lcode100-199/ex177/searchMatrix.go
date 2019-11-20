// Copyright 2019 readailib.com. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import "fmt"

func getIndex(start, end int, matrix []int, target int) (bool, int) {
	l1, r1 := 0, end-1
	for l1 < r1 {
		mid1 := l1 + (r1-l1)>>1
		if matrix[mid1] == target {
			return true, mid1
		} else if matrix[mid1] > target {
			r1 = mid1
		} else if matrix[mid1] < target {
			l1 = mid1+1
		}
	}
	if matrix[l1] > target {
		return false, l1-1
	}
	return false, l1
}

func searchMatrix(matrix [][]int, target int) bool {
	if matrix == nil || matrix[0][0] > target {
		return false
	}
	rows, cols := len(matrix), len(matrix[0])
	// 找到第一行中最接近（小于等于）target的索引
	i, j := 0, cols-1
	for i < rows && j >= 0 {
		if(matrix[i][j] == target) {
			return true;
		}
		if(matrix[i][j] < target) {
			i++
		}else {
			j--
		}
	}
	return false
}

func main() {
	matrix := [][]int{
		[]int{1, 1},
	}
	fmt.Println(searchMatrix(matrix, 2))
}

/*
编写一个高效的算法来搜索 m x n 矩阵 matrix 中的一个目标值 target。该矩阵具有以下特性：

每行的元素从左到右升序排列。
每列的元素从上到下升序排列。
示例:

现有矩阵 matrix 如下：

[
  [1,   4,  7, 11, 15],
  [2,   5,  8, 12, 19],
  [3,   6,  9, 16, 22],
  [10, 13, 14, 17, 24],
  [18, 21, 23, 26, 30]
]
给定 target = 5，返回 true。

给定 target = 20，返回 false
*/
