class Solution:
    def subsets(self, nums):
        res = []
        def count(start, num):
            res.append(num)
            for i in range(start, len(nums)):
                count(i+1, num+[nums[i]])
            return
        count(0, [])
        return res

s = Solution()
print(s.subsets([1,2,3]))