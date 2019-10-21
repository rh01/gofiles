// Copyright 2019 readailib.com. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// 给定一个链表: 1->2->3->4->5, 和 n = 2.
// 当删除了倒数第二个节点后，链表变为 1->2->3->5.
package main

import "fmt"

/**
 * Definition for singly-linked list.
 * */

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {

	if head == nil || n <= 0 {
		return head
	}

	length := 0
	cur := head
	for cur != nil {
		cur = cur.Next
		length++
	}
	print(length)
	// 如果是第一个则
	if n >= length {
		head = head.Next
		return head
	}

	cur = head
	for i := length; i != n+1; i = i - 1 {
		cur = cur.Next
	}
	cur.Next = cur.Next.Next
	return head
}

func main() {
	head := &ListNode{
		1, &ListNode{
			2, &ListNode{
				3, &ListNode{
					4, &ListNode{
						5, nil,
					},
				},
			},
		},
	}
	head2 := removeNthFromEnd(head, 1)
	cur := head2
	res := make([]int, 0)
	for cur != nil {
		res = append(res, cur.Val)
		cur = cur.Next
	}
	fmt.Println(res)

}

/*
给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。

示例：

给定一个链表: 1->2->3->4->5, 和 n = 2.

当删除了倒数第二个节点后，链表变为 1->2->3->5.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
