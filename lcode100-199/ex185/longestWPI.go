// Copyright 2019 The darwin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

func longestWPI(hours []int) int {

	size := len(hours) // 初始化新的数组

	helpArray := make([]int, size) // 辅助数组

	for i := 0; i < size; i++ {
		if hours[i] > 8 {
			helpArray[i] = 1
		} else {
			helpArray[i] = -1
		}
	}

	presum := make([]int, size+1) // 前缀和

	for i := 1; i <= size; i++ { // 计算前缀和
		presum[i] = presum[i-1] + helpArray[i-1]
	}

	// 顺序生成单调递减栈,
	stack := make([]int, 0)
	for i := 0; i < size+1; i++ {
		if len(stack) == 0 || presum[stack[len(stack)-1]] > presum[i] {
			stack = append(stack, i) // 这里栈中存放着索引
		}
	}

	// 然后从后往前，求最大长度坡
	i := size
	ans := 0
	for i > ans {
		for len(stack) != 0 && presum[stack[len(stack)-1]] < presum[i] {
			ans = max(ans, i-stack[len(stack) - 1])
			stack = stack[:len(stack)-1]
		}
		i--
	}
	return ans

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
