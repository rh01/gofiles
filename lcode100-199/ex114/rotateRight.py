# Definition for singly-linked list.
class ListNode(object):
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution(object):
    def rotateRight(self, head, k):
        """
        :type head: ListNode
        :type k: int
        :rtype: ListNode
        """
        if head is None:
            return None
        nums = []
        length = 0
        cur = head
        while cur != None:
            nums.append(cur.val)
            cur = cur.next
            length+=1
        
        k %= length
        nums.reverse()
        def reverse(nums, start, end):
            while start < end:
                nums[start], nums[end] = nums[end], nums[start]
                start+=1
                end -=1
        reverse(nums, 0, k-1)
        reverse(nums, k, length-1)
        
        newHead = ListNode(-1)
        for i in res:
            newHead.next = ListNode(i)
            newHead = newHead.next
        
        return newHead.next
        