// Copyright 2019 The darwin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ex164

type MinStack struct {
	minStack  []int
	dataStack []int
	size      int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		minStack:  make([]int, 0),
		dataStack: make([]int, 0),
		size:      0,
	}
}

func (this *MinStack) Push(x int) {
	if this.size == 0 {
		this.dataStack = append(this.dataStack, x)
		this.minStack = append(this.minStack, x)
	}

	minVal := this.GetMin()
	if x > minVal {
		this.dataStack = append(this.dataStack, x)
		this.minStack = append(this.minStack, minVal)
	} else {
		this.dataStack = append(this.dataStack, x)
		this.minStack = append(this.minStack, x)
	}
	this.size++
}

func (this *MinStack) Pop() {
	if this.size > 0 {
		this.dataStack = this.dataStack[:this.size-1]
		this.minStack = this.minStack[:this.size-1]
		this.size--
	}
}

func (this *MinStack) Top() int {
	if this.dataStack != nil && this.size != 0 {
		return this.dataStack[this.size-1]
	}
	return -1
}

func (this *MinStack) GetMin() int {
	if this.minStack != nil && this.size != 0 {
		return this.minStack[len(this.minStack)-1]
	}
	return -1
}
