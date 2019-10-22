# Definition for a binary tree node.
class TreeNode(object):
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Solution(object):
    def isValidBST(self, root):
        """
        :type root: TreeNode
        :rtype: bool
        """
        res = []
        
        self.preOrder(root, res)
        if len(res) < 2:
            return True
        print(res)
        for i in range(1, len(res)):
            if res[i] < res[i-1]:
                return False
        return True

    def preOrder(self, root, res):
        if root is None:
            return
        self.preOrder(root.left, res)
        res.append(root.val)

        self.preOrder(root.right, res)

sol = Solution()
root = TreeNode(2)
root.left = TreeNode(1)
root.right = TreeNode(3)

print(sol.isValidBST(root))