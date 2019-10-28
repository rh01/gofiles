// Copyright 2019 The darwin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"strconv"
)

func grayCode(n int) []int {
	if n <= 0 {
		return []int{}
	}

	res := []string{"0", "1"}
	walk(&res, n, 1)

	result := make([]int, 0)
	for _, value := range res {
		v, _ := strconv.ParseInt(value, 2, 0)
		result = append(result, int(v))
	}
	return result
}

func walk(res *[]string, n int, cur int) {
	if cur == n {
		return
	}

	temp := make([]string, 0)
	for i := 0; i < len(*res); i++ {
		temp = append(temp, "0"+(*res)[i])
	}

	for i := len(*res) - 1; i >= 0; i-- {
		temp = append(temp, "1"+(*res)[i])
	}

	*res = temp

	walk(res, n, cur+1)
	return
}

// func main() {
// 	fmt.Println(grayCode(1))
// }
