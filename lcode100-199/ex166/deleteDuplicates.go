// Copyright 2019 The darwin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import "fmt"

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	// special case
	if head == nil || head.Next == nil {
		return head
	}
	before, after := head, head.Next

	for after != nil {
		for after != nil && before.Val == after.Val {
			after = after.Next
		}
		before.Next = after
		before = before.Next
		if after == nil {
			break
		}
		after = after.Next
	}
	return head
}

func main() {
	head := &ListNode{
		2,
		&ListNode{
			2,
			&ListNode{
				2,
				nil,
			},
		},
	}
	head2 := deleteDuplicates(head)
	for head2 != nil {
		fmt.Println(head2.Val)
		head2 = head2.Next
	}
}
