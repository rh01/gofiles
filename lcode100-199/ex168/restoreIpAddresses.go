// Copyright 2019 The darwin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ex168

import (
	"fmt"
	"strconv"
)

// 输入: "25525511135"
// 输出: ["255.255.11.135", "255.255.111.35"]
func restoreIpAddresses(s string) []string {
	res := make([]string, 0)
	hmap := make(map[string]bool, 0)
	size := len(s)

	walk(s, &res, size, "", "", 0, 1, 1, &hmap)
	walk(s, &res, size, "", "", 0, 2, 1, &hmap)
	walk(s, &res, size, "", "", 0, 3, 1, &hmap)

	return res
}

// "010010"
// ["0.1.0.010","0.1.00.10","0.1.001.0","0.10.0.10","0.10.01.0","0.100.1.0","01.0.0.10","01.0.01.0","01.00.1.0","010.0.1.0"]
// ["0.10.0.10","0.100.1.0"]

func walk(s string, res *[]string, size int, totalstr string, totalstr2 string, left, right int, cur int, hmap *map[string]bool) {
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
	totalstr2 += strconv.Itoa(cuint)
	if cur == 4 && right == size {

		if _, exist := (*hmap)[totalstr2]; exist {
			return
		}
		(*hmap)[totalstr2] = true
		*res = append(*res, totalstr)
		return
	}

	walk(s, res, size, totalstr+".", totalstr2+".", right, right+1, cur+1, hmap)
	walk(s, res, size, totalstr+".", totalstr2+".", right, right+2, cur+1, hmap)
	walk(s, res, size, totalstr+".", totalstr2+".", right, right+3, cur+1, hmap)
	return
}

func main() {
	fmt.Println(restoreIpAddresses("010010"))
}
