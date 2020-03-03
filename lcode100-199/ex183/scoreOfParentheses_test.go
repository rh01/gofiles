// Copyright 2019 The darwin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ex183

import "testing"

func TestScoreOfParentheses(t *testing.T) {
	S := "(()())"
	want := scoreOfParentheses(S)
	if want != 4 {
		t.Errorf("expect %d, but got %d", 1, want)
	}
}
