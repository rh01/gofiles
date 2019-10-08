package main

import (
	"container/list"
	//"log"
)

/*
# -*- coding:utf-8 -*-
class Solution:
    def LastRemaining_Solution(self, n, m):
        # write code here


每年六一儿童节,牛客都会准备一些小礼物去看望孤儿院的小朋友,今年亦是如此。
HF作为牛客的资深元老,自然也准备了一些小游戏。其中,有个游戏是这样的:

首先,让小朋友们围成一个大圈。
然后,他随机指定一个数m,让编号为0的小朋友开始报数。
每次喊到m-1的那个小朋友要出列唱首歌,
然后可以在礼品箱中任意的挑选礼物,并且不再回到圈中,从他的下一个小朋友开始,继续0...m-1报数....
这样下去....直到剩下最后一个小朋友,
可以不用表演,并且拿到牛客名贵的“名侦探柯南”典藏版(名额有限哦!!^_^)。

请你试着想下,哪个小朋友会得到这份礼品呢？(注：小朋友的编号是从0到n-1)
*/

func LastRemaining_Solutionv1(n, m int) int {
	//TODOTODOTODO
	if n == 0 {
		return -1
	}

	if n == 1 {
		return 0
	} else {
		return (LastRemaining_Solutionv1(n-1, m) + m ) % n
	}
}

func LastRemaining_Solution(n, m int) int {
	//TODOTODOTODO
	li := list.New()
	for i := 0; i < n; i++ {
		li.PushBack(i)
	}
	count := 0
	res := 0
	for li.Len() > 0 {
		if count == m-1 {
			res = li.Remove(li.Front()).(int)
			//log.Print(res)
			count = 0
		}else{
			li.PushBack(li.Remove(li.Front()))
			count++
		}
	}
	return res
}

func main() {
	println(LastRemaining_Solution(10, 2))
}
