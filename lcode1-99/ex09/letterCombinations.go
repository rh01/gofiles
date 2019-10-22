// Copyright 2019 readailib.com. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"strconv"
	"strings"
)

func letterCombinations(digits string) []string {
	res := []string{}
	var mapping = []string{"0", "1", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	res = append(res, "")

	digitsArray := strings.Split(digits, "") // []string
	length := len(digitsArray)
	for i := 0; i < length; i++ {
		x, _ := strconv.Atoi(digitsArray[i])
		for len(res[0]) == i {
			remove := res[0]
			res = res[1:]

			for i := 0; i < len(mapping[x]); i++ {
				res = append(res, remove+string(mapping[x][i]))
			}
		}
	}
	return res
}
