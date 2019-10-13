"""
给定一个整数数组，判断数组中是否有两个不同的索引 i 和 j，
使得 nums [i] 和 nums [j] 的差的绝对值最大为 t，
并且 i 和 j 之间的差的绝对值最大为 ķ。

示例 1:

输入: nums = [1,2,3,1], k = 3, t = 0
输出: true
示例 2:

输入: nums = [1,0,1,1], k = 1, t = 2
输出: true
示例 3:

输入: nums = [1,5,9,1,5,9], k = 2, t = 3
输出: false

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/contains-duplicate-iii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
"""
class Solution:
    def containsNearbyAlmostDuplicate(self, nums, k, t) -> bool:
        if t < 0 or k < 0:
            return False
        all_buckets = {}
        bucket_size = t + 1                     # 桶的大小设成t+1更加方便
        for i in range(len(nums)):
            bucket_num = nums[i] // bucket_size # 放入哪个桶
            
            if bucket_num in all_buckets:       # 桶中已经有元素了
                return True
            
            all_buckets[bucket_num] = nums[i]   # 把nums[i]放入桶中
            
            if (bucket_num - 1) in all_buckets and abs(all_buckets[bucket_num - 1] - nums[i]) <= t: # 检查前一个桶
                return True
            
            if (bucket_num + 1) in all_buckets and abs(all_buckets[bucket_num + 1] - nums[i]) <= t: # 检查后一个桶
                return True
            
            # 如果不构成返回条件，那么当i >= k 的时候就要删除旧桶了，以维持桶中的元素索引跟下一个i+1索引只差不超过k
            if i >= k:
                all_buckets.pop(nums[i-k]//bucket_size)
                
        return False

            

sol = Solution()
# print(sol.containsNearbyAlmostDuplicate(nums = [1,2,3,1], k = 3, t = 0))
# print(sol.containsNearbyAlmostDuplicate( nums = [1,0,1,1], k = 1, t = 2))
# print(sol.containsNearbyAlmostDuplicate(nums = [1,5,9,1,5,9], k = 2, t = 3))
print(sol.containsNearbyAlmostDuplicate(nums = [-3,3], k = 2, t = 4))
print(sol.containsNearbyAlmostDuplicate(nums=[0,2147483647],k=1,t=2147483647))