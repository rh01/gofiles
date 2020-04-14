package main

import "fmt"

func generateParenthesis(n int) []string {
	var res []string
	var gen func(int, int, string)

	gen = func(l, r int, curStr string) {
		if l > n || r > n || r > l {
			return
		}

		if l == n && r == n {
			res = append(res, curStr)
		}

		gen(l+1, r, curStr+"(")
		gen(l, r+1, curStr+")")
		return
	}

	gen(0, 0, "")
	return res
}

func main() {
	fmt.Println(generateParenthesis(3))
}
