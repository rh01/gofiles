// Copyright 2019 readailib.com. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

/*
# -*- coding:utf-8 -*-
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None
class Solution:
    def deleteDuplication(self, pHead):
        # write code here
*/

/*
在一个排序的链表中，存在重复的结点，
请删除该链表中重复的结点，
重复的结点不保留，返回链表头指针。
例如，链表1->2->3->3->4->4->5 处理后为 1->2->5*/

// listNode数据结构定义
type ListNode struct {
	val  int
	next *ListNode
}

// 删除重复的节点
// 返回链表头指针
func deleteDeplicationv1(pHead *ListNode) *ListNode {
	hmap := make(map[int]bool)
	cur := pHead

	for cur != nil {
		_, exist := hmap[cur.val]
		if !exist {
			hmap[cur.val] = true
		}
	}

	newHead := &ListNode{}
	cur = newHead
	for value, _ := range hmap {
		cur.next = &ListNode{val: value}
		cur = cur.next
	}
	return newHead.next
}

func main() {
}
