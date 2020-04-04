package main

import "fmt"

func nthUglyNumber(n int) int {
	if n == 1 {
		return 1
	}
	i, j, k := 0, 0, 0
	result := []int{1}
	res := 0
	for n > 1 {
		a := result[i] * 2
		b := result[j] * 3
		c := result[k] * 5
		fmt.Println("a, b, c", a, b, c)
		res = min(min(a, b), c)
		if res == a {
			i++
		}
		if res == b {
			j++
		}
		if res == c {
			k++
		}
		result = append(result, res)
		fmt.Println("res", res)
		n--
	}
	return result[len(result)-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(nthUglyNumber(10))
}

func isUgly(n int) bool {
	if n < 1 {
		return false
	}
	for n%2 == 0 {
		n /= 2
	}
	for n%3 == 0 {
		n /= 3
	}
	for n%5 == 0 {
		n /= 5
	}
	return n == 1
}
