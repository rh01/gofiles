class Solution(object):
    def firstMissingPositive(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        for idx, v in enumerate(nums):
            if v != idx and 0 < v <= len(nums):
                pass
                

# sol = Solution()
# print(sol.firstMissingPositive([3,4,-1,1]))
