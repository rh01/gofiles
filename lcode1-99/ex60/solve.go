// Copyright 2019 readailib.com. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
给定一个二维的矩阵，包含 'X' 和 'O'（字母 O）。

找到所有被 'X' 围绕的区域，并将这些区域里所有的 'O' 用 'X' 填充。

示例:

X X X X
X O O X
X X O X
X O X X
运行你的函数后，矩阵变为：

X X X X
X X X X
X X X X
X O X X
解释:

被围绕的区间不会存在于边界上，换句话说，任何边界上的 'O' 都不会被填充为 'X'。 任何不在边界上，或不与边界上的 'O' 相连的 'O' 最终都会被填充为 'X'。如果两个元素在水平或垂直方向相邻，则称它们是“相连”的。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/surrounded-regions
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package ex60

// 基本思路是对编辑的O进行深搜，并将其标记为小o
// 然后标记完之后，最后将剩下的O标记为X，小o变回O即可
func solve(board [][]byte) {
	if board == nil || len(board) == 0 {
		return
	}
	rows := len(board)
	cols := len(board[0])

	// 对第一列和最后一列进行遍历
	for i := 0; i < rows; i++ {
		dfs(board, i, 0, rows, cols)
		dfs(board, i, cols-1, rows, cols)
	}
	// 对一行和最后一行进行遍历
	for i := 0; i < cols; i++ {
		dfs(board, 0, i,rows, cols)
		dfs(board, rows-1, i,rows, cols)
	}
	// 接下来进行最后的替换
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
			if board[i][j] == 'o' {
				board[i][j] = 'O'
			}
		}
	}
}


func dfs(board [][]byte, row, col ,rows, cols int)  {
	// 首先判断dfs的终止条件
	if row < 0 || row >=rows || col < 0 || col >= cols || board[row][col] == 'o' {
		return
	}

	// 然后判断O的条件
	if board[row][col] == 'O' {
		board[row][col] = 'o'
		dfs(board, row+1, col, rows, cols)
		dfs(board, row-1, col, rows, cols)
		dfs(board, row, col+1, rows, cols)
		dfs(board, row, col-1, rows, cols)
	}
	return
}