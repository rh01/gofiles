# Definition for a binary tree node.
# 给定两个二叉树，编写一个函数来检验它们是否相同。

# 如果两个树在结构上相同，并且节点具有相同的值，则认为它们是相同的。

# 示例 1:

# 输入:       1         1
#           / \       / \
#          2   3     2   3

#         [1,2,3],   [1,2,3]

# 输出: true
# 示例 2:

# 输入:      1          1
#           /           \
#          2             2

#         [1,2],     [1,null,2]

# 输出: false
# 示例 3:

# 输入:       1         1
#           / \       / \
#          2   1     1   2

#         [1,2,1],   [1,1,2]

# 输出: false

# 来源：力扣（LeetCode）
# 链接：https://leetcode-cn.com/problems/same-tree
# 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
class TreeNode(object):
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Solution(object):
    def isSameTree(self, p, q):
        """
        :type p: TreeNode
        :type q: TreeNode
        :rtype: bool
        """
        if p is None:
            return q == None
        
        if q is None:
            return p == None
             
        return p.val == q.val and self.isSameTree(p.left, q.left) and self.isSameTree(p.right, q.right)
        