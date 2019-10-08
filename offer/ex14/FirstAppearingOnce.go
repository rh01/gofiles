// Copyright 2019 readailib.com. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"container/list"
)

/*
# -*- coding:utf-8 -*-
class Solution:
    # 返回对应char
    def FirstAppearingOnce(self):
        # write code here
    def Insert(self, char):
        # write code here
*/

var (
	s      *string
	hmap   = make(map[string]int)
	hqueue = list.New()
)

func FirstAppearingOnce() string {
	data := []byte(*s)
	length := len(data)
	for i := 0; i < length; i++ {
		if hmap[string(data[i])] == 1 {
			return string(data[i])
		}
	}
	return "#"
}

func Insert(char string) {
	_, exist := hmap[char]
	*s += char
	if exist {
		hmap[char] += 1
	} else {
		hmap[char] = 1
	}

}

func main() {
}
