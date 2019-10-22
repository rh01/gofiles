# 给定一个整数数组和一个整数 k，判断数组中是否存在两个不同的索引 i 和 j，
# 使得 nums [i] = nums [j]，并且 i 和 j 的差的绝对值最大为 k。

class Solution(object):
    def containsNearbyDuplicate(self, nums, k):
        """
        :type nums: List[int]
        :type k: int
        :rtype: bool
        """
        hmap = {}
        for i, j in enumerate(nums):
            if j not in hmap:
                hmap[j] = [i]
            else:
                # hmap[j].append(i)
                tlist = hmap[j]
                length = len(tlist)

                #print(hmap)
                i2 = 0
                while i2 < length:
                    if i < tlist[i2]:
                        if (abs(i - tlist[i2]) <= k or abs(i-tlist[i2-1]) <= k):
                            return True
                        else:
                            hmap[j] = hmap[j][:i2] +[i] + hmap[j][i2:]
                            break
                    i2+=1
                
                if  i2 == length:  
                    if abs(i - tlist[-1]) <= k:
                        return True
                    else:
                        hmap[j].append(i)
               
        return False

sol = Solution()
print(sol.containsNearbyDuplicate( nums = [1,2,3,1], k = 3))
print(sol.containsNearbyDuplicate(nums = [1,0,1,1], k = 1))
print(sol.containsNearbyDuplicate(nums = [1,2,3,1,2,3], k = 2))
                    

          