// Copyright 2019 readailib.com. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ex153

func singleNumber(nums []int) int {
	one, two := 0, 0
	var three int
	for _, num := range nums {
		two |= (one & num) // two相应的位等于1表示该位出现2次
		one ^= num // one相应的位等于1，表示该位出现了1次
		three = (one & two) // 表示该位出现3次

		// 如果出现3次，将对应的位置为0
		two &= ^three
		one &= ^three
	}
	return one
}


/*
给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现了三次。找出那个只出现了一次的元素。

说明：

你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？

示例 1:

输入: [2,2,3,2]
输出: 3
示例 2:

输入: [0,1,0,1,0,1,99]
输出: 99

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/single-number-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/