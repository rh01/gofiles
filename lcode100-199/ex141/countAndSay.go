package main

import (
	"fmt"
	"strconv"
)

func countAndSay(n int) string {
	a1 := []byte{'1'}
	for i := 1; i < n; i++ {
		compress(&a1)
		//fmt.Println("=====================")
	}
	return string(a1)
}

func compress(chars2 *[]byte) {

	chars := *chars2
	left, right := 0, 0
	//size := len(chars)
	// res := 0
	for right < len(chars) {
		for right < len(chars) && chars[right] == chars[left] {
			right++
		}
		count := right - left

		//fmt.Println(count)
		// if count == 1 {
		// 	left++
		// }

		if count >= 0 {
			
			charsRight:= string(chars[right:]) // 变成静态
			chaLeft := chars[left]

			chars = append(chars[:left], []byte(strconv.Itoa(count))...)
			chars = append(chars, chaLeft)
			temp := len(chars)
			chars = append(chars, []byte(charsRight)...)
			left, right = temp, temp
		}
		//fmt.Println(string(chars))

	}

	*chars2 = chars

	return
}

// func compress(chars *[]byte) {
// 	fmt.Println(string(*chars))
// 	left, right := 0, 0
// 	//size := len(chars)
// 	// res := 0
// 	for right < len(*chars) {
// 		for right < len(*chars) && (*chars)[right] == (*chars)[left] {
// 			right++
// 		}
// 		fmt.Println("right", right)
// 		count := right - left
// 		// if count == 1 {
// 		// 	left++
// 		// }
// 		fmt.Println("chars___", string(*chars))

// 		if count >= 1 {
// 			charsRight := (*chars)[right:]
// 			fmt.Println("charsRight", string(charsRight))
// 			chaLeft := (*chars)[left]
// 			*chars = append((*chars)[:left], []byte(strconv.Itoa(count))...)
// 			fmt.Println("append before left, ", string(*chars))
// 			//fmt.Println(chars)
// 			*chars = append(*chars, chaLeft)
// 			fmt.Println("append left, ", string(*chars))

// 			temp := len((*chars))
// 			*chars = append(*chars, charsRight...)
// 			fmt.Println("charsRight", string(charsRight))
// 			fmt.Println("apppend right", string(*chars))
// 			left, right = temp, temp
// 		}

// 		if left >= len(*chars) {
// 			break
// 		}
// 	}
// 	return
// 	//return len(chars)
// }

func main() {
	fmt.Println(countAndSay(6))
}

/*
报数序列是一个整数序列，按照其中的整数的顺序进行报数，得到下一个数。其前五项如下：

1.     1
2.     11
3.     21
4.     1211
5.     111221
6.     312211
1 被读作  "one 1"  ("一个一") , 即 11。
11 被读作 "two 1s" ("两个一"）, 即 21。
21 被读作 "one 2",  "one 1" （"一个二" ,  "一个一") , 即 1211。

给定一个正整数 n（1 ≤ n ≤ 30），输出报数序列的第 n 项。

注意：整数顺序将表示为一个字符串。



示例 1:

输入: 1
输出: "1"
示例 2:

输入: 4
输出: "1211"

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/count-and-say
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
