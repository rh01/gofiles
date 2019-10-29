package main

import (
	"fmt"
	"strconv"
)

func multiply(num1 string, num2 string) string {
	// 思路：将每个元素减去'0'得到差值另个也需要得到对应的差值
	var res, size1, size2 uint32
	size1 = uint32(len(num1))
	size2 = uint32(len(num2))


	for i := uint(1); i < size1+1; i++ {
		for j := uint(1); j < uint(len(num2)+1); j++ {
			res = res + int(num1[len(num1)-i]) * uint(num2[len(num2)-j]) * int(10<<uint32(i+j-2))
		}
	}
	return strconv.Itoa(int(res))
}

func main() {
	fmt.Println(multiply("123",
		"456"))
}
