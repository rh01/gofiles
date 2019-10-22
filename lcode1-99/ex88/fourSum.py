"""
// Copyright 2019 readailib.com. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ex88

// nums = [1, 0, -1, 0, -2, 2]，和 target = 0。
// [
//   [-1,  0, 0, 1],
//   [-2, -1, 1, 2],
//   [-2,  0, 0, 2]
// ]
func fourSum(nums []int, target int) [][]int {

}


/*
给定一个包含 n 个整数的数组 nums 和一个目标值 target，判断 nums 中是否存在四个元素 a，b，c 和 d ，使得 a + b + c + d 的值与 target 相等？找出所有满足条件且不重复的四元组。

注意：

答案中不可以包含重复的四元组。

示例：

给定数组 nums = [1, 0, -1, 0, -2, 2]，和 target = 0。

满足要求的四元组集合为：
[
  [-1,  0, 0, 1],
  [-2, -1, 1, 2],
  [-2,  0, 0, 2]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/4sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
"""

class Solution(object):
    def threeSum(self, nums):
        """
        :type nums: List[int]
        :rtype: List[List[int]]
        """
        res = []
        nums.sort()
        size = len(nums)
        if size < 3:
            return []
        for idx in range(size):
          for idx in range(size):
            if idx == 0 or nums[idx]>nums[idx-1]:
                l = idx+1
                r = len(nums)-1
                v = nums[idx]
           
                while l < r:
                    t = nums[l] + nums[r] + v
                    if t < 0:
                        l+=1
                    elif t > 0:	
                        r-=1
                    else:
                        elems = [v, nums[l], nums[r]] 
                        res.append([v, nums[l], nums[r]])
                        l+=1
                        r-=1
                        while l < r and nums[l] == nums[l-1]:
                            l += 1
                        while r > l and nums[r] == nums[r+1]:
                            r -= 1
        return res