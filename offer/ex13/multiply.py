#!/usr/bin/env python
#-*- coding:utf-8 -*-

'''
Filename: multiply
Created Date: 2019/10/8 9:20
Author: Shine

Copyright (c) 2019 41sh.cn
'''

# -*- coding:utf-8 -*-
class Solution:
    def multiply(self, A):
        # write code here
        length = len(A)

        B = [0 for i in range(length)]

        # if i == 0:
        B[0] = self.fn(A[1:])
            # continue
        # elif i == length-1:
        B[length-1] = self.fn(A[:length-1])
            # continue
        for i in range(1,length-1):
           B[i] = self.fn(A[:i])*self.fn(A[i+1:])
        return B

    def fn(self, nums):
        res = 1
        for i in nums:
            res*=i
        return res

sol = Solution()
print(sol.multiply([1, 2, 3, 4, 5]))