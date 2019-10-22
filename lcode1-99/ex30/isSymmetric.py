# Definition for a binary tree node.
class TreeNode(object):
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Solution(object):
    def isSymmetric(self, root):
        """
        :type root: TreeNode
        :rtype: bool
        """
        if root is None:
            return True
        return self.isChildSymmetric(root.left, root.right)
        
    def isChildSymmetric(self, left, right):
        if left is None:
            return right == None

        # 这个就多余了
        if right is None:
            return left == None

        return left.val == right.val and\
             self.isChildSymmetric(left.right, right.left) and \
             self.isChildSymmetric(left.left, right.right)

        
