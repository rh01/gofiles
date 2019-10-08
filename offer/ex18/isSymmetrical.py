#!/usr/bin/env python
# -*- coding:utf-8 -*-

'''
Filename: isSymmetrical
Created Date: 2019/10/8 14:56
Author: Shine

Copyright (c) 2019 41sh.cn
'''


# -*- coding:utf-8 -*-
class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


class Solution:
    def isSymmetrical(self, pRoot):
        # write code here
        if pRoot is None:
            return True
        return self.left2right(pRoot.left, pRoot.right)

    def left2right(self, left, right):
        if left is None:
            return right == None

        if right is None:
            return left == None
        if left.val != right.val:
            return False

        return self.left2right(left.left, right.right) and \
               self.left2right(left.right, right.left)
