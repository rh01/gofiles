/*
 * Copyright 2019 @rh01 https://github.com/rh01
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

/**
You are given two non-empty linked lists representing two non-negative integers.
The digits are stored in reverse order and each of their nodes contain a single digit.
Add the two numbers and return it as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Example:

Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
Explanation: 342 + 465 = 807.
*/

package leetcode002

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	head := new(ListNode)
	cur := head
	for l1 != nil || l2 != nil || carry != 0 {
		n1, n2 := 0, 0
		if l1 != nil {
			n1, l1 = l1.Val, l1.Next
		}
		if l2 != nil {
			n2, l2 = l2.Val, l2.Next
		}
		num := n1 + n2 + carry
		carry = num / 10
		cur.Next = &ListNode{Val:num%10, Next:nil}
		cur = cur.Next
	}
	return head.Next
}


/** 常规做 */
func addTwoNumbersV1(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{} // 初始化新的头结点
	cur := res         // 这个作为遍历使用
	carry := 0         // 进位

	h1 := l1
	h2 := l2
	for ; h1 != nil && h2 != nil; {
		v1 := h1.Val
		v2 := h2.Val

		curV := v1 + v2 + carry
		carry = curV / 10
		curV = curV % 10
		cur.Next = &ListNode{Val: curV}

		// 遍历
		h1 = h1.Next
		h2 = h2.Next
		cur = cur.Next
	}

	for ; h1 != nil; {
		v1 := h1.Val
		curV := v1 + carry
		carry = curV / 10
		curV = curV % 10
		cur.Next = &ListNode{Val: curV}

		h1 = h1.Next
		cur = cur.Next
	}

	for ; h2 != nil; {
		v1 := h2.Val
		curV := v1 + carry
		carry = curV / 10
		curV = curV % 10
		cur.Next = &ListNode{Val: curV}

		h2 = h2.Next
		cur = cur.Next
	}

	if carry != 0 {
		cur.Next = &ListNode{Val: carry}
	}

	return res.Next
}
