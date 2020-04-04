func numRookCaptures(board [][]byte) int {
    if board == nil || len(board) == 0 || len(board[0]) == 0 {
		return 0
	}
	// col, row := make([]int, 8), make([]int, 8)
	ri, rj := -1, -1
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if board[i][j] == 'R' {
				ri = i 
				rj = j
			}
		}
	}
	cnt, i := 0, ri + 1
	for i < 8 && board[i][rj] != 'B' {
		if board[i][rj] == 'p' {
			cnt++
			break
		}
		i++
	}
	i = ri - 1
	for i >= 0 && board[i][rj] != 'B' {
		if board[i][rj] == 'p' {
			cnt++
			break
		}
		i--
	}
	j := rj + 1
	for j < 8 && board[ri][j] != 'B' {
		if board[ri][j] == 'p' {
			cnt++
			break
		}
		j++
	}
	j := rj - 1
	for j >= 0 && board[ri][j] != 'B' {
		if board[ri][j] == 'p' {
			cnt++
			break
		}
		j--
	}
	return cnt
}