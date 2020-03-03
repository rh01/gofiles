// Copyright 2019 The darwin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ex183

func scoreOfParentheses(S string) int {
	chars := []byte(S)
	stack := make([]int, 1)

	for _, char := range chars {
		if char == '(' {
			stack = append(stack, 0)
		} else {
			N := len(stack)
			w := stack[N-1]
			v := stack[N-2]
			stack = stack[:N-2]
			stack = append(stack, v+max(2*w, 1))
		}
	}

	return stack[0]
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
