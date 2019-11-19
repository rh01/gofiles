// Copyright 2019 The darwin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ex174

func canCompleteCircuit(gas []int, cost []int) int {
	size := len(gas)
	visited := make([]bool, size) // 这个用来查看开始的是否有效

	for i := 0; i < size; i++ {
		// 如果 gas[i] < cost[i],
		if gas[i] < cost[i] {
			visited[i] = true
			continue
		}

		visited[i] = true
		// temp := make([]int, size)
		j := 0
		curGas := gas[i]
		// curCost := 0
		for {
			if j == size {
				return i%size
			}
			// 这里检查当前gas是否小于cost值
			if curGas < cost[i%size] {
				break
			}
			curGas = curGas - cost[i%size] + gas[(i+1)%size]
			i++
			j++
		}
	}

	return -1
}
