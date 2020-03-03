// Copyright 2019 readailib.com. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ex92

func isUgly(num int) bool {
	if num < 1 {
		return false
	}

	for num%5 == 0 {
		num /= 5
	}

	for num%3 == 0 {
		num /= 3
	}

	for num%2 == 0 {
		num >>= 1
	}

	return num == 1
}
