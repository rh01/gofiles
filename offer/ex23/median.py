#!/usr/bin/env python
# -*- coding:utf-8 -*-


# 如何得到一个数据流中的中位数？如果从数据流中读出奇数个数值，
# 那么中位数就是所有数值排序之后位于中间的数值。如果从数据流中读出偶数个数值，
# 那么中位数就是所有数值排序之后中间两个数的平均值。
# 我们使用Insert()方法读取数据流，使用GetMedian()方法获取当前读取数据的中位数。

# -*- coding:utf-8 -*-
from heapq import *


class Solution:
    def __init__(self):
        self.heaps = [], []

    def Insert(self, num):
        # write code here
        small, large = self.heaps
        heappush(small, -heappushpop(large, num))  # 将num放入大根堆，并弹出大根堆的最小值，取反，放入大根堆small
        if len(large) < len(small):
            heappush(large, -heappop(small))  # 弹出small中最小的值，取反，即最大的值，放入large

    def GetMedian(self, ss):
        # write code here
        small, large = self.heaps
        if len(large) > len(small):
            return float(large[0])
        return (large[0] - small[0]) / 2.0
