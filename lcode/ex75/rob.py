# Definition for a binary tree node.
class TreeNode(object):
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Solution(object):
    def rob(self, root):
        """
        :type root: TreeNode
        :rtype: int
        """

        if root is None:
            return 0
        


        # a表示带有root节点 b表示不带有root节点
        def dfs(root):
            if root is None:
                return 0, 0
            
            la, lb = dfs(root.left)
            ra, rb = dfs(root.right)

            a = root.val + lb + rb
            b = max(la, lb) + max(ra, rb)

            return a, b
        a, b = dfs(root)
        return max(a, b)
