"""
给定一个单链表 L：L0→L1→…→Ln-1→Ln ，
将其重新排列后变为： L0→Ln→L1→Ln-1→L2→Ln-2→…

你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

示例 1:

给定链表 1->2->3->4, 重新排列为 1->4->2->3.
示例 2:

给定链表 1->2->3->4->5, 重新排列为 1->5->2->4->3.
"""

# Definition for singly-linked list.
class ListNode(object):
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution(object):
    def reorderList(self, head):
        """
        :type head: ListNode
        :rtype: None Do not return anything, modify head in-place instead.
        """
        
        if head == None or head.next == None or head.next.next == None:
            return 
        
        slow, fast = head, head
        while fast.next != None and fast.next.next != None:
            slow = slow.next
            fast = fast.next.next
        
        # 查到中间的链表节点，然后将两个连起来
        head2 = slow.next
        slow.next = None
        
        # 反转链表
        pre = None
        nextNode = None
        cur = head2
        while cur != None:
            nextNode = cur.next
            cur.next = pre
            pre = cur 
            cur = nextNode
        # 此时pre节点为头节点
        head2 = pre
        
        cur1 = head
        cur2 = head2
        
        while cur1 != None and cur2 != None:
            next1Node = cur1.next
            next2Node = cur2.next
            
            cur2.next = next1Node
            cur1.next = cur2
            
            cur1 = next1Node
            cur2 = next2Node
            
        return head
            
        
        
        
        