// Copyright 2019 readailib.com. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func hammingDistance(x int, y int) int {
	// 转
	strx := strconv.FormatInt(int64(x), 2)
	stry := strconv.FormatInt(int64(y), 2)

	sizex := len(strx)
	sizey := len(stry)

	if sizex < sizey {
		strx = strings.Repeat("0", sizey-sizex) + strx
	}

	if sizex > sizey {
		stry = strings.Repeat("0", sizex-sizey) + stry
	}

	// fmt.Println(strx)
	// fmt.Println(stry)

	ret := 0
	for i := 0; i < len(strx); i++ {
		if strx[i] != stry[i] {
			ret++
		}
	}
	return ret
}

func main() {
	fmt.Println(hammingDistance(1, 4))
}

/*
两个整数之间的汉明距离指的是这两个数字对应二进制位不同的位置的数目。

给出两个整数 x 和 y，计算它们之间的汉明距离。

注意：
0 ≤ x, y < 231.

示例:

输入: x = 1, y = 4

输出: 2

解释:
1   (0 0 0 1)
4   (0 1 0 0)
       ↑   ↑

上面的箭头指出了对应二进制位不同的位置

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/hamming-distance
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
