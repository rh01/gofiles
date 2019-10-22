"""
给定一个没有重复数字的序列，返回其所有可能的全排列。

示例:

输入: [1,2,3]
输出:
[
  [1,2,3],
  [1,3,2],
  [2,1,3],
  [2,3,1],
  [3,1,2],
  [3,2,1]
]
"""
class Solution(object):
    def permute(self, nums):
        """
        :type nums: List[int]
        :rtype: List[List[int]]
        """
        res = []
        self.walk(res, nums, [], 0, len(nums))
        return res
    
    def walk(self, res, nums, curnums, level, n):
        if level == n:
            res.append(curnums)
            return
        
        for idx, val in enumerate(nums):
            self.walk(res, nums[:idx]+nums[idx+1:], curnums+[val], level+1, n)
        return

sol = Solution()
print(sol.permute([1,2,3]))        