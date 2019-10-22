class Solution(object):
    def convert(self, s, numRows):
        """
        :type s: str
        :type numRows: int
        :rtype: str
        """
        if s is None or len(s) == 0: return ""
        if numRows < 2: return s
        length = len(s) # 计算字符串s的长度，结合numRows计算出numCols
        # 上一步我放弃了
        res = ""
        i = 0
        while i<numRows:
            count = 0
            shift = list(range(2*(numRows-1), 0, -2)) + [2*(numRows-1)]
            cur = i
            res+=s[cur]
            while shift[i] + cur < length:
                j = shift[i] + cur
                res+=s[j]
                count+=1
                cur = j
                shift.reverse()
               
            i+=1
        return res

sol = Solution()
print(sol.convert("LEETCODEISHIRING", 3))
print(sol.convert(s = "LEETCODEISHIRING", numRows = 4))