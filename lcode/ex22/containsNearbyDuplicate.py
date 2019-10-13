# 给定一个整数数组和一个整数 k，判断数组中是否存在两个不同的索引 i 和 j，
# 使得 nums [i] = nums [j]，并且 i 和 j 的差的绝对值最大为 k。

class Solution(object):
    def containsNearbyDuplicate(self, nums, k):
        """
        :type nums: List[int]
        :type k: int
        :rtype: bool
        """
        hmap = {}
        for i, j in enumerate(nums):
            if j not in hmap:
                hmap[j] = i
            else:
                if abs(i - hmap[j]) <= k:
                    return True
                else:
                    hmap[j]=i
               
        return False

sol = Solution()
print(sol.containsNearbyDuplicate( nums = [1,2,3,1], k = 3))
print(sol.containsNearbyDuplicate(nums = [1,0,1,1], k = 1))
print(sol.containsNearbyDuplicate(nums = [1,2,3,1,2,3], k = 2))
                    

          