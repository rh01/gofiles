package main

/*
# -*- coding:utf-8 -*-
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None
class Solution:
    def TreeDepth(self, pRoot):
        # write code here
*/


// 输入一棵二叉树，求该树的深度。从根结点到叶结点依次经过的结点（含根、叶结点）形成树的一条路径，最长路径的长度为树的深度。
type TreeNode struct {
	value int       // 当前节点的值
	left  *TreeNode // 左孩子节点
	right *TreeNode //右孩子节点
}

//func (t *TreeNode) NewTreeNode()  {
//
//}

func TreeDepth(pRoot *TreeNode) int {
	if pRoot == nil {
		return 0
	}
	return max(TreeDepth(pRoot.left), TreeDepth(pRoot.right)) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
}
