// Copyright 2019 The darwin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ex129

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

 func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	// 合并两个有序链表
	cur1, cur2 := l1, l2

	newHead1 := &ListNode{-1, nil}
	res := 0

	// 新的头节点
	cur := newHead1
	for cur1 != nil || cur2 != nil {
		if cur1 == nil {
			res = cur2.Val
			cur2 = cur2.Next
		} else if cur2 == nil {
			res = cur1.Val
			cur1 = cur1.Next
		} else if cur2 != nil && cur1 != nil {
			if cur1.Val > cur2.Val {
				res = cur2.Val
				cur2 = cur2.Next
			} else {
				res = cur1.Val
				cur1 = cur1.Next
			}
		}

		cur.Next = &ListNode{res, nil}
		cur = cur.Next
	}

	return newHead1.Next
}
