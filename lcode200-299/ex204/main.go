package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	strArr1 := strings.Split(strings.TrimSpace(s), " ")
	strArr2 := make([]string, 0)
	for i := len(strArr1) - 1; i >= 0; i-- {
		if strArr1[i] != " " {
			strArr2 = append(strArr2, strArr1[i])
		}
	}
	return strings.Join(strArr2, " ")
}

func main() {
	s := "  hello world!  "
	fmt.Println(reverseWords(s))
}
