// Copyright 2019 The darwin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ex174

func canCompleteCircuit2(gas []int, cost []int) int {
	size := len(gas)
	totalTank, curTank := 0, 0
	start := 0
	for i := 0; i < size; i++ {
		curTank += (gas[i] - cost[i])
		totalTank += (gas[i] - cost[i])
		if curTank < 0 {
			start = i + 1
			curTank = 0
		}
	}

	if totalTank >= 0 {
		return start
	}
	return -1
}
