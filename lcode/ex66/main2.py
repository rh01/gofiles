def find(nums, target):
    length = len(nums)
    low, high = 0, length-1
    ret1, ret2 = -1, -1
    while low < high:
        mid = (low + high) // 2
        if nums[mid] < target:
            low = mid+1
        else:
            high = mid
    if nums[low] != target:
        return [-1, -1]
    else:
        ret1 = low
    
    high = length-1
    while low<high:
        mid = (low + high) // 2
        if nums[mid] > target:
            high = mid-1
        else:
            low = mid
    ret2 = high
    return [ret1, ret2]