class Solution(object):
    def strStr(self, haystack, needle):
        """
        :type haystack: str
        :type needle: str
        :rtype: int
        """
        length = len(haystack)
        k = 0
        while k < length:
            j = 0
            i = k
            while i<length and j < len(needle) and haystack[i] == needle[j]:
                j+=1
                i+=1
            if j == len(needle):
                return i - len(needle)
            
            k+=1
        return -1


            



    def strStrv1(self, haystack, needle):
        """
        :type haystack: str
        :type needle: str
        :rtype: int
        """
        try:
            return haystack.index(needle)
        except ValueError:
            return -1

sol = Solution()
print(sol.strStr(haystack = "hello", needle = "el"))
print(sol.strStr(haystack = "aaaaa", needle = "bba"))
print(sol.strStr("mississippi","issip"))