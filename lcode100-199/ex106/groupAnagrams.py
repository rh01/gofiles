"""
给定一个字符串数组，将字母异位词组合在一起。
字母异位词指字母相同，但排列不同的字符串。

示例:

输入: ["eat", "tea", "tan", "ate", "nat", "bat"],
输出:
[
  ["ate","eat","tea"],
  ["nat","tan"],
  ["bat"]
]
说明：

所有输入均为小写字母。
不考虑答案输出的顺序
"""

from collections import Counter
# from functools import 
class Solution(object):
    def groupAnagrams(self, strs):
        """
        :type strs: List[str]
        :rtype: List[List[str]]
        """
        # 两个字典
        dictA = {}
        dictB = {}
        
        for s in strs:
            val = Counter(s)
            if  val not in dictA.values():
                dictA[s] = Counter(s)
                dictB[s] = [s]
            else:
                for i in dictA.keys():
                    if dictA[i] == val:
                        dictB[i].append(s)
        res = []
        for i in dictB.values():
            res.append(i)
        return res

sol = Solution()
print(sol.groupAnagrams(["eat", "tea", "tan", "ate", "nat", "bat"]))
