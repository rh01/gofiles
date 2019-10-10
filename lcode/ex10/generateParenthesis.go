// Copyright 2019 readailib.com. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import "fmt"

// 给出 n 代表生成括号的对数，请你写出一个函数，
// 使其能够生成所有可能的并且有效的括号组合。
// n = 3
// [
//   "((()))",
//   "(()())",
//   "(())()",
//   "()(())",
//   "()()()"
// ]
var res []string

func generateParenthesis(n int) []string {
	res = make([]string, 0)
	walk("", 0, 0, n)
	return res
}

func walk(str string, left, right, n int) {

	if left < right || left > n || right > n {
		return
	}

	if left == n && right == n {
		res = append(res, str)
		return
	}

	walk(str+"(", left+1, right, n)
	walk(str+")", left, right+1, n)
	return
}

func main() {
	fmt.Println(generateParenthesis(3))
}
