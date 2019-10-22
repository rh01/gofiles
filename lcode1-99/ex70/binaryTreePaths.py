
# Definition for a binary tree node.
class TreeNode(object):
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


"""
给定一个二叉树，返回所有从根节点到叶子节点的路径。

说明: 叶子节点是指没有子节点的节点。

示例:

输入:

   1
 /   \
2     3
 \
  5

输出: ["1->2->5", "1->3"]

解释: 所有根节点到叶子节点的路径为: 1->2->5, 1->3

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/binary-tree-paths
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
"""

class Solution(object):
    def binaryTreePaths(self, root):
        """
        :type root: TreeNode
        :rtype: List[str]
        """

        if root is None:
            return []
        res = []
        self.dfs(root, "", res)
        return res

    def dfs(self, root, cur, res):
        # 判断终止条件
        if root == None:
            return
        
        cur+=str(root.val)
        if not (root.left == None and root.right == None):
            cur+="->"
        else:
            res.append(cur)
            return
        
        dfs(root.left, cur, res)
        dfs(root.right, cur, res)
        return