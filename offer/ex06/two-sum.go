package main

import "fmt"

// 输入一个递增排序的数组和一个数字S，在数组中查找两个数，使得他们的和正好是S，
// 如果有多对数字的和等于S，输出两个数的乘积最小的。
//
//Python模板：
//# -*- coding:utf-8 -*-
//class Solution:
//    def FindNumbersWithSum(self, array, tsum):
//        # write code here

func FindNumberWithSum(array []int, tsum int) [2]int {
	plow, phigh := 0, len(array)-1
	res := [2]int{}
	// << 表示有意
	minNum := 2<<16 - 1
	for plow < phigh {
		csum := array[plow] + array[phigh]
		if csum == tsum {
			cmul := array[plow] * array[phigh]
			if cmul < minNum {
				minNum = cmul
				res = [2]int{array[plow], array[phigh]}
			}
			plow++
		} else if csum < tsum {
			plow++
		} else {
			phigh--
		}
	}

	return res
}

func main() {
	fmt.Println(FindNumberWithSum([]int{1, 2, 4, 7, 11, 15}, 15))
}
