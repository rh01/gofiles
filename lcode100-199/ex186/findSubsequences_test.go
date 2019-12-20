// Copyright 2019 The darwin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ex186

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFindSubsequences(t *testing.T) {
	res := findSubsequences([]int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15})
	expext := [][]int{
		[]int{4, 6},
		[]int{4, 7},
		[]int{4, 6, 7},
		[]int{4, 6, 7, 7},
		[]int{6, 7},
		[]int{6, 7, 7},
		[]int{7, 7},
		[]int{4, 7, 7},
	}

	fmt.Println(res)
	if reflect.DeepEqual(res, expext) {
		t.Fatalf("expect is %v, but got %v", expext, res)
	}
}
