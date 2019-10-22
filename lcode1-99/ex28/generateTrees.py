# Definition for a binary tree node.
class TreeNode(object):
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Solution(object):
    def generateTrees(self, n):
        """
        :type n: int
        :rtype: List[TreeNode]
        """
        if n == 0: return []
        return self.createBSTs(1, n)
    
    def createBSTs(self, start, end):
        res = []
        if start > end:
            res.append(None)
            return res
        
        for i in range(start, end+1, 1):
            left_res = self.createBSTs(start, i-1)
            right_res  = self.createBSTs(i+1, end)

            lsize = len(left_res)
            rsize = len(right_res)

            for j in range(lsize):
                for k in range(rsize):
                    root = TreeNode(i)
                    root.left = left_res[j]
                    root.right = right_res[k]
                    res.append(root)
        return res