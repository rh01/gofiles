#!/usr/bin/env python
#-*- coding:utf-8 -*-

'''
Filename: firstAppear
Created Date: 2019/10/8 9:56
Author: Shine

Copyright (c) 2019 41sh.cn
'''

class Solution:

    def __init__(self):
        self.s=""
    def FirstAppearingOnce(self):

        res=list(filter(lambda c:self.s.count(c)==1,self.s))
        return res[0] if res else "#"

    def Insert(self, char):

        self.s+=char