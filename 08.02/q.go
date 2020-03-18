
var visted [][]bool 
var res [][]int 
func pathWithObstacles(obstacleGrid [][]int) [][]int {
	if obstacleGrid == nil || len(obstacleGrid) == 0 || len(obstacleGrid[0]) == 0 {
		return nil
	}

	// 这个题利用DFS来做
	m, n := len(obstacleGrid), len(obstacleGrid[0]) // 先确定行数和列数
	visted = make([][]bool, m)
	for i := 0; i < m; i++ {
		visted[i] = make([]bool, n)
	}
	res = make([][]int, 0)

	if walk(obstacleGrid, 0, 0, m, n) {
		return res
	}
	return nil
}


// walk 返回一条可行的即可
func walk(obstacleGrid [][]int, row, col int, rows, cols int) bool {
	// 剪枝
	if row >= rows || col >= cols || visted[row][col] || obstacleGrid[row][col] == 1  {
		return false
	}

	visted[row][col] = true
	res = append(res, []int{row, col})

	// 确定返回的条件
	if col == cols -1 && row == rows-1 && obstacleGrid[row][col] == 0 {
		return true
	}

	if walk(obstacleGrid, row+1, col, rows, cols) {
		return true
	}
    
    
	if walk(obstacleGrid, row, col+1, rows, cols) {
		return true
	}
	
	// visted[row][col] = false

	res = res[:len(res)-1] //在结果中移除
	return false
}

