#!/usr/bin/env python
# -*- coding:utf-8 -*-

'''
Filename: movingCount
Created Date: 2019/10/9 9:04
Author: Shine

地上有一个m行和n列的方格。一个机器人从坐标0,0的格子开始移动，
每一次只能向左，右，上，下四个方向移动一格，
但是不能进入行坐标和列坐标的数位之和大于k的格子。
例如，当k为18时，机器人能够进入方格（35,37），
因为3+5+3+7 = 18。但是，它不能进入方格（35,38），
因为3+5+3+8 = 19。请问该机器人能够达到多少个格子？

Copyright (c) 2019 41sh.cn
'''


# -*- coding:utf-8 -*-
class Solution:
    def __init__(self):
        self.count = 0

    def movingCount(self, threshold, rows, cols):
        # 初始化arr每个元素为一，然后访问每个元素
        # 并将其置为0
        arr = [[1 for _ in range(cols)] for _ in range(rows)]
        self.walk(threshold, rows, cols, arr, 0, 0)
        return self.count

    def walk(self, threshold, rows, cols, arr, i, j):
        # 判断是否满足递归结束的条件
        if i >= rows or j >= cols or i < 0 or j < 0 or arr[i][j] != 1:
            return

        # 下面计算数位之和
        tempi = sum(map(int, list(str(i))))
        tempj = sum(map(int, list(str(j))))

        # 下面判断数位之和是否大于threshold
        # 如果大于threshold，那么就会返回
        if tempi + tempj > threshold:
            return

        # 上面都满足的话，说明可以进入这个格子(i, j)
        # 接下来将arr[i][j]置为0即可
        arr[i][j] = 0
        self.count += 1

        # 接下来递归的返回上下左右即可
        self.walk(threshold, rows, cols, arr, i - 1, j)
        self.walk(threshold, rows, cols, arr, i + 1, j)
        self.walk(threshold, rows, cols, arr, i, j - 1)
        self.walk(threshold, rows, cols, arr, i, j + 1)
