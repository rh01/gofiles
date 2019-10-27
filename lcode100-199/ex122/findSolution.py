"""
   This is the custom function interface.
   You should not implement it, or speculate about its implementation
   class CustomFunction:
       # Returns f(x, y) for any given positive integers x and y.
       # Note that f(x, y) is increasing with respect to both x and y.
       # i.e. f(x, y) < f(x + 1, y), f(x, y) < f(x, y + 1)
       def f(self, x, y):
  
"""

import math
class Solution(object):
    def findSolution(self, customfunction, z):
        """
        :type num: int
        :type z: int
        :rtype: List[List[int]]
        """
        res = []
        if customfunction == 1:
            for i in range(1, z//2):
                res.add([i, z-i])
                res.add([z-i, i])
            return res
        
        if customfunction == 2:
            for i in range(1, int(math.sqrt(z))):
                if z % i == 0:
                    res.append([z/i, i])
                    res.append([i, z])
            return res
        
                