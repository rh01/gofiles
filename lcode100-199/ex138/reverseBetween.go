// Copyright 2019 readailib.com. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ex138

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil || m > n {
		return nil
	}
	// 	先考虑一般情况
	pre := &ListNode{0, nil}
	pre.Next = head

	//cur := pre

	first, end := pre, pre
	// 此时的pre就是m的
	var i int
	i = 1
	for ; i < m; i++ {
		first = first.Next
		// cur = cur.Next
		end = end.Next
	}

	// 找到tail节点
	for ; end != nil && i < n+1; i++ {
		end = end.Next
	}

	for first.Next != end {
		cur := first.Next
		first.Next = cur.Next
		cur.Next = end.Next
		end.Next = cur
	}
	return pre.Next
}
