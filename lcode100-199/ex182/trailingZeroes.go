// Copyright 2019 The darwin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ex182

func fn(n int) int {
	ret := 1
	for n != 0 {
		ret *= n
		n--
	}
	return ret
}

func trailingZeroes(n int) int {
	ret := 0
	for n >= 5 {
		ret += n / 5
		n /= 5
	}
	return ret
}
