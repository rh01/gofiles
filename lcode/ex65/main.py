def vaild(strs):
    length = len(strs)
    if strs == None or len(strs) == 0:
        return 1
    hstack = []
    sArr = list(strs)
    hstack.append(sArr.pop(0))
    while hstack and sArr:
        next_item = sArr.pop(0)
        if next_item in ["(", "[", "{"]:
            hstack.append(next_item)
        elif next_item == ")":
            top = hstack.pop()
            if top != "(":
                return 0
        elif next_item == "]":
            top = hstack.pop()
            if top != "[":
                return 0
        elif next_item == "}":
            top = hstack.pop()
            if top != "{":
                return 0
    if len(hstack) != 0:
        return 0
    return 1

s = input()
print(vaild(s))