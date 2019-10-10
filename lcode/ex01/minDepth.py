#!/usr/bin/env python
#-*- coding:utf-8 -*-

'''
Filename: minDepth
Created Date: 2019/10/9 9:59
Author: Shine

Copyright (c) 2019 41sh.cn
'''

# Definition for a binary tree node.
class TreeNode(object):
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Solution(object):
    def minDepth(self, root):
        """
        :type root: TreeNode
        :rtype: int
        """
        if root is None:
            return 0
        return self._find(root, 1)
    
    def _find(self, root, level):
        if root.left is None and root.right is None:
            return level
        elif root.left is None and root.right != None:
            return self._find(root.right, level)
        elif root.left != None and root.right == None:
            return self._find(root.left, level+1)
        else:
            return min(self._find(root.left, level+1), self._find(root.right, level+1))