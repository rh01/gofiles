// Copyright 2019 The darwin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"strings"
)

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	size := len(s)
	l, r := 0, size-1
	for l <= r {
		lStr := s[l]
		rStr := s[r]
		if !((48 <= lStr && lStr <= 57) || (97 <= lStr && lStr <= 122)) {
			l++
			continue
		}

		if !((48 <= rStr && rStr <= 57) || (97 <= rStr && rStr <= 122)) {
			r--
			continue
		}

		if lStr != rStr {
			return false
		}

		l++
		r--
	}
	return true
}

func main() {
	fmt.Println('0', '9', 'a', 'z')
}

/*
给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。

说明：本题中，我们将空字符串定义为有效的回文串。

示例 1:

输入: "A man, a plan, a canal: Panama"
输出: true
示例 2:

输入: "race a car"
输出: false

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/valid-palindrome
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
