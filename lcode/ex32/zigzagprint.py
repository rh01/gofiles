# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution(object):
    def zigzagLevelOrder(self, pRoot):
        """
        :type root: TreeNode
        :rtype: List[List[int]]
        """
        if pRoot is None:
            return []
        queue = []
        help_stack = []
        res = []
        queue.append(pRoot)
        res.append([pRoot.val])
        flip = True
        q2 = []
        while queue:
            if flip:
                cur = queue.pop(0)
            else:
                cur = queue.pop()

            if cur.left != None:
                help_stack.append(cur.left)
            if cur.right != None:
                help_stack.append(cur.right)
            if len(queue) == 0:
                if flip:
                    while help_stack:
                        cur = help_stack.pop()
                        queue.append(cur)
                        q2.append(cur.val)
                    flip = False
                else:
                    while help_stack:
                        cur = help_stack.pop(0)
                        queue.append(cur)
                        q2.append(cur.val)
                    flip = True
                if len(q2) != 0:
                    res.append(q2)
                q2 = []
        return res
