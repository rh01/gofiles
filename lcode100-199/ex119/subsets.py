class Solution(object):
    def subsets(self, nums):
        """
        :type nums: List[int]
        :rtype: List[List[int]]
        """
        
        # bit computation
        res = []
        size = len(nums)
        for i in range(2**size):
            print(i)
            s = bin(i)[2:]
            s = s[::-1]
            strs = []
            for i in range(len(s)):
                if s[i] == '1':
                    strs.append(nums[i])
            # if strs not in res: 
            res.append(strs)
        return res
    
sol = Solution()
print(sol.subsets([0]))