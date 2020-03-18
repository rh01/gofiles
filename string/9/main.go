package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {
	s := strconv.Itoa(x)
	l := len(s)
	res := 0
	for i := 0; i < l; i++ {
		a := extendFromCenter(s, i, i)
		b := extendFromCenter(s, i, i+1)
		res = max(res, max(a, b))
	}

	fmt.Println(res)

	if res == l {
		return true
	}

	return false
}

func extendFromCenter(s string, l, r int) int {
	size := len(s)
	// res := 0
	for l >= 0 && r < size && s[l] == s[r] {
		l--
		r++
	}

	return r - l - 1
}

func max(i0, i1 int) int {
	if i0 > i1 {
		return i0
	}
	return i1
}

func main() {
	s := 121
	fmt.Println(isPalindrome(s))
}
