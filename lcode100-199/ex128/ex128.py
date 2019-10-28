class Solution(object):
    def search(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: int
        """
        if nums  is None or len(nums ) == 0:
            return False
        size = len(nums)
        low1, high2 = 0, size -1
        high1 = low1
        low2 = high2
        
        while high1+1 < size and nums[high1] <= nums[high1+1]:
            high1+=1

        while low2-1 > 0 and nums[low2] >= nums[low2-1]:
            low2 -=1

        # case1 Not Exist
        if (target < nums[low1] and target < nums[low2]) or \
            (target > nums[high1] and target > nums[high2]) :
            return False
        
        # case2
        # You may assume no duplicate exists in the array.
        if target >= nums[low1]:
            while low1 <= high1:
                mid = (low1 + high1) // 2
                midVal = nums[mid]
                if midVal > target:
                    high1 = mid-1
                elif midVal < target:
                    low1 = mid +1
                else:
                    return True
        
        
        if target <= nums[high2]:
            while low2 <= high2:
                mid = (low2 + high2) // 2
                midVal = nums[mid]
                if midVal > target:
                    high2 = mid-1
                    
                elif midVal < target:
                    low2 = mid +1
                else:
                    return True

        return False