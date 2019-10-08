#!/usr/bin/env python
# -*- coding:utf-8 -*-

'''
Filename: deleteDuplication
Created Date: 2019/10/8 9:58
Author: Shine

Copyright (c) 2019 41sh.cn
'''


# -*- coding:utf-8 -*-
class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution:
    def deleteDuplication(self, pHead):
        cur = pHead
        hmap = []
        removed = []
        while cur != None:
            if cur.val not in hmap:
                if cur.val not in removed:
                    hmap.append(cur.val)
            else:
                index = hmap.index(cur.val)
                pval = hmap.pop(index)
                removed.append(pval)
            cur = cur.next

        newHead = ListNode(0)
        cur = newHead
        for i in hmap:
            cur.next = ListNode(i)
            cur = cur.next
        return newHead.next
