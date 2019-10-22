# Definition for a binary tree node.
class TreeNode(object):
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


class Solution(object):
    def pathSum(self, root, sum):
        """
        :type root: TreeNode
        :type sum: int
        :rtype: List[List[int]]
        """
        if root is None:
            return []
        res = []
        self.walk(res, root, [], sum)
        return res
    

    def walk(self, res, root, curres, totalsum):

        # 递归终止停止
        if root == None:
            return
        # 判断递归的停止条件
        if root.left == None and root.right == None:
            if  sum(curres) + root.val == totalsum:
                curres.append(root.val)
                res.append(curres)
            return
        
        self.walk(res, root.left, curres+[root.val], totalsum)
        self.walk(res, root.right, curres+[root.val], totalsum)
        return