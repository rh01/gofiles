package main

import "fmt"

func main() {
	var array = [...]int8{0, 1, 2, 3, 4, 5, 6, 7}

	switch array[4] { // 这里不同的是，如果case中无符号的常量会自动转换为switch中表达式的结果类型
	case array[0],array[1],array[2]:  
		fmt.Println("array[0], array[1]")
	case array[0],array[1],array[2]:
		fmt.Println("array[2],array[3],array[4]")
	case array[0],array[1],array[2]:
		fmt.Println("array[5], array[6], array[7]")
	}
}
