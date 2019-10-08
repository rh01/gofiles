// Copyright 2019 The rh01 Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package leetcode003

import (
	"strings"
)



// 返回最长的子字符串
func lengthOfLongestSubstring1(s string) int {
	bs := []byte(s)
	var bsMaxLen, start int   // bsMaxLen表示最长字符串的长度，start表示刚开始出现的相同字符的下一个位置
	buckets := map[byte]int{} // 存储着当前遇到相同字符的索引

	for i := 0; i < len(bs); i++ {
		if _, exist := buckets[bs[i]]; exist && start <= buckets[bs[i]] {
			start = buckets[bs[i]] + 1
		} else {
			if bsMaxLen < i-start+1 {
				bsMaxLen = i - start + 1
			}
		}
		buckets[bs[i]] = i
	}
	return bsMaxLen
}

// go:path
func lengthOfLongestSubstringBrute(s string) int {
	if len(s) < 2 {
		return len(s)
	}
	length := len(s)
	i := 0
	maxLength := 0
	for ; i < length-1; i += 1 {
		curLength := 1
		for j := i + 1; j < length; j += 1 {
			if !strings.Contains(s[i:j], string(s[j])) {
				curLength += 1
			} else {
				break
			}
		}
		if curLength > maxLength {
			maxLength = curLength
		}
	}
	return maxLength
}
