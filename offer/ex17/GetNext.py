#!/usr/bin/env python
#-*- coding:utf-8 -*-

'''
Filename: GetNext
Created Date: 2019/10/8 14:54
Author: Shine

Copyright (c) 2019 41sh.cn
'''

# -*- coding:utf-8 -*-
# class TreeLinkNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None
#         self.next = None
class Solution:
    def GetNext(self, pNode):
        # write code here
        if pNode is None:
            return None

        if pNode.right != None:
            cur = pNode.right
            while cur.left != None:
                cur = cur.left
            return cur
        else:
            cur = pNode
            p = pNode.next
            while p != None and p.right==cur:
                p = p.next
                cur = cur.next
            return p

