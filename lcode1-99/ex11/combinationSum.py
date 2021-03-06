
"""
给定一个无重复元素的数组 candidates 和一个目标数 target ，
找出 candidates 中所有可以使数字和为 target 的组合。

candidates 中的数字可以无限制重复被选取。

说明：

所有数字（包括 target）都是正整数。
解集不能包含重复的组合。 
示例 1:

输入: candidates = [2,3,6,7], target = 7,
所求解集为:
[
  [7],
  [2,2,3]
]
示例 2:

输入: candidates = [2,3,5], target = 8,
所求解集为:
[
  [2,2,2,2],
  [2,3,3],
  [3,5]
]
"""
class Solution(object):
    def combinationSum(self, candidates, target):
        """
        :type candidates: List[int]
        :type target: int
        :rtype: List[List[int]]
        """

        res = set()
        self.walk(res, candidates, [], target)
        return list(map(list, list(res)))
        

    def walk(self, res, candidates, intnums, target):
      
      # 判断当前是否满足递归的终止条件
      # intnums = map(int, curnums.strip().split(","))

      if sum(intnums) > target:
        return

      if sum(intnums) == target:
        res.add(tuple(sorted(intnums)))
        return

      for c in candidates:
        self.walk(res, candidates, intnums+[c], target)
      return

sol = Solution()
print(sol.combinationSum([2,3,5], 8))