
class Solution(object):
    def containsDuplicate(self, nums):
        """
        :type nums: List[int]
        :rtype: bool
        """
        hmap = {}
        for i in nums:
            if i in hmap:
                hmap[i]+=1
                return True
            else:
                hmap[i] =1
        return False