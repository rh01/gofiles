// Copyright 2019 The darwin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ex180

// 复用上一个问题的答案
func isRectangleOverlap(rec1 []int, rec2 []int) bool {
	// 第一个矩形
	x11, y11 := rec1[0], rec1[1]
	x12, y12 := rec1[2], rec1[3]

	// 第二个矩形
	x21, y21 := rec2[0], rec2[1]
	x22, y22 := rec2[2], rec2[3]

	// 为了方便起见，这里利用排除法，排除掉不相交的情况
	//case1
	if y11 >= y22 || y12 <= y21 {
		return false
	}

	// case2
	if x12 <= x21 || x22 <= x11 {
		return false
	}

	return true
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func computeArea(A int, B int, C int, D int, E int, F int, G int, H int) int {
	rec1 := []int{A, B, C, D}
	rec2 := []int{E, F, G, H}

	s := 0
	if isRectangleOverlap(rec1, rec2) {

		x1 := max(A, E)
		x2 := min(C, G)

		y1 := max(B, F)
		y2 := min(D, H)

		s = (x2 - x1) * (y2 - y1)

	}
	return (C-A)*(D-B) + (G-E)*(H-F) - s

}
