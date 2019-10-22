"""
给定一个可包含重复数字的序列，返回所有不重复的全排列。

示例:

输入: [1,1,2]
输出:
[
  [1,1,2],
  [1,2,1],
  [2,1,1]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/permutations-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
"""
class Solution(object):
    def permuteUnique(self, nums):
        """
        :type nums: List[int]
        :rtype: List[List[int]]
        """
        res = []
        self.walk(res, nums, [])
        return res
    
    def walk(self, res, nums, curnums):
        if len(nums) == 0:
            if curnums not in res:
                res.append(curnums)
            return
        for idx, val in enumerate(nums):
            self.walk(res, nums[:idx]+nums[idx+1:], curnums+[val])
        return

sol = Solution()
print(sol.permuteUnique( [1,1,2]))