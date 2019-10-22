// Copyright 2019 The darwin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ex94

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// 递归解法
func swapPairs(head *ListNode) *ListNode {
	// special case
	if head == nil || head.Next == nil {
		return head
	}
	next := head.Next
	head.Next = swapPairs(next.Next)
	next.Next = head
	return next
}

// 非递归，利用空头节点
func swapPairsv2(head *ListNode) *ListNode {
	pre := &ListNode{0, nil}
	pre.Next = head
	temp := pre

	for temp.Next != nil && temp.Next.Next != nil {
		start := temp.Next
		end := temp.Next.Next
		temp.Next = end
		start.Next = end.Next
		end.Next = start
		temp = start
	}
	return pre.Next
}

//给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。
//你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

//示例:
//给定 1->2->3->4, 你应该返回 2->1->4->3.
