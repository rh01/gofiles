// Copyright 2019 The darwin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"strconv"
)

// 输入: "25525511135"
// 输出: ["255.255.11.135", "255.255.111.35"]
func restoreIpAddresses(s string) []string {
	res := make([]string, 0)
	size := len(s)

	walk(s, &res, size, "", 0, 1, 1)
	walk(s, &res, size, "", 0, 2, 1)
	walk(s, &res, size, "", 0, 3, 1)

	return res
}

// "010010"
// ["0.1.0.010","0.1.00.10","0.1.001.0","0.10.0.10","0.10.01.0","0.100.1.0","01.0.0.10","01.0.01.0","01.00.1.0","010.0.1.0"]
// ["0.10.0.10","0.100.1.0"]

func walk(s string, res *[]string, size int, totalstr string, left, right int, cur int) {
	// 判断越界条件
	if left > size-1 || right > size || cur > 4 || (right-left > 1 && s[left] == '0') {
		return
	}
	cstr := s[left:right]
	cuint, _ := strconv.Atoi(cstr)
	if cuint > 255 {
		return
	}

	totalstr += cstr
	if cur == 4 && right == size {
		*res = append(*res, totalstr)
		return
	}

	walk(s, res, size, totalstr+".", right, right+1, cur+1)
	walk(s, res, size, totalstr+".", right, right+2, cur+1)
	walk(s, res, size, totalstr+".", right, right+3, cur+1)
	return
}

func main() {
	fmt.Println(restoreIpAddresses("010010"))
}
