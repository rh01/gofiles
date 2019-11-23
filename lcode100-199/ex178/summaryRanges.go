// Copyright 2019 The darwin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"strconv"
	"strings"
)

// [0,1,2,4,5,7]
// 输出: ["0->2","4->5","7"]
func summaryRanges(nums []int) []string {
	// 初始化一个ret，类型为string的切片
	ret := make([]string, 0)

	i := 0
	N := len(nums)
	for i < N {
		tmp := make([]string, 0)
		tmp = append(tmp, strconv.Itoa(nums[i]))
		for i < N-1 && nums[i+1]-nums[i] == 1 {
			tmp = append(tmp, strconv.Itoa(nums[i+1]))
			i++
		}
		if len(tmp) == 1 {
			ret = append(ret, tmp[0])
		} else {
			ret = append(ret, strings.Join([]string{tmp[0], tmp[len(tmp)-1]}, "->"))
		}
		i++
	}
	return ret
}

func main() {
	a := []int{0, 1, 2, 4, 5}
	fmt.Println(summaryRanges(a))
}
