class Solution:
    def subsetsWithDup(self, nums):
        res = []
        def count(start, num):
            if sorted(num) not in res:
                res.append(sorted(num))
            for i in range(start, len(nums)):
                count(i+1, num+[nums[i]])
                                
            return
        count(0, [])
        return res

s = Solution()
print(s.subsetsWithDup([4,4,4,1,4]))