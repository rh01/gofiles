// Copyright 2019 readailib.com. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ex172

var res [][]string

func partition(s string) [][]string {
	res := make([][]string, 0)
	cur := make([]string, 0)
	dfs(0, &cur)
}

func dfs(cur int, curset *[]string) {
	res = append(res, *curset)
	
}
