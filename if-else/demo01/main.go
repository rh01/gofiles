package main

import "fmt"

func main() {
	var array = [...]int{1, 2, 3, 4, 5, 6} // 注意这个是数组
	var maxLen int

	maxLen = len(array) - 1
	for i, e := range array {
		if i == maxLen {
			array[0] += e
		} else {
			array[i+1] += e
		}
	}

	fmt.Println(array)
}
