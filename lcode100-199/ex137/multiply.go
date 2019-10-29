package main

import "fmt"

func multiply(num1 string, num2 string) string {
	// 思路：将每个元素减去'0'得到差值另个也需要得到对应的差值
	var res1 int64
	var res2 int64
	// _ = res2
	// _ = res1
	for _, str1 := range num1 {
		res1 = res1*10 + (int64(str1) - int64('0'))
	}

	for _, str1 := range num2 {
		res2 = res2*10 + (int64(str1) - int64('0'))
	}

	mul := res1 * res2
	fmt.Println(res1)

	result := ""
	for mul != 0 {
		mod := mul % 10
		result = string(byte(mod+int64('0'))) + result
		mul = (mul - mod) / 10
	}

	return result
}

func main() {
	fmt.Println(multiply("498828660196",
		"840477629533"))
}
