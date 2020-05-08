package main

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverse(l1 *ListNode) *ListNode {
	// 头插法
	head := &ListNode{-1, nil}
	cur := head

	for l1 != nil {
		temp := l1.Next
		l1.Next = cur.Next
		cur.Next = l1
		l1 = temp
	}
	return head.Next
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	newHead := &ListNode{-1, nil}
	acc := 0
	cur := newHead
	l1 = reverse(l1)
	l2 = reverse(l2)
	for l1 != nil || l2 != nil {
		var a, b int
		if l1 != nil {
			a = l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			b = l2.Val
			l2 = l2.Next
		}

		s := a + b + acc
		node := &ListNode{s % 10, nil}
		acc = s / 10
		cur.Next = node
		cur = cur.Next
	}
	return reverse(newHead.Next)
}

func main() {

}
