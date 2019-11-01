// Copyright 2019 readailib.com. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ex146

import "sort"

func findMaximizedCapital(k int, W int, Profits []int, Capital []int) int {
	// Capital 表示最少的启动资金
	// profit 获得的利润

	ar := [][2]int{}
	for i, c := range Capital {
		ar = append(ar, [2]int{c, Profits[i]})
	}

	sort.Slice(ar, func(i, j int) bool {
		a, b := ar[i], ar[j]
		// less function
		return a[0] < b[0] || (a[0] == b[0] && a[1] < b[1])
	})

	for i := 0; i < k; i++ {
		l, r := 0, len(ar)
		k := 0
		mx := 0
		// 这一步是找到满足W的索引，
		for l < r {
			mid := l + (r-1)>>1
			if ar[mid][0] <= W {
				l = mid + 1
			} else {
				r = mid
			}
		}

		// 接下来从该索引往下找到获得最大利润的
		for j := r - 1; j >= 0; j-- {
			// 找到可以获得Profit的最大的
			if t := ar[j][1]; t > mx {
				mx, k = t, j
			}
		}

		W += mx
		if k < len(ar) {
			copy(ar[k:], ar[k+1:]) // copy ar[k+1:] to ar[K:], and remove the last
			ar = ar[:len(ar)-1]
		}
	}
	return W
}
