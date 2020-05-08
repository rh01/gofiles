package main

import "fmt"

func updateMatrix(matrix [][]int) [][]int {
	if matrix == nil || len(matrix) == 0 || len(matrix[0]) == 0 {
		return [][]int{}
	}

	m, n := len(matrix), len(matrix[0])
	var dp = make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	var walk func(x, y int) int
	walk = func(x, y int) int {
		if x < 0 || x >= m || y < 0 || y >= n {
			return 0
		}

		if matrix[x][y] == 0 {
			return 0
		}

		if dp[x][y] != 0 {
			return dp[x][y]
		}

		a := walk(x-1, y)
		b := walk(x+1, y)
		c := walk(x, y-1)
		d := walk(x, y+1)

		return min(min(min(a, b), c), d) + 1
	}

	// for i := 0; i < m; i++ {
	// 	for j := 0; j < n; j++ {
	// 		dp[i][j] = walk(i, j)
	// 	}
	// }
	walk(0, 0)
	return dp
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	a := [][]int{
		[]int{0, 0, 0},
		[]int{0, 1, 0},
		[]int{1, 1, 1},
	}
	fmt.Println(updateMatrix(a))
}
