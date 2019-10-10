
def solve(people, limit):
    people.sort()
    low = 0
    res = 0
    high = len(people)-1
    while low < high:
        if people[high] == limit: 
            res+=1
            high-=1
        elif people[high] < limit:
            if people[low] + people[high] > limit:
                res+=1
                high-=1
            if people[low] + people[high] <= limit:
                res+=1
                high-=1
                low+=1
        else:
            return -1
    return res


print(solve([3, 5, 3, 5],5))