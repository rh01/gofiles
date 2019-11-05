// Copyright 2019 The darwin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ex153

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func trap(height []int) int {
	if height == nil || len(height) == 0 {
		return 0
	}

	size := len(height)
	left := make([]int, size)
	right := make([]int, size)

	left[0] = height[0]
	right[size-1] = height[size-1]
	for i := 1; i < size; i++ {
		left[i] = max(height[i], left[i-1])
	}

	for i := size - 2; i >= 0; i-- {
		right[i] = max(height[i], right[i+1])
	}

	ret := 0
	for i := 0; i < size; i++ {
		ret += min(left[i], right[i]) - height[i]
	}
	return ret
}
