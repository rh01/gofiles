class Solution(object):
    def rob(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        # return max(sum(nums[1::2]), sum(nums[::2]))
        length = len(nums)

        dp = [0 for _ in range(length)]
        # dp[i] = max(dp[i-2]+nums[i], dp[i-1])
        dp[1] = nums[0]
        dp[2] = max(nums[0], nums[1])
        for i in range(2, length):
            dp[i] = max(dp[i-2]+nums[i], dp[i-1])
        return dp[length]