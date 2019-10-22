


def isGlodonNumber(num):
    hset = set()
    t = 0
    while t == len(hset):
        num = sum(map(lambda x: int(x)**2, list(str(num))))
        if num == 1:
            return True
        t+=1
        hset.add(num)
    return False


        # if 1 < num < 4:
            # return False
    # return True
    # map(lambda x: x^2, intlist)
print(isGlodonNumber(32))