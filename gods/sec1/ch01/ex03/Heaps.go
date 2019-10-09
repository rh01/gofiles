// Copyright 2019 readailib.com. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//A heap is a data structure that is based on the heap property.
//The heap data structure is used in selection, graph, and k-way merge algorithms.
//Operations such as finding, merging, insertion, key changes,
//and deleting are performed on heaps. Heaps are part of the container/heap package in Go.
//According to the heap order (maximum heap)
//property, the value stored at each node is greater than or equal to its children.
package main

import (
	"container/heap"
	"fmt"
)

// heap 包提供了对于任意类型的堆操作，最小堆是具有每个节点
// 都是以其为根的子树中最小值属性的树
// 默认为小顶堆，即堆顶元素为最小值
// 问题1.怎么样才能构建一个小顶堆
// 必须实现heap的Interface的接口的数据就可以实现堆
//
// 先看一下heap的Interface接口的定义
// type Interface interface {
//     sort.Interface
//     Push(x interface{}) // 向末尾添加元素
//     Pop() interface{}   // 从末尾删除元素
// }
//
// sort.Interface 定义
// type Interface interface {
//     // Len方法返回集合中的元素个数
//     Len() int
//     // Less方法报告索引i的元素是否比索引j的元素小
//     Less(i, j int) bool
//     // Swap方法交换索引i和j的两个元素
//     Swap(i, j int)
// }
// 即我们定义的数据类型必须是实现sort接口且实现Push
// 和Pop方法

// This example demonstrates an integer heap built using the heap interface.

//// An IntHeap is a min-heap of ints.
type IntHeap []int

// 实现sort接口，即要实现Len，Less，Swap方法
func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// 实现Push方法
func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

// 实现Pop方法
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// This example inserts several ints into an IntHeap, checks the minimum,
// and removes them in order of priority.
func Example_intHeap() {
	h := &IntHeap{2, 1, 5}
	heap.Init(h)
	heap.Push(h, 3)
	fmt.Printf("minimum: %d\n", (*h)[0])
	for h.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(h))
	}
	// Output:
	// minimum: 1
	// 1 2 3 5
}

func main() {
	Example_intHeap()
}
