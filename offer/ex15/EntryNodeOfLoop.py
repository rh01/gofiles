#!/usr/bin/env python
# -*- coding:utf-8 -*-

'''
Filename: EntryNodeOfLoop
Created Date: 2019/10/8 9:58
Author: Shine

Copyright (c) 2019 41sh.cn
'''


# -*- coding:utf-8 -*-
class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None


# -*- coding:utf-8 -*-
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None
class Solution:
    def EntryNodeOfLoop(self, pHead):

        pfast = pHead
        pslow = pHead

        while pfast!= None and pfast.next != None:
            pfast = pfast.next.next
            pslow = pslow.next
            if pfast is pslow:
                break
        if pfast is None or pfast.next is None:
            return None

        pslow = pHead
        while pfast != pslow:
            pfast = pfast.next
            pslow = pslow.next

        return pslow