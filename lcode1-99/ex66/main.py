def find(nums, target):
    if target in nums:
        return [nums.index(target), nums[::-1].index(target)]
    return [-1, -1]