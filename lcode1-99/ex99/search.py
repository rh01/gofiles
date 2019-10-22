class Solution(object):
    def search(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: int
        """
        size = len(nums)
        low1, high2 = 0, size -1
        high1 = low1
        low2 = high2
        
        while nums[high1] < nums[high1+1]:
            high1+=1

        while  nums[low2] > nums[low2-1]:
            low2 -=1

        # case1 Not Exist
        if (target < nums[low1] and target < nums[low2]) or \
            (target > nums[high1] and target > nums[high2]) :
            return -1
        
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
                    return low1
        
        
        if target <= nums[high2]:
            while low2 <= high2:
                mid = (low2 + high2) // 2
                midVal = nums[mid]
                if midVal > target:
                    high2 = mid-1
                    
                elif midVal < target:
                    low2 = mid +1
                else:
                    return low2

        return -1

sol = Solution()
print(sol.search([4,5,6,7,0,1,2], 0))