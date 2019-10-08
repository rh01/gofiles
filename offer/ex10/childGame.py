#!/usr/bin/env python
#-*- coding:utf-8 -*-

'''
Filename: childGame
Created Date: 2019/10/7 20:36
Author: Shine

Copyright (c) 2019 41sh.cn
'''



# -*- coding:utf-8 -*-
class Solution:
    def LastRemaining_Solution(self, n, m):
        if n < 1:
            return -1

        con = range(n)

        final = -1
        start = 0
        while con:
            k = (start + m - 1) % n
            final = con.pop(k)
            n -= 1
            start = k

        return final
