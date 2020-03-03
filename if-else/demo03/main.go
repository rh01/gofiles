package main

import "fmt"

func main() {
	var array = [...]int8{0, 1, 2, 3, 4, 5, 6, 7}

	switch int8(1 + 3) { // 如果这里不用强制类型转换是不会通过编译的
	case array[0], array[1]:
		fmt.Println("array[0], array[1]")
	case array[2], array[3], array[4]:
		fmt.Println("array[2],array[3],array[4]")
	case array[5], array[6], array[7]:
		fmt.Println("array[5], array[6], array[7]")
	}
}
