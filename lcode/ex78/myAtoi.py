class Solution(object):
    def myAtoi(self, str):
        """
        :type str: str
        :rtype: int
        """
        INT_MIN = -2**31
        INT_MAX = 2**31-1
        s = str.strip()
        if s[0] not in "1234567890+-":
            return 0

        res = "0"
        neg =False
        i = 0

        if s[0] == "+":
            neg = False
            i = 1
        if s[0] == "-":
            neg = True
            i = 1
        
    
        while i < len(s) and s[i] in "1234567890":
            res+=s[i]
            i+=1

        if neg:
            return max(-1*int(res), INT_MIN)
        else:
            return min(int(res), INT_MAX)


sol = Solution()
print(sol.myAtoi("-91283472332"))
