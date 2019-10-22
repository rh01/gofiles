// Copyright 2019 readailib.com. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
给定一个由 '1'（陆地）和 '0'（水）组成的的二维网格，计算岛屿的数量。
一个岛被水包围，并且它是通过水平方向或垂直方向上相邻的陆地连接而成的。
你可以假设网格的四个边均被水包围。

示例 1:

输入:
11110
11010
11000
00000

输出: 1
示例 2:

输入:
11000
11000
00100
00011

输出: 3

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/number-of-islands
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package ex62

func numIslands(grid [][]byte) int {
	// 首先判断临界条件
	if grid == nil || len(grid) == 0 {
		return 0
	}
	var res = 0
	// 获得行数和列数
	rows := len(grid)
	cols := len(grid[0])
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == '1' {
				dfs(&grid, rows, cols, i, j)
				res++
			}
		}
	}

	return res
}

func dfs(grid *[][]byte, rows, cols, x, y int) {
	// 判断递归的终止条件
	if x < 0 || x >= rows || y < 0 || y >= cols || (*grid)[x][y] == '0' {
		return
	}

	// 主要逻辑
	(*grid)[x][y] = '0'
	dfs(grid, rows, cols, x+1, y)
	dfs(grid, rows, cols, x-1, y)
	dfs(grid, rows, cols, x, y+1)
	dfs(grid, rows, cols, x, y-1)

	return
}
