package main

import "fmt"

func hasGroupsSizeX(deck []int) bool {
	if deck == nil || len(deck) < 2 {
		return false
	}
	m := make(map[int]int, 0)
	for _, v := range deck {
		m[v]++
	}
	min := 1<<32 - 1
	for _, v := range m {
		if v < min {
			min = v
		}
	}
	if min == 1 {
		return false
	}
	for _, v := range m {
		min =  gcd(v, min) 
		if min < 2 {
			return false
		}
	}
	return true
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func main() {
	fmt.Println(hasGroupsSizeX([]int{0, 0, 0, 1, 1, 1, 2, 2, 2}))
}
