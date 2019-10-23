"""
还记得童话《卖火柴的小女孩》吗？现在，你知道小女孩有多少根火柴，
请找出一种能使用所有火柴拼成一个正方形的方法。不能折断火柴，
可以把火柴连接起来，并且每根火柴都要用到。

输入为小女孩拥有火柴的数目，每根火柴用其长度表示。
输出即为是否能用所有的火柴拼成正方形。

示例 1:

输入: [1,1,2,2,2]
输出: true

解释: 能拼成一个边长为2的正方形，每边两根火柴。
示例 2:

输入: [3,3,3,3,4]
输出: false

解释: 不能用所有火柴拼成一个正方形。
注意:

给定的火柴长度和在 0 到 10^9之间。
火柴数组的长度不超过15。

"""

class Solution(object):
    def makesquare(self, nums):
        """
        :type nums: List[int]
        :rtype: bool
        """
        # 检查
        if nums is None or len(nums) < 4:
            return False
        nums.sort(reverse=True)
        # 判断是否火柴的长度之和是否为4的倍数
        if sum(nums) // 4 * 4 != sum(nums):
            return False
        
        return self.dfs(nums, 0, len(nums), 0,  0, 0, 0, sum(nums)//4)

    def dfs(self, nums, cur, length, i1, i2, i3, i4, i5):
        # 判断是否遍历完且每个边的长度相等
        if cur == length:
            if i5 == i1 and i5 == i2 and i5 == i3 and i5 == i4:
                return True
            else:
                return False
        
        # 剪枝
        if i1 > i5 or  i2 > i5  or i3 > i5 or i4 > i5:
            return False

        return self.dfs(nums, cur+1, length, i1+nums[cur], i2, i3, i4, i5) or  \
                self.dfs(nums, cur+1, length, i1, i2+nums[cur], i3, i4, i5) or  \
                self.dfs(nums, cur+1, length, i1, i2, i3+nums[cur], i4, i5) or  \
                self.dfs(nums, cur+1, length, i1, i2, i3, i4+nums[cur], i5)
        
sol = Solution()
print(sol.makesquare( [1,1,2,2,2]))
print(sol.makesquare([3,3,3,3,4]))
print(sol.makesquare([5969561,8742425,2513572,3352059,9084275,2194427,1017540,2324577,6810719,8936380,7868365,2755770,9954463,9912280,4713511]))