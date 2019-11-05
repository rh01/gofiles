// Copyright 2019 The darwin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ex157

import "fmt"

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func partitionv1(head *ListNode, x int) *ListNode {

	newh := &ListNode{-1, nil}
	newh.Next = head

	l, r1 := newh, newh
	for r1.Next != nil {
		r1 = r1.Next
	}
	fmt.Println(r1.Val)

	for l.Next != nil && l.Next != r1 {
		curVal := l.Next.Val
		if curVal >= x {
			tmp := r1.Next
			r1.Next = &ListNode{curVal, tmp}
			l.Next = l.Next.Next
		} else {
			l = l.Next
		}
	}
	return newh.Next
}

func partition(head *ListNode, x int) *ListNode {

	newh := &ListNode{-1, nil}
	newh.Next = head

	l, r1 := newh, newh
	for r1.Next != nil {
		r1 = r1.Next
	}
	r := r1
	for l.Next != nil && l.Next != r1.Next {
		curVal := l.Next.Val
		if curVal >= x {
			r.Next = &ListNode{curVal, nil}
			r = r.Next
			l.Next = l.Next.Next
		} else {
			l = l.Next
		}
	}
	return newh.Next
}
