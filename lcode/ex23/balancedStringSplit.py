class Solution(object):
    def balancedStringSplit(self, s):
        """
        :type s: str
        :rtype: int
        """
        hqueue = []
        res = 0
        if len(s) < 2:
            return 0
        i = 0

        while i<len(s):
            if len(hqueue) != 0 and s[i] != hqueue[-1] :
                hqueue.pop()
                if len(hqueue) == 0:
                    res+=1
            else:
                hqueue.append(s[i])
            i+=1
        if len(hqueue) != 0:
            return -1
        else:
            return res

sol = Solution()
print(sol.balancedStringSplit("RLRRLLRLRL"))