// Copyright 2019 The darwin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import "fmt"

func divide(dividend int, divisor int) int {
	res := 0
	flag := false
	if divisor < 0 {
		flag = true
		divisor = -divisor
	}
	for dividend > divisor {
		dividend -= divisor
		res++
	}

	

	// [−231,  231 − 1]
	if flag {
		res = -res
	}

	if res > 1<<31-1 || res < - 1<< 31 {
		return 1<<31 - 1
	}

	return res

}

func main() {
	fmt.Println(1 << 31)
	fmt.Println(2 << 31)
	fmt.Println(divide(7, -3))
}
