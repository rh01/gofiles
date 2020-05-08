/**
 * Definition for singly-linked list.
 */
type ListNode struct {
     Val int
     Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	newHead := &ListNode{-1, nil}
	cur := newHead
	for l1 != nil || l2 != nil {
		if l1 == nil {
			cur.Next = l2
			break
		}

		if l2 == nil {
			cur.Next = l1
			break
		}

		if l1.Val < l2.Val {
			l1 = l1.Next
			cur.Next = &ListNode{l1.Val, nil}
		} else {
			l2 = l2.Next
			cur.Next = &ListNode{l2.Val, nil}
		}
		cur = cur.Next
	}

	return newHead.Next
}