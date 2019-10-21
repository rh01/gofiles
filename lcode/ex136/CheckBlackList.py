#!/bin/python
# -*- coding: utf8 -*-
import sys
import os
import re

#请完成下面这个函数，实现题目要求的功能
#当然，你也可以不按照下面这个模板来作答，完全按照自己的想法来 ^-^ 
#******************************开始写代码******************************

def  CheckBlackList(userIP, blackIP):
    subs = blackIP.split("/")
    if len(subs) == 1:
        return userIP == blackIP
    else:
        ip1s = subs[0].split(".")
        ip2s = userIP.split(".")
        ip1size = len(ip1s)
        ip2size = len(ip2s)

        if ip1size != ip2size:
            return False
        res = []
        from operator import xor
        for i in range(ip1size):
            res.append(xor(int(ip1s[i]), int(ip2s[i])))
        subnet = subs[1]
        comp = 0
        for i in res:
            comp = comp*(2**8) + i
        if comp < 2**(32-int(subnet)):
            return True



#******************************结束写代码******************************


if __name__ == "__main__":
    try:
        _userIP = input()
    except:
        _userIP = None

    try:
        _blackIP = input()
    except:
        _blackIP = None


    res=CheckBlackList(_userIP, _blackIP)

    print(str(int(res)) + "\n")
    #print(res)