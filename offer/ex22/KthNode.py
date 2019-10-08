#!/usr/bin/env python
#-*- coding:utf-8 -*-


# -*- coding:utf-8 -*-
class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

# 给定一棵二叉搜索树，请找出其中的第k小的结点。
# 例如， （5，3，7，2，4，6，8） 中，
# 按结点数值大小顺序第三小结点的值为4。
class Solution:
    def __init__(self):
        self.res = []
    # 返回对应节点TreeNode
    def KthNode(self, pRoot, k):
        
        return TreeNode(self.res[k-1])

    def inOrder(self, pRoot):
        if pRoot is None:
            return
        self.inOrder(pRoot.left)
        self.res.append(pRoot.val)
        self.inOrder(pRoot.right)
