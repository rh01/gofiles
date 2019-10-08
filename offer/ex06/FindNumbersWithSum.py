#!/usr/bin/env python
#-*- coding:utf-8 -*-

'''
Filename: FindNumbersWithSum
Created Date: 2019/10/7 18:13
Author: Shine

Copyright (c) 2019 41sh.cn
'''

# -*- coding:utf-8 -*-
import sys
class Solution:
    def FindNumbersWithSum(self, array, tsum):
        # write code here
        low, high = 0, len(array)-1
        minVal = sys.maxsize
        res = []

        while low < high:
            csum = array[low] + array[high]
            if csum == tsum:
                curMul = array[low] * array[high]
                if curMul < minVal:
                    minVal = curMul
                    res =[array[low],array[high]]
                low+=1
                # high=low+1
            elif csum < tsum:
                low+=1
            else:
                high-=1
        return res

    def FindNumbersWithSumv1(self, array, tsum):
        # write code here
        low, high = 0, 1
        minsum = sys.maxsize
        res = []

        hmap = {}
        for idx, v in enumerate(array):
            # t = tsum - v
            if v not in hmap:
                hmap[v] = idx
            t = tsum -v
            if t in hmap and t*v < minsum:
                minsum = t*v
                res =[t, v]
        return res

sol = Solution()
print(sol.FindNumbersWithSum([1, 2, 4, 7, 11, 15], 15))