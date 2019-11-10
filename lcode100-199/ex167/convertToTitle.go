// Copyright 2019 The darwin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import "fmt"

// 168
func convertToTitle(n int) string {
	res := ""
	for n != 0 {
		tmp := (n-1) % 26
		n = (n - tmp) / 26
		// fmt.Println(string(byte(tmp + 65)))
		res = string(byte(tmp+65)) + res
	}
	return res
}

func main() {
	fmt.Println('A', 'Z')
	fmt.Println(convertToTitle(26))
	fmt.Println(convertToTitle(28))
	fmt.Println(convertToTitle(701))
}
