// Copyright 2019 The darwin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ex165

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	l1, l2 := 0, 0
	a, b := headA, headB

	for a != nil {
		a = a.Next
		l1++
	}

	for b != nil {
		b = b.Next
		l2++
	}

	// l1表示最长的
	if l1 < l2 {
		l1, l2 = l2, l1
		headA, headB = headB, headA
	}

	// find
	a = headA
	for l1 != l2 {
		l1--
		a = a.Next
	}

	b = headB
	for a != nil && b != nil && a != b {
		a = a.Next
		b = b.Next
	}

	return a
}
