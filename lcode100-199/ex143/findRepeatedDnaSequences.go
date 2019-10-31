// Copyright 2019 readailib.com. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ex143

func findRepeatedDnaSequences(s string) []string {
	// 检查是否小于10
	size := len(s)
	if size < 10 {
		return []string{}
	}

	// 初始化两个hset
	res := make(map[string]bool, 0)
	ret := make(map[string]bool, 0)
	result := make([]string, 0)

	for i := 0; i < size-10; i++ {
		str := s[i:i+10]
		_, exist := res[str]
		if exist {
			_, ok := ret[str]
			if ok {
				continue
			} else {
				ret[str] = true
				result = append(result, str)
			}
		} else {
			res[str] = true
		}
	}
	return result
}

/*
所有 DNA 都由一系列缩写为 A，C，G 和 T 的核苷酸组成，
例如：“ACGAATTCCG”。在研究 DNA 时，识别 DNA 中的重复序列有时会对研究非常有帮助。

编写一个函数来查找 DNA 分子中所有出现超过一次的 10 个字母长的序列（子串）。



示例：

输入：s = "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"
输出：["AAAAACCCCC", "CCCCCAAAAA"]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/repeated-dna-sequences
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
