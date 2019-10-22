#!/usr/bin/env python
# -*- coding:utf-8 -*-

'''
Filename: maxPoints
Created Date: 2019/10/10 10:18
Author: Shine

Given n points on a 2D plane, find the maximum number
of points that lie on the same straight line.

Example 1:

Input: [[1,1],[2,2],[3,3]]
Output: 3
Explanation:
^
|
|        o
|     o
|  o
+------------->
0  1  2  3  4
Example 2:

Input: [[1,1],[3,2],[5,3],[4,1],[2,3],[1,4]]
Output: 4
Explanation:
^
|
|  o
|     o        o
|        o
|  o        o
+------------------->
0  1  2  3  4  5  6
NOTE: input types have been changed on April 15, 2019.
Please reset to default code definition to get new method signature.

Copyright (c) 2019 41sh.cn
'''
from collections import Counter


class Solution(object):
    def maxPoints(self, points):
        """
        :type points: List[List[int]]
        :rtype: int
        """
        # points: [[x1, y1], [x2, y2]] 二维数组
        # for i in points:
        #     if points.count(i) > 1:
        #         points.remove(i)
        if len(points) < 2:
            return len(points)
        maxnum = 0
        for i1, j1 in enumerate(points):
            once = {}
            vertical = 1
            overleap = 0
            oncemax = 0
            for i2, j2 in enumerate(points):
                if i1 == i2:
                    continue
                if j1 == j2:
                    overleap += 1
                    continue
                x1, y1 = j1
                x2, y2 = j2
                if x1 - x2 != 0:
                    slope = (y2 - y1) * 1.0 / (x2 - x1)
                    if slope not in once:
                        once[slope] = 2
                    else:
                        once[slope] += 1
                else:
                    vertical += 1
            if once != {}:
                # print(once)
                oncemax = once[max(once, key=lambda a: once[a])]
            maxnum = max(maxnum, oncemax + overleap, vertical + overleap)
        return maxnum

    def maxPointsv2(self, points):
        if not len(points):
            return 0
        # take care of repeated points
        points = [(i[0], i[1]) for i in points]
        repeated = Counter(points)
        points = list(repeated.keys())
        if len(points) < 3:
            return sum(repeated.values())
        line = {}

        def slant(p1, p2):
            if p2[1] == p1[1]:
                return None
            return (p2[0] - p1[0]) / (p2[1] - p1[1])

        mxt = -float('inf')

        def update(p1, p2, line=line):
            ind = (p1, slant(p1, p2))
            if ind not in line:
                # init line
                line[ind] = repeated[p1]
            line[ind] += repeated[p2]
            return line[ind]

        ## n square time enumerate lines
        for i, p1 in enumerate(points):
            for p2 in points[i + 1:]:
                mxt = max(mxt, update(p1, p2))

        return mxt


sol = Solution()
print(sol.maxPoints([[0, 0]]))
print(sol.maxPoints([[1, 1], [2, 2], [3, 3]]))
