"""
给定二叉树的根节点 root，找出存在于不同节点 A 和 B 之间的最大值 V，
其中 V = |A.val - B.val|，且 A 是 B 的祖先。

（如果 A 的任何子节点之一为 B，或者 A 的任何子节点是 B 的祖先，
那么我们认为 A 是 B 的祖先）

 

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/maximum-difference-between-node-and-ancestor
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
"""
# Definition for a binary tree node.
class TreeNode(object):
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

import sys

class Solution(object):
    def __init__(self):
        self.maxmin = 0
    def maxAncestorDiff(self, root, sum):
        """
        :type root: TreeNode
        :type sum: int
        :rtype: List[List[int]]
        """
        if root is None:
            return []
        res = []
        self.walk(root, sys.maxsize, -sys.maxsize, 0, res)
        return max(res)
    

    def walk(self, root, minnum, maxnum):
        # 递归终止条件
        if root == None:
            return
        
        if root.val > maxnum:
            maxnum = root.val
        if root.val < minnum:
            minnum = root.val

        if abs(maxnum - minnum) > self.maxmin:
            self.maxmin = abs(maxnum - minnum)

        if root.left == None and root.right == None:
            return

        self.walk(root.left, minnum, maxnum)
        self.walk(root.left, minnum, maxnum)
        return