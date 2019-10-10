"""
给出集合 [1,2,3,…,n]，其所有元素共有 n! 种排列。

按大小顺序列出所有排列情况，并一一标记，当 n = 3 时, 所有排列如下：

"123"
"132"
"213"
"231"
"312"
"321"
给定 n 和 k，返回第 k 个排列。

说明：

给定 n 的范围是 [1, 9]。
给定 k 的范围是[1,  n!]。
示例 1:

输入: n = 3, k = 3
输出: "213"
示例 2:

输入: n = 4, k = 9
输出: "2314"

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/permutation-sequence
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
"""

import math
class Solution(object):
    def getPermutationv1(self, n, k):
        """
        :type n: int
        :type k: int
        :rtype: str
        """
        nums = list(range(1, n+1))
        res = []
        self.walk(res, nums, [],k)
        return res
    
    def walk(self, res, nums, curnums, k):
        
        if len(nums) == 0:
            k-=1
            if k == 0:
                return curnums
        # if len(nums) == 0:
        #     #curtime+=1
        #     if k == 0:
        #         res.append(curnums)
               # return
        for idx, val in enumerate(nums):
            self.walk(res, nums[:idx]+nums[idx+1:], curnums+[val], k)

    
    def getPermutation(self, n: int, k: int) -> str:
        result = ''
        mod = k-1
        if n==1:return '1'
        n_set = [str(i) for i in range(1,n+1)]
        fac_n = math.factorial(n-1)
        for i in range(n):
            val,mod = divmod(mod,fac_n)
            #print(i,fac_n,val,mod,result,n_set)
            fac_n = fac_n//(n-i-1)
            result+=n_set[val]
            del n_set[val]
            if mod==0:
                for i in n_set:result+=i
                return result
            
        return result
        

sol = Solution()
print(sol.getPermutation(3, 3))        