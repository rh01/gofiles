// Copyright 2019 readailib.com. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ex175

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	// fakeHead point to the head
	fakeHead := &ListNode{-1, head}

	pre, cur := fakeHead, fakeHead.Next
	for cur != nil {
		for cur != nil && cur.Val == val {
			pre.Next = cur.Next
			cur = cur.Next
		}

		if cur == nil {
			break
		}
		cur = cur.Next
		pre = pre.Next
	}

	return fakeHead.Next
}
