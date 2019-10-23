from collections import Counter
class Solution(object):
    def isAnagram(self, s, t):
        """
        :type s: str
        :type t: str
        :rtype: bool
        """
        return Counter(t) == Counter(s)
    
sol = Solution()
print(sol.isAnagram(s = "anagram", t = "nagaram"))
print(sol.isAnagram(s = "rat", t = "car"))