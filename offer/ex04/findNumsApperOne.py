#!/usr/bin/env python
# -*- coding:utf-8 -*-

'''
Filename: findNumsApperOne
Created Date: 2019/10/7 16:38
Author: Shine

Copyright (c) 2019 41sh.cn
'''

from operator import xor


# -*- coding:utf-8 -*-
class Solution:
    # 返回[a,b] 其中ab是出现一次的两个数字
    def FindNumsAppearOnce(self, array):
        # write code here
        # 这里使用异或的性质来做
        res = 0
        # 此时的res一定是两个不同值得抑或结果
        for i in array:
            res = xor(res, i)

        # 找到最低位为1得位置，使得成为这样00001000
        res = (((~res) + 1) & res)
        num1, num2 = 0, 0
        for i in array:
            if res & i:
                num1 = xor(num1, i)
            else:
                num2 = xor(num2, i)
        return [num1, num2]
