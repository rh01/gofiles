// Copyright 2019 The darwin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import "fmt"

// 递归大法好
func reverse(x int) int {
	res := 0

	walk(&x, &res)
	fmt.Println((1<<2)-1 )
	if res < -1<<31 || res > (1<<31)-1 {
		return 0
	}

	return res
}

func walk(x *int, res *int) {
	if *x == 0 {
		return
	}

	gewei := *x % 10
	shiwei := (*x - gewei) / 10
	*res = *res*10 + gewei
	*x = shiwei
	walk(x, res)
	return
}

func main() {
	fmt.Println(reverse(1563847412))
}
