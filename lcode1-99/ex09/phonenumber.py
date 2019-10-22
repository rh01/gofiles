# 输入："23"
# 输出：["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].


class Solution(object):
    def letterCombinations(self, digits):
        """
        :type digits: str
        :rtype: List[str]
        """
        # 利用队列做这个方法很精巧
        # 尽力还原该算法
        queue = []
        mapping = ["0", "1", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"]
        queue.append("")
        for index, i in enumerate(digits):
            x = int(i)
            #print(queue)
            while len(queue[0])==index:
                s = queue.pop(0)
                for t in mapping[x]:
                    queue.append(s+t)

        return queue

# print(findAllPhoneNumber("23"))
