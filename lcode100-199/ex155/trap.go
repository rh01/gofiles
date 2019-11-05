// Copyright 2019 The darwin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import "fmt"

func trap(height []int) int {
	size := len(height)
	l, r := 0, size-1
	lmax, rmax := 0, 0
	ret := 0
	for l <= r {
		if height[l] <= height[r] {
			if height[l] > lmax {
				lmax = height[l]
			} else {
				ret += lmax - height[l]
			}
			l++
		} else {
			if height[r] > rmax {
				rmax = height[r]
			} else {
				ret += rmax - height[r]
			}
			r--
		}
	}
	return ret
}

func main() {
	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
}
