/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func reverseKGroup(head *ListNode, k int) *ListNode {
	 // fake 
	if head == nil {
		return nil
	}

	// ddummy node 
	dummy := *ListNode{-1, nil}
	dummy.Next = head
	pre, tail = dummy, dummy
	for {
		count := k
		// 找到tail节点的位置
		for count != 0 && tail != nil {
			tail = tail.Next
			count--
		}

		// 如果tail为nil,表示k没结束就到表尾， 此时不需要反转
		if tail == nil {
			break
		}

		head = pre.Next
		// 尾插法
		for pre.Next != tail {
			cur := pre.Next
			pre.Next = cur.Next
			cur.Next = tail.Next
			tail.Next = cur
		}
	}
	return dummy.Next
}