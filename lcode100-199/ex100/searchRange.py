#coding: utf-8
class Solution(object):
    def searchRange(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: List[int]
        """

        low, high = 0, len(nums)-1

        # 先查最左边        
        while low < high:
            mid = (low + high)//2
            if nums[mid] < target:
                low = mid +1
            else:
                high = mid
        
        if not nums or nums[low] != target:
            return [-1, -1]
        
        a, b = low, len(nums) -1
        while a < b:
            mid = (a + b + 1) // 2
            if nums[mid] > target:
                b = mid -1
            else:
                a = mid
        return [low, a]
    
sol = Solution()
print(sol.searchRange(nums = [8,8,8], target = 8))