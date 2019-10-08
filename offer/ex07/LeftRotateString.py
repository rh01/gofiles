#!/usr/bin/env python
#-*- coding:utf-8 -*-

'''
Filename: LeftRotateString
Created Date: 2019/10/7 19:11
Author: Shine

Copyright (c) 2019 41sh.cn
'''


# -*- coding:utf-8 -*-
class Solution:
    def LeftRotateStringv1(self, s, n):
        return (s[:n][::-1]+s[n:][::-1])[::-1]



