"""
给定一个单链表，其中的元素按升序排序，将其转换为高度平衡的二叉搜索树。

本题中，一个高度平衡二叉树是指一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1。

示例:

给定的有序链表： [-10, -3, 0, 5, 9],

一个可能的答案是：[0, -3, 9, -10, null, 5], 它可以表示下面这个高度平衡二叉搜索树：

      0
     / \
   -3   9
   /   /
 -10  5

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/convert-sorted-list-to-binary-search-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
"""
# Definition for singly-linked list.
class ListNode(object):
    def __init__(self, x):
        self.val = x
        self.next = None

# Definition for a binary tree node.
class TreeNode(object):
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None
class Solution(object):
    def sortedListToBST(self, head):
        """
        :type head: ListNode
        :rtype: TreeNode
        """
        if head == None:
            return None
        if head.next == None:
            return TreeNode(head.val)
        slow, fast = head, head
        # while (fast != null && fast.next != null) {
        #     pre = pre.next;
        #     slow = slow.next;
        #     fast = fast.next.next;
        # }   
        # ListNode pre = head, slow = head.next, fast = head.next.next;
        while fast.next != None and  fast.next.next != None:
            slow, fast = slow.next, fast.next.next
        root = TreeNode(slow.val)
        rightHead = slow.next
        
        if head.val == root.val:
            head=None
        elif head.next.val == root.val:
            head.next = None
        else:
            cur = head
            while cur.next.val != root.val :
                cur = cur.next
            cur.next=None
        root.left = self.sortedListToBST(head)
        root.right = self.sortedListToBST(rightHead)
        return root