// Copyright 2019 The darwin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ex156

func twoSum(numbers []int, target int) []int {
	l, r := 0, len(numbers)-1
	for l < r {
		sum := numbers[l] + numbers[r]
		switch {
		case sum > target:
			r--
		case sum < target:
			l++
		case sum == target:
			return []int{l, r}
		}
	}
	return []int{}
}
