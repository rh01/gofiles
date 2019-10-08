#!/usr/bin/env python
# -*- coding:utf-8 -*-

'''
Filename: IsContinuous
Created Date: 2019/10/7 19:56
Author: Shine

Copyright (c) 2019 41sh.cn
'''


# -*- coding:utf-8 -*-
class Solution:
    def IsContinuous(self, numbers):
        if len(numbers) != 5:
            return False
        minn, maxn = 14, -1
        for i in numbers:
            if i == 0:
                continue
            if numbers.count(i) > 1:
                return False
            if i > maxn:
                maxn = i

            if i < minn:
                minn = i

        if maxn - minn >= 5:
            return False
        return True
