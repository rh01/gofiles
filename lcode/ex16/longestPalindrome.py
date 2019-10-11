class Solution(object):
    def longestPalindrome(self, s):
        """
        :type s: str
        :rtype: str
        """
        if len(s) <= 1:
            return s
        length = len(s)
        dp = [0 for _ in range(length)]
        dp[0] = 1
        end = 1
        for i in range(1, length):
            
            if i - dp[i-1] +2 > 2 and s[i-dp[i-1]-1:i+1]==s[i-dp[i-1]-1:i+1][::-1]:
                dp[i] = dp[i-1]+2
                end = i+1
            elif s[i-dp[i-1]:i+1]==s[i-dp[i-1]:i+1][::-1]:
                dp[i] = dp[i-1]+1
                end=i+1
            else:
                dp[i] = dp[i-1]
        return s[end-dp[-1]:end]
            


sol = Solution()
print(sol.longestPalindrome("aba"))