# -*- coding:utf-8 -*-
class Solution:
    def Permutation(self, ss):
        
        if not ss or len(ss) == 0:
            return []
        
        res = []
        self.walk(ss, res, "")
        return res
    
    def walk(self, ss, res, p):
        # 剪枝
        if len(ss) == 0: 
            res.append(p)
            return 
                        
        for i, v in enumerate(ss):
            self.walk(ss[:i]+ss[i+1:], res, p+v)
            
if __name__ == "__main__":
    s = Solution()
    print s.Permutation("abc")