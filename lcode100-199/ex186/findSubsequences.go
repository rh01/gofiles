// Copyright 2019 The darwin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ex186

func findSubsequences(nums []int) [][]int {
	if nums == nil || len(nums) < 2 {
		return [][]int{}
	}

	ret := make([][]int, 0)
	walk(nums, 0, []int{}, &ret)

	return ret
}

func walk(nums []int, start int, curset []int, ret *[][]int) {
	if len(curset) >= 2 {
		t1 := make([]int, len(curset))
		copy(t1, curset)
		*ret = append(*ret, t1)
	}

	for i := start; i < len(nums); i++ {
		if i != start && nums[i] == nums[i-1] {
			continue
		}

		if (curset != nil && len(curset) != 0) && nums[i] < curset[len(curset) - 1] {
			continue
		}
		curset = append(curset, nums[i])
		walk(nums, i+1, curset, ret)
		curset = curset[:len(curset) - 1]
	}
	return
}



func findSubsequencesV2(nums []int) [][]int {
	o, t := [][]int{}, []int{}
	var cal func(int)
	cal = func(i0 int) {
		if len(t) >= 2 {
			t1 := make([]int, len(t))
			copy(t1, t)
			o = append(o, t1)
		}
		s := map[int]bool{}
		for i := i0; i < len(nums); i++ {
			n := nums[i]
			if (0 != len(t) && t[len(t)-1] > n) || s[n] {
				continue
			}
			t, s[n] = append(t, n), true
			cal(i + 1)
			t = t[:len(t)-1]
		}
	}
	cal(0)
	return o
}
