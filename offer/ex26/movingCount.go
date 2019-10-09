// Copyright 2019 readailib.com. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ex26

import "strconv"

var (
	totalCount = 0
)

func movingCount(thread, rows, cols int) int {
	// 声明并初始化初始化一个arr，这个arr可以被叫做访问
	// arr，初始化全部为1
	var visitArr [rows][cols]int
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			visitArr[i][j] = 1
		}
	}

	// 接下来就开始执行我们的递归过程了
	walk(thread, rows, cols, &visitArr, 0, 0)
	return totalCount
}

func walk(thread, rows, cols int, visitArr *[rows][cols]int, i, j int) {
	// 判断递归的终止条件
	if i < 0 || j < 0 || i >= rows || j >= cols || visitArr[i][j] != 1 {
		return
	}

	// 接下来来计算数位之和
	res1 := 0
	res2 := 0
	stri := strconv.Itoa(i)
	strj := strconv.Itoa(j)
	for i := 0; i < len(stri); i++ {
		res1 += int((byte(stri[i]) - '0'))
	}
	for i := 0; i < len(strj); i++ {
		res2 += int((byte(strj[i]) - '0'))
	}

	if res1+res2 > thread {
		return
	}

	// 上面不满足的话
	visitArr[i][j] = 0
	totalCount += 1

	// 接下来继续
	walk(thread, rows, cols, visitArr, i+1, j)
	walk(thread, rows, cols, visitArr, i-1, j)
	walk(thread, rows, cols, visitArr, i, j+1)
	walk(thread, rows, cols, visitArr, i, j-1)
}


