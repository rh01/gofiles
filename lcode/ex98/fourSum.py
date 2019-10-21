class Solution(object):
    def fourSum(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: List[List[int]]
        """
        if len(nums) < 4 or nums is None: return []

        nums.sort()
        res = []
        size = len(nums)
        for i in range(size-2):
            if i>0 and nums[i]==nums[i-1]:
                continue
            for j in range(i+1, size-1):
                if j>i+1 and nums[j]==nums[j-1]:
                    continue
                num1 = nums[i]
                num2 = nums[j]

                low = j + 1
                high = size -1
                while low < high:
                    num3 = nums[low]
                    num4 = nums[high]

                    csum = num1 + num2 + num3 + num4
                    if csum <target:
                        low+=1
                    elif csum > target:
                        high-=1
                    else:
                        res.append([num1, num2, num3, num4])
                        while low+1 < high and nums[low+1] == nums[low]:
                            low+=1
                        while high-1 > low and nums[high-1] == nums[high]:
                            high-=1
                        low+=1
                        high-=1
        return res 

sol = Solution()
print(sol.fourSum(nums = [1, 0, -1, 0, -2, 2],target = 0))