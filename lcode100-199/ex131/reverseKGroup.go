// Copyright 2019 The darwin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ex131

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

// 方法1：递归法
func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{-1, nil}
	dummy.Next = head
	pre, tail := dummy, dummy

	for {
		count := k
		// 找到tail节点的位置
		for count != 0 && tail != nil {
			count--
			tail = tail.Next
		}
		// 加入tail为nil，表示k没有结束就到表尾了，
		// 就不需要反转了
		if tail == nil {
			break
		}
		head = pre.Next
		// <> -> 1 -> 2-> 3
		// <> -> 2 -> 3 -> 1
		for pre.Next != tail {
			// 获取下一个元素
			cur := pre.Next
			// pre与cur.next连接起来,此时cur(孤单)掉了出来
			pre.Next = cur.Next
			cur.Next = tail.Next
			tail.Next = cur
		}
		pre = head
		tail = head
	}
	return dummy.Next
}

/*
class Solution:
    def reverseKGroup(self, head: ListNode, k: int) -> ListNode:
        dummy = ListNode(0)
        dummy.next = head
        pre = dummy
        tail = dummy
        while True:
            count = k
            while count and tail:
                count -= 1
                tail = tail.next
            if not tail: break
            head = pre.next
            while pre.next != tail:
                cur = pre.next # 获取下一个元素
                # pre与cur.next连接起来,此时cur(孤单)掉了出来
                pre.next = cur.next
                cur.next = tail.next # 和剩余的链表连接起来
                tail.next = cur #插在tail后面
            # 改变 pre tail 的值
            pre = head
            tail = head
        return dummy.next

*/
/*
给你一个链表，每 k 个节点一组进行翻转，请你返回翻转后的链表。

k 是一个正整数，它的值小于或等于链表的长度。

如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。

示例 :

给定这个链表：1->2->3->4->5

当 k = 2 时，应当返回: 2->1->4->3->5

当 k = 3 时，应当返回: 3->2->1->4->5

说明 :

你的算法只能使用常数的额外空间。
你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/reverse-nodes-in-k-group
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
