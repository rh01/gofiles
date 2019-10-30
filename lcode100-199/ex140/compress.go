package main

import (
	"fmt"
	"strconv"
)

func compress(chars []byte) int {
	left, right := 0, 0
	//size := len(chars)
	// res := 0
	for right < len(chars) {
		for right < len(chars) && chars[right] == chars[left] {
			right++
		}
		count := right - left
		if count == 1 {
			left++
		}

		if count > 1 {
			charsRight := chars[right:]
			chars = append(chars[:left+1], []byte(strconv.Itoa(count))...)
			fmt.Println(chars)
			left, right = len(chars), len(chars)
			chars = append(chars, charsRight...)
		}

		if left >= len(chars) {
			break
		}
	}

	return len(chars)
}

func main() {
	fmt.Println(compress([]byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}))
	fmt.Println(compress([]byte{'a'}))
	fmt.Println(compress([]byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'}))

}
