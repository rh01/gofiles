// Copyright 2019 The darwin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package main

import "fmt"

func judgeCircle(moves string) bool {
	// 先做最简单的判断, 剪枝
	size := len(moves)
	if size%2 != 0 {
		return false
	}

	ud_stack, lr_stack := make([]byte, 0), make([]byte, 0)
	for _, move := range moves {
		switch move {
		case 'U':
			if len(ud_stack) == 0 || ud_stack[len(ud_stack)-1] != byte(move) {
				ud_stack = append(ud_stack, byte(move))
			} else {
				ud_stack = ud_stack[:len(ud_stack)-1]
			}
		case 'D':
			if len(ud_stack) == 0 || ud_stack[len(ud_stack)-1] != byte(move) {
				ud_stack = append(ud_stack, byte(move))
			} else {
				ud_stack = ud_stack[:len(ud_stack)-1]
			}
		case 'L':
			if len(lr_stack) == 0 || lr_stack[len(lr_stack)-1] != byte(move) {
				lr_stack = append(lr_stack, byte(move))
			} else {
				lr_stack = lr_stack[:len(lr_stack)-1]
			}
		case 'R':
			if len(lr_stack) == 0 || lr_stack[len(lr_stack)-1] != byte(move) {
				lr_stack = append(lr_stack, byte(move))
			} else {
				lr_stack = lr_stack[:len(lr_stack)-1]
			}
		}
	}
	return (len(ud_stack) + len(lr_stack)) == 0
}

func main() {
	fmt.Println(judgeCircle("LLLLRULLRDDLDRRDLDRLDURLUUDRRLUDULURLLLLURUDRULDDLRLLLULLLLLULRUDRRRLDDLDLRULLDRUUDLRLDDURRUDRULLRLU"))
}
