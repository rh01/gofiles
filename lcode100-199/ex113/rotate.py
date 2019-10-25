class Solution(object):
    def rotatev1(self, nums, k):
        """
        :type nums: List[int]
        :type k: int
        :rtype: None Do not return anything, modify nums in-place instead.
        """
        if nums is None or len(nums) ==0:
            return 
        k %= n
        for i in range(k):
            nums.insert(0, nums.pop())
        
    
    def rotate(self, nums, k):
        n = len(nums)
        k %= n
        nums.reverse()
        def reverse(nums, start, end):
            while start < end:
                nums[start], nums[end] = nums[end], nums[start]
                start+=1
                end -=1
        reverse(nums, 0, k-1)
        reverse(nums, k, n-1)

        
        #nums =  nums[-k:] + nums[:-k]

sol = Solution()
arr =  [1,2,3,4,5,6,7]
sol.rotate(arr, k = 3)
print(arr)