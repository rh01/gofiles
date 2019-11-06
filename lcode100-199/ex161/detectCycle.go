// Copyright 2019 The darwin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ex161

/**
 * Definition for singly-linked list.
 * */
type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	isCycle := false
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast != nil && slow == fast {
			isCycle = true
			break
		}
	}

	if !isCycle {
		return nil
	}

	fast = head
	for fast != slow {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}