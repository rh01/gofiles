class Solution(object):
    def longestCommonPrefix(self, strs):
        """
        :type strs: List[str]
        :rtype: str
        """

        if strs == None or len(strs) == 0: return ""
        if len(strs) < 2: return strs[0]

        targ = min(strs, key=len)
        i = 0
        res = ""
        while i<len(targ):
            for j in strs:
                if j[i] != targ[i]:
                    return res
            res+=targ[i]
            i+=1
        return res

sol = Solution()
print(sol.longestCommonPrefix(["floweraaaaa","flower","flight"]))
