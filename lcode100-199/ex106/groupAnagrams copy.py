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
        res = []
        
        for s in strs:
            ss ="".join(sorted(s))
            if  ss not in dictA.keys():
                dictA[ss] = [s]
            else:
                dictA[ss].append(s)
        return list(dictA.values())

sol = Solution()
print(sol.groupAnagrams(["eat", "tea", "tan", "ate", "nat", "bat"]))