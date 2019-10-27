"""
5239. 循环码排列 显示英文描述 
用户通过次数5
用户尝试次数6
通过次数5
提交次数6
题目难度Medium
给你两个整数 n 和 start。你的任务是返回任意 (0,1,2,,...,2^n-1) 的排列 p，并且满足：

p[0] = start
p[i] 和 p[i+1] 的二进制表示形式只有一位不同
p[0] 和 p[2^n -1] 的二进制表示形式也只有一位不同
 

示例 1：

输入：n = 2, start = 3
输出：[3,2,0,1]
解释：这个排列的二进制表示是 (11,10,00,01)
     所有的相邻元素都有一位是不同的，另一个有效的排列是 [3,1,0,2]
示例 2：

输出：n = 3, start = 2
输出：[2,6,7,5,4,0,1,3]
解释：这个排列的二进制表示是 (010,110,111,101,100,000,001,011)
 

提示：

1 <= n <= 16
0 <= start < 2^n
"""


class Solution(object):
    def circularPermutation(self, n, start):
        """
        :type n: int
        :type start: int
        :rtype: List[int]
        """
        if n <= 0:
            return []
        
        res = []
        startstr = (n - len(bin(start)[2:]))*'0'+bin(start)[2:]
        print(startstr)
        self.dfs(n, start, res, 0,  list(startstr))
        return res
    
    def dfs(self, n, start, res, curnum, startstr):
        hmap = {"0":"1", "1":"0"}
        if start in res:
            return
        res.append(start)
        if curnum == 2**n-1:
            return
        
        for i in range(n):
            startstr[i] = hmap[startstr[i]]
            s1 = int("".join(startstr), 2)
            if s1 in res: continue
            self.dfs(n, s1, res, curnum+1, startstr)
            startstr[i] = hmap[startstr[i]]
        return
    
sol = Solution()
print(sol.circularPermutation(10, 2))