package main

import (
	"fmt"
	"strings"
)

func isIsomorphic(s string, t string) bool {
	hmap1 := make(map[string]string)
	hmap2 := make(map[string]string)

	sArray := strings.Split(s, "")
	tArray := strings.Split(t, "")

	for i := 0; i < len(sArray); i++ {
		if _, exist := hmap1[sArray[i]]; !exist {
			hmap1[sArray[i]] = tArray[i]
		}

		if _, exist := hmap2[tArray[i]]; !exist {
			hmap2[tArray[i]] = sArray[i]
		}

		if hmap1[sArray[i]] != tArray[i] || hmap2[tArray[i]] != sArray[i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isIsomorphic("paper", "title"))
}
