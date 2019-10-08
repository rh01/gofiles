package main

import "strings"

//汇编语言中有一种移位指令叫做循环左移（ROL），现在有个简单的任务，
//就是用字符串模拟这个指令的运算结果。
//对于一个给定的字符序列S，请你把其循环左移K位后的序列输出。
//例如，字符序列S=”abcXYZdef”,要求输出循环左移3位后的结果，
//即“XYZdefabc”。是不是很简单？OK，搞定它！
//
//python实例代码
//
//# -*- coding:utf-8 -*-
//class Solution:
//	def LeftRotateString(self, s, n):
//		# write code here

func main() {
	println(LeftRotateString("aasssss", 2))
	//fmt.Println("ssssss")
}

func LeftRotateString(s string, n int) string {
	//bytes := []byte(s)
	//u := string(s[1])
	// split := strings.Split(s, "")
	//fmt.Println(split)

	// (xTyT)T=yx
	// strings.Builder{}
	sArray := strings.Split(s, "")
	length := len(sArray)
	// X - > X->T
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		swap(sArray, i, j)
	}
	//
	for i, j := n, length-1; i <= j; i, j = i+1, j-1 {
		swap(sArray, i, j)
	}

	for i, j := 0, length-1; i <= j; i, j = i+1, j-1 {
		swap(sArray, i, j)
	}

	return strings.Join(sArray, "")

}

func swap(array []string,i, j int) {
	var s3 string
	s3 = array[i]
	array[i] = array[j]
	array[j] = s3
}

//		# write code here
