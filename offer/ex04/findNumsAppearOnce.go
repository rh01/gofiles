package main

import "fmt"

/*
// 一个整型数组里除了两个数字之外，其他的数字都出现了两次。请写程序找出这两个只出现一次的数字。
# -*- coding:utf-8 -*-
class Solution:
    # 返回[a,b] 其中ab是出现一次的两个数字
    def FindNumsAppearOnce(self, array):
        # write code here
*/

func main() {
	fmt.Println(FindNumsAppearOnce([]int{1,2,3,4,2,3}))
}

// 返回出现一次得两个数字
func FindNumsAppearOnce(array []int) [2]int {
	//res是两个只出现一次得两个数字得抑或结果
	var res int = 0
	for _, value := range array {
		res = res ^ value
	}

	if res == 0 {
		return [2]int{}
	}

	//这里应该加一判断，因为对于有符号整数来说~xxx表示补码得反码
	// 参考：https://colobu.com/2016/06/20/dive-into-go-5/#%E9%80%BB%E8%BE%91%E8%BF%90%E7%AE%97%E7%AC%A6
	res = (^(^-res + 1) + 1) ^ res
	num1, num2 := 0, 0
	for _, value := range array {
		if (value & res) != 0 {
			num1 = num1 ^ value
		} else {
			num2 = num2 ^ value
		}
	}

	return [2]int{num2, num1}
}
