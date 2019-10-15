// Copyright 2019 readailib.com. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package main

import "fmt"

func getMaximumGold(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])
	//path := make([][]int, 0)
	res := make([]int, 0)
	total := make([]int, 1)
	
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] != 0 {
				walk(grid, rows, cols, i, j, &res, &total)
			}
		}
	}
	return total[0]
}

func walk(grid [][]int /* visit [][]bool, */, rows, cols int, curR, curC int, res, total *[]int) {
	// 确定
	if curR < 0 || curR >= rows || curC < 0 || curC >= cols || grid[curR][curC] == 0 {
		return
	}

	*res = append(*res, grid[curR][curC])
	temp := grid[curR][curC]
	grid[curR][curC] = 0
	moves := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for i := 0; i < 4; i++ {
		move_x := moves[i][0]
		move_y := moves[i][1]
		walk(grid, rows, cols, curR+move_x, curC+move_y, res, total)
	}
	t := 0
	for i := 0; i < len(*res); i++ {
		t += (*res)[i]
	}
	if t > (*total)[0] {
		(*total)[0] = t
	}

	*res = (*res)[:len(*res)-1]

	grid[curR][curC] = temp
	return
}

func main() {
	fmt.Println(getMaximumGold([][]int{
		{
			0, 6, 0,
		},
		{
			5, 8, 7,
		},
		{
			0, 9, 0,
		},
	}))
}
