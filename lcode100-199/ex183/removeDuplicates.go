package main

import (
	"fmt"
	"strings"
)

func removeDuplicates(s string, k int) string {

	stack := make([][2]uint, 0)
	for i := 0; i < len(s); i++ {
		if len(stack) == 0 || stack[len(stack)-1][0] != uint(s[i]) {
			stack = append(stack, [2]uint{uint(s[i]), 1})
		} else if stack[len(stack)-1][1]+1 < uint(k) {
			stack[len(stack)-1][1]++
		} else {
			stack = stack[:len(stack)-1]
		}
	}

	
	ret := ""
	for _, val := range stack {
		ret += strings.Repeat(string(byte(val[0])), int(val[1]))
	}
	return ret
}

func main() {
	fmt.Printf("result: %v", removeDuplicates("pbbcggttciiippooaais", 2))
}
