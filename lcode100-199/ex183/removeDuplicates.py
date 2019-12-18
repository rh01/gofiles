
class Solution(object):
    def removeDuplicates(self, s, k):
        """
        :type s: str
        :type k: int
        :rtype: str
        """
        flag = False
        
        idx = 1
        stack = [s[0]]
        count = 1
        
        while idx < len(s):
            temp = k
            stack.append(s[idx])   
            if s[idx] == s[idx-1]:
                count+=1
                if count == temp:
                    flag = True
                    while temp!=0:
                        stack.pop()
                        temp-=1

            else:
                count = 1
            idx+=1
        if not flag:
            return s
        return self.removeDuplicates(''.join(stack), k)

                



sol = Solution()
print(sol.removeDuplicates("yfttttfbbbbnnnnffbgffffgbbbbgssssgthyyyy", 4))