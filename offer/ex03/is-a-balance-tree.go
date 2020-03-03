package main

//输入一棵二叉树，判断该二叉树是否是平衡二叉树。


/*
# -*- coding:utf-8 -*-
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None
class Solution:
    def IsBalanced_Solution(self, pRoot):
        # write code here
*/

// 输入一棵二叉树，求该树的深度。从根结点到叶结点依次经过的结点
// （含根、叶结点）形成树的一条路径，最长路径的长度为树的深度。
type TreeNode struct {
	value int       // 当前节点的值
	left  *TreeNode // 左孩子节点
	right *TreeNode //右孩子节点
}

func IsBalanced_Solution(pRoot *TreeNode) bool {
	if pRoot == nil {
		return true
	}
	return (abs(maxTreeDepth(pRoot.right)-maxTreeDepth(pRoot.left)) <= 1) &&
		IsBalanced_Solution(pRoot.left) && IsBalanced_Solution(pRoot.right)
}

// 返回书的高度
func maxTreeDepth(pRoot *TreeNode) int {
	if pRoot == nil {
		return 0
	}
	return max(maxTreeDepth(pRoot.left), maxTreeDepth(pRoot.right)) + 1
}

func abs(res int) int {

	if res < 0 {
		return -res
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
