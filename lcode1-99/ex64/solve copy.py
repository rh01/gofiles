
def solve(length, nums, sumnum):
    # method1
    hmap = {}
    for i in range(length):
        cnum = nums[i]
        if cnum not in hmap.keys():
            hmap[cnum] = i
        target = sumnum - cnum
        if target in hmap.keys():
            return [i, hmap[target]]
    return [-1, -1]
    

length = int(input())
nums = list(map(int, input().split()))
target = int(input().strip())

res = solve(length, nums, target)
print(" ".join(list(map(str, res))))
