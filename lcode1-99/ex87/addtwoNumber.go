// Copyright 2019 readailib.com. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ex87

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

// 输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
// 输出：7 -> 0 -> 8
// 原因：342 + 465 = 807
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	head := &ListNode{-1, nil}
	carry := 0
	cnum := 0
	cur := head
	l1Val, l2Val := 0, 0
	for l1 != nil || l2 != nil {
		if l1 == nil {
			l1Val = 0
		} else {
			l1Val = l1.Val
			l1 = l1.Next
		}

		if l2 == nil {
			l2Val = 0
		} else {
			l2Val = l2.Val
			l2 = l2.Next
		}

		csum := l1Val + l2Val + carry
		carry = csum / 10
		cnum = csum % 10
		cur.Next = &ListNode{cnum, nil}
		cur = cur.Next
	}

	if carry != 0 {
		cur.Next = &ListNode{carry, nil}
		cur = cur.Next
	}
	return head.Next
}

/*
给出两个 非空 的链表用来表示两个非负的整数。
其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。

如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。

您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

示例：

输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/add-two-numbers
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
