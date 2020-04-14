func rotate(matrix [][]int)  {
	if matrix == nil || len(matrix) == 0 || len(matrix[0]) == 0 {
		return
	}

	sz := len(matrix)
	help := make([][]int, sz)
	for i := 0; i < sz; i++ {
		help[i] = make([]int, sz)
	}

	for  i := 0; i < sz; i++ {
		for  j := 0; j < sz; j++ {
			help[j][sz - i - 1] = matrix[i][j]
		}
	}
	for  i := 0; i < sz; i++ {
		for  j := 0; j < sz; j++ {
			matrix[i][j] = help[i][j];
		}
	}
	
}


