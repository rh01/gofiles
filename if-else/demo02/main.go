package main

import "fmt"

func main() {
	var slice = []int{1, 2, 3, 4, 5, 6} // 注意这个是
	var maxLen int

	maxLen = len(slice) - 1
	for i, e := range slice {
		if i == maxLen {
			slice[0] += e
		} else {
			slice[i+1] += e
		}
	}

	fmt.Println(slice)
}
