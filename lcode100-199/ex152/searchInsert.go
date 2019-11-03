// Copyright 2019 The darwin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import "fmt"

func searchInsert(nums []int, target int) int {

	l, r := 0, len(nums)-1
	if target > nums[r] {
		return r + 1
	}

	for l < r {
		mid := l + (r-l-1)>>2
		switch {
		case nums[mid] == target:
			return mid
		case nums[mid] > target:
			r = mid - 1
		case nums[mid] < target:
			l = mid + 1
		}
	}
	return l
}

func main() {
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 7))
}
