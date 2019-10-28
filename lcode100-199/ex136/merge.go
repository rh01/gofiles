// Copyright 2019 The darwin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	res := make([]int, 0)
	i, j := 0, 0
	for i < m || j < n {
		switch {
		case i == m:
			res = append(res, nums2[j])
			j++
		case j == n:
			res = append(res, nums1[i])
			i++
		case i < m && j < n:
			if nums1[i] < nums2[j] {
				res = append(res, nums1[i])
				i++
			} else {
				res = append(res, nums2[j])
				j++
			}
		}
	}
	fmt.Println(res)
	for i := 0; i < m+n; i++ {
		nums1[i] = res[i]
	}
}

func main() {

	nums1 := []int{1, 2, 3, 0, 0, 0}
	merge(nums1, 3, []int{2, 5, 6}, 3)
	fmt.Println(nums1)
}
