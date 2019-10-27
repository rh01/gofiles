"""给定两个整数 n 和 k，返回 1 ... n 中所有可能的 k 个数的组合。

示例:

输入: n = 4, k = 2
输出:
[
  [2,4],
  [3,4],
  [2,3],
  [1,2],
  [1,3],
  [1,4],
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/combinations
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。"""

class Solution(object):
    def combine(self, n, k):
        """
        :type n: int
        :type k: int
        :rtype: List[List[int]]
        """
        if n <= 1 or n <1:
            return []
        
        if n == 1:
            return list(range(1, n+1))
        
        narr = list(range(1, n+1))
        res = []
        for i in range(1, n+1):
            res.append([i])
        self.dfs(narr, res, n, 1, 0)
        return res
    
    def dfs(self, narr, res, n, cur, k):
        
        if n == k:
            return
        
        
        
            
        
        for i in range(cur, n):
           
           
           
           
sol = Solution()
print(sol.combine(4, 2))
        