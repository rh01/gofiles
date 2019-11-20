// Copyright 2019 readailib.com. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ex176

import "sort"

func hIndex(citations []int) int {
	// IntSlice给[]int添加方法以满足Interface接口，以便排序为递增序列。
	sort.Sort(sort.IntSlice(citations))

	N := len(citations)
	l, r := 0, N
	for l < r {
		mid := l + (r-l)>>2

		tmp := citations[mid]
		if tmp == N-mid {
			return N - mid
		} else if tmp < N-mid {
			l = mid + 1
		} else if tmp > N-mid {
			r = mid
		}
	}
	return N - l
}
