package main

/*

将一个字符串转换成一个整数，要求不能使用字符串转换整数的库函数。
数值为0或者字符串不是一个合法的数值则返回0

输入描述:
输入一个字符串,包括数字字母符号,可以为空

输出描述:
如果是合法的数值表达则返回该数字，否则返回0

示例1
输入
+2147483647
    1a33
输出
2147483647
    0
# -*- coding:utf-8 -*-
class Solution:
    def StrToInt(self, s):
        # write code here
*/

func StrToInt(s string) int {
	sArray := []byte(s)
	res := 0
	isPositive := false
	for _, value := range sArray {
		if value == '+'{
			isPositive = true
			continue
		}
		if value == '-'{
			isPositive = false
			continue
		}
		if value-'0' >= 0 && value-'0' <= 9 {
			cnum := value - '0'
			res = res*10 + int(cnum)
		}else{
			return 0
		}
	}

	if !isPositive {
		return -res
	}else {
		return res
	}
}

func main() {
	println(StrToInt("-2147483647"))
	println(StrToInt("1a33"))
}
