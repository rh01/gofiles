// Copyright 2019 readailib.com. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

// 带环链表的入口节点
/*
# -*- coding:utf-8 -*-
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None
class Solution:
    def EntryNodeOfLoop(self, pHead):
        # write code here

这个需要多重复几遍
*/

// 定义链表
type ListNode struct {
	val  int
	next *ListNode
}

// 给一个链表，若其中包含环，请找出该链表的环的入口结点，
// 否则，输出null。
func EntryNodeOfLoop(pHead *ListNode) *ListNode {

	//这里常见的方法就是快慢指针
	fast := pHead
	slow := pHead

	// 两个快慢指针
	// 这里必须处理异常TODO
	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
		if slow == fast {
			break
		}
	}

	if fast == nil || fast.next == nil {
		return nil
	}

	fast = pHead
	for fast != slow {
		fast = fast.next
		slow = slow.next
	}
	return fast
}

func main() {
}
