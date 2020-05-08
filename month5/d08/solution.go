func maximalSquare(matrix [][]byte) int {
    if matrix == nil  || len(matrix) == 0 || len(matrix[0]) == 0{
        return 0
    }
	rows, cols := len(matrix), len(matrix[0])
	
	// 初始化
	dp := make([][]int, rows)
	for i := 0; i < rows; i++ {
		dp[i] = make([]int, cols)
	}
    
    
	// 记录最大值
	ret := 0

	for i := 0; i < rows; i++ {
		dp[i][0] = int(matrix[i][0] - '0')
        if dp[i][0] > ret {
			ret =  dp[i][0] 
        }
	}

	for i := 0; i < cols; i++ {
		dp[0][i] = int(matrix[0][i] - '0')
        if dp[0][i] > ret {
            ret = dp[0][i]
        }
	}

	for i := 1; i < rows; i++ {
		for j := 1; j < cols; j++ {
			if matrix[i][j] == '1' {
				dp[i][j] = 1 + min(dp[i-1][j-1], dp[i-1][j], dp[i][j-1])
				if dp[i][j] > ret {
					ret = dp[i][j]
				}
			}
		}
	}

	return ret*ret
}

func min(a ...int) int {
	ret := a[0]
	for _, v := range a {
		if v < ret {
			ret = v
		} 
	}
	return ret
}
