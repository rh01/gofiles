
# 6
# 1 2 3 4 5 6
# 3

def solve(nums, target):
    # method1
    if len(nums) < 2:
        return [-1, -1]
    nums.sort()
    low, high = 0, len(nums)-1
    while low < high:
        csum = nums[low] + nums[high]
        if csum < target:
            low+=1
        elif csum > target:
            high-=1
        else:
            return [low, high]
    return [-1, -1]