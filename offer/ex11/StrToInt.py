#!/usr/bin/env python
# -*- coding:utf-8 -*-

'''
Filename: StrToInt
Created Date: 2019/10/8 8:05
Author: Shine

Copyright (c) 2019 41sh.cn
'''


# -*- coding:utf-8 -*-
class Solution:
    def StrToInt(self, s):
        res = 0
        isPostive = True
        for i in s:
            if i == "+":
                isPostive = True
                continue
            if i == "-":
                isPostive = False
                continue
            val = ord(i) - ord("0")
            if 0 <= val <= 9:
                res = res * 10 + val
            else:
                return 0
        if not isPostive:
            return -res
        else:
            return res


# write code here

sol = Solution()
print(sol.StrToInt("+2147483647"))
