"""
给定一个数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。

candidates 中的每个数字在每个组合中只能使用一次。

说明：

所有数字（包括目标数）都是正整数。
解集不能包含重复的组合。 
示例 1:

输入: candidates = [10,1,2,7,6,1,5], target = 8,
所求解集为:
[
  [1, 7],
  [1, 2, 5],
  [2, 6],
  [1, 1, 6]
]
示例 2:

输入: candidates = [2,5,2,1,2], target = 5,
所求解集为:
[
  [1,2,2],
  [5]
]
"""

class Solution(object):
    def combinationSum2(self, candidates, target):
        """
        :type candidates: List[int]
        :type target: int
        :rtype: List[List[int]]
        """
        res = []
        self.walk(res, candidates, [], target)
        return res
        

    def walk(self, res, candidates, intnums, target):
      if sum(intnums) > target:
        return

      if sum(intnums) == target:
          intnums.sort()
          if intnums not in res:
              res.append(intnums)
          return

      for i, c in enumerate(candidates):
          #if c not in intnums:
            self.walk(res, candidates[:i]+candidates[i+1:], intnums+[c], target)
      return

sol = Solution()
print(sol.combinationSum2([5,24,28,14,13,28,12,29,22,8,16,28,11,5,8,20,10,27,16,19,16,15,14,14,9,23,30,13,33,24,24,33,14,18,5,14,33,12,30,21,15,12,14,13,34,9,20,9,31,32,16],29))