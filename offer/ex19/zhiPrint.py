#!/usr/bin/env python
# -*- coding:utf-8 -*-

'''
Filename: zhiPrint
Created Date: 2019/10/8 15:10
Author: Shine

请实现一个函数按照之字形打印二叉树，
即第一行按照从左到右的顺序打印，
第二层按照从右至左的顺序打印，
第三行按照从左到右的顺序打印，
其他行以此类推。

Copyright (c) 2019 41sh.cn
'''


# -*- coding:utf-8 -*-
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None
# -*- coding:utf-8 -*-
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None
class Solution:
    def Print(self, pRoot):
        # write code here
        if pRoot is None:
            return []
        queue = []
        help_stack = []
        res = []
        queue.append(pRoot)
        res.append([pRoot.val])
        flip = True
        q2 = []
        while queue:
            if flip:
                cur = queue.pop(0)
            else:
                cur = queue.pop()

            if cur.left != None:
                help_stack.append(cur.left)
            if cur.right != None:
                help_stack.append(cur.right)
            if len(queue) == 0:
                if flip:
                    while help_stack:
                        cur = help_stack.pop()
                        queue.append(cur)
                        q2.append(cur.val)
                    flip = False
                else:
                    while help_stack:
                        cur = help_stack.pop(0)
                        queue.append(cur)
                        q2.append(cur.val)
                    flip = True
                res.append(q2)
                q2 = []
        res.pop()
        return res
