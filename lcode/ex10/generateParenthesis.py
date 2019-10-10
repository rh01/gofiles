class Solution(object):
    def generateParenthesis(self, n):
        """
        :type n: int
        :rtype: List[str]
        """
        res = []
        # bfs加剪枝
        self.fn(res, "", 0, 0, n)
        return res

    def fn(self, res, string, left, right, n):
        if left > n or right > n or right > left:
            return
        # 递归的终止条件
        if left == n and right == n:
            res.append(string)
            return
        
        self.fn(res, string + "(", left+1, right, n)
        self.fn(res, string+")", left, right+1, n)
        return


sol = Solution()
print(sol.generateParenthesis(3))