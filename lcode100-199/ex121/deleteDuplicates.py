# Definition for singly-linked list.
class ListNode(object):
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution(object):
    def deleteDuplicates(self, head):
        """
        :type head: ListNode
        :rtype: ListNode
        """
        
        deleted = []
        added = []
        cur = head
        while cur != None:
            value = cur.val
            if value not in deleted:
                if value not in added:
                    added.append(value)
                else:
                    added.remove(value)
                    deleted.append(value)
                cur = cur.next
            else:
                cur = cur.next
                continue
            
        
        cur = ListNode(-1)
        newHead = cur
        for i in added:
            newHead.next = ListNode(i)
            newHead = newHead.next
        return cur.next
        
        

"""
 1->2->3->3->4->4->5
 给定一个排序链表，删除所有含有重复数字的节点，只保留原始链表中 没有重复出现 的数字。
"""