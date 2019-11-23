// Copyright 2019 The darwin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ex182

import (
	"testing"
)

func TestTrailingZeroes(t *testing.T) {
	// fmt.Println(fn(7))
	res := trailingZeroes(15)
	wanted := 3
	if res != wanted {
		t.Errorf("wanted: %d, expect: %d\n", res, wanted)
	}
}
