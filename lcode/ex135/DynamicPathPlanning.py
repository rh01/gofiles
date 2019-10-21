#!/bin/python
# -*- coding: utf8 -*-
import sys
import os
import re

#请完成下面这个函数，实现题目要求的功能
#当然，你也可以不按照下面这个模板来作答，完全按照自己的想法来 ^-^ 
#******************************开始写代码******************************


def  DynamicPathPlanning(matrixGrid):
    rows = len(matrixGrid)
    cols = len(matrixGrid[0])
    # initialize dp 
    dp = [[0 for _ in range(cols)]for _ in range(rows)]
    if matrixGrid[0][0] == "1" or matrixGrid[rows-1][cols-1]: return 0
    for i in range(rows):
        if matrixGrid[i][0] != "1":
            dp[i][0] = 1
        else:
            for i in range(i, rows, 1):
                dp[i][0] = 0
            break
    for i in range(cols):
        if matrixGrid[0][i] != "1":
            dp[0][i] = 1
        else:
            for i in range(i, cols, 1):
                dp[0][i] = 0
            break

    # compute dp
    for i in range(1, rows):
        for j in range(1, cols):
            if matrixGrid[i][j] == "1":
                dp[i][j] = 0
            else:
                dp[i][j] = dp[i-1][j] + dp[i][j-1] 
    return dp[rows-1][cols-1]

def isVaild(position):
    x, y = position
    if x >= _matrixGrid_rows or x < 0  or  y >- _matrixGrid_cols or y  < 0:
        return False
    return True


#******************************结束写代码******************************


_matrixGrid_rows = 0
_matrixGrid_cols = 0
_matrixGrid_rows = int(input())
_matrixGrid_cols = int(input())

_matrixGrid = []
for _matrixGrid_i in range(_matrixGrid_rows):
    _matrixGrid_temp = list(map(int,re.split(r'\s+', input().strip())))
    _matrixGrid.append(_matrixGrid_temp)

  
res = DynamicPathPlanning(_matrixGrid)

print(str(res) + "\n")