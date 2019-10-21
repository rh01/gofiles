class Solution(object):
    def isValid(self, strs):
        """
        :type s: str
        :rtype: bool
        """
        length = len(strs)
        if strs == None or len(strs) == 0:
            return True
        hstack = []
        sArr = list(strs)
        hstack.append(sArr.pop(0))
        while  sArr:
            next_item = sArr.pop(0)
            if next_item in ["(", "[", "{"]:
                hstack.append(next_item)
            else:
                if hstack==None or len(hstack) ==0:
                    return False
                top = hstack.pop()
                if next_item == ")":
                    if top != "(":
                        return False
                elif next_item == "]":
                    if top != "[":
                        return False
                elif next_item == "}":
                    if top != "{":
                        return False

        if len(hstack) != 0:
            return False
        return True

sol = Solution()
print(sol.isValid("()]"))