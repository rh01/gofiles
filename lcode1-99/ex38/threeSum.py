class Solution(object):
    def threeSum(self, nums):
        """
        :type nums: List[int]
        :rtype: List[List[int]]
        """
        res = []
        nums.sort()
        size = len(nums)
        if size < 3:
            return []
        for idx in range(size):
            if idx == 0 or nums[idx]>nums[idx-1]:
                l = idx+1
                r = len(nums)-1
                v = nums[idx]
           
                while l < r:
                    t = nums[l] + nums[r] + v
                    if t < 0:
                        l+=1
                    elif t > 0:	
                        r-=1
                    else:
                        elems = [v, nums[l], nums[r]] 
                        res.append([v, nums[l], nums[r]])
                        l+=1
                        r-=1
                        while l < r and nums[l] == nums[l-1]:
                            l += 1
                        while r > l and nums[r] == nums[r+1]:
                            r -= 1
        return res

sol = Solution()
print(sol.threeSum( nums = [-1, 0, 1, 2, -1, -4]))



