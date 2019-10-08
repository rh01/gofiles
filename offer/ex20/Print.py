#!/usr/bin/env python
# -*- coding:utf-8 -*-

'''
Filename: Print
Created Date: 2019/10/8 15:20
Author: Shine

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
    # 返回二维列表[[1,2],[4,5]]
    def Print(self, pRoot):
        # write code here
        res = []
        if pRoot is None:
            return []

        queue = []
        tempNodes = []
        queue.append(pRoot)
        res.append([pRoot.val])
        while queue:
            cur = queue.pop(0)

            data = []
            if cur.left != None:
                tempNodes.append(cur.left)
            if cur.right != None:
                tempNodes.append(cur.right)

            if len(queue) == 0:
                while tempNodes:
                    cur = tempNodes.pop(0)
                    curn = cur.val
                    data.append(curn)
                    queue.append(cur)
                res.append(data)
        res.pop()
        return res
