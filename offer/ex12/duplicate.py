#!/usr/bin/env python
# -*- coding:utf-8 -*-

'''
Filename: duplicate
Created Date: 2019/10/8 8:38
Author: Shine

Copyright (c) 2019 41sh.cn
'''


# -*- coding:utf-8 -*-
class Solution:
    # 这里要特别注意~找到任意重复的一个值并赋值到duplication[0]
    # 函数返回True/False
    def duplicatev1(self, numbers, duplication):
        # write code here
        hlist = []
        for i in numbers:
            if i not in hlist:
                hlist.append(i)
            else:
                duplication[0] = i
                break
        if duplication[0] == -1:
            return False
        else:
            return True

    def duplicate(self, numbers, duplication):
        length = len(numbers)
        for i in range(length):
            index = numbers[i]
            if index >= length:
                index = index - length

            if numbers[index] >= length:
                duplication[0] = index
                return True
            else:
                numbers[index] += length
        return False


sol = Solution()
print(sol.duplicate([2, 4, 3, 1, 4], [-1]))
