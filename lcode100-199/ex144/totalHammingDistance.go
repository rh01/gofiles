// Copyright 2019 readailib.com. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ex144

import (
	"strconv"
	"strings"
)

func totalHammingDistance(nums []int) int {
	size := len(nums)
	total := 0
	for i := 0; i < 30; i++ {
		oneCount := 0
		temp := 0
		for j := 0; j < size; j++ {
			oneCount += nums[j] & 1
			nums[j] >>= 1
			if nums[j] == 0 {
				temp += 1
			} else {
				temp += 0
			}
		}
		//两两之间的汉明距离的总和实际上为：1的个数*0的个数
		total += oneCount * (size - oneCount)
		if temp == size {
			break
		}
	}
	return total
}

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

/*
两个整数的 汉明距离 指的是这两个数字的二进制数对应位不同的数量。

计算一个数组中，任意两个数之间汉明距离的总和。

示例:

输入: 4, 14, 2

输出: 6

解释: 在二进制表示中，4表示为0100，14表示为1110，2表示为0010。（这样表示是为了体现后四位之间关系）
所以答案为：
HammingDistance(4, 14) + HammingDistance(4, 2) + HammingDistance(14, 2) = 2 + 2 + 2 = 6.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/total-hamming-distance
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
