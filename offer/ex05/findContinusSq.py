#!/usr/bin/env python
#-*- coding:utf-8 -*-

'''
Filename: findContinusSq
Created Date: 2019/10/7 17:26
Author: Shine

Copyright (c) 2019 41sh.cn
'''

class Solution:
    def FindContinuousSequence(self, tsum):
        # write code here
        # 记住这里是1,2,3,...,n得序列
        res = []
        pLow, pHigh = 1, 2
        while (pLow < pHigh):
            # 等差数列公式
            csum = (pLow + pHigh)*(pHigh - pLow + 1) / 2
            if csum == tsum:
                res.append(list(range(pLow, pHigh+1)))
                pLow+=1
            elif csum < tsum:
                pHigh+=1
            else:
                pLow+=1
        return res

s = Solution()
s.FindContinuousSequence(100)