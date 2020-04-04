
func gameOfLife(board [][]int) {
	if board == nil || len(board) == 0 || len(board[0]) == 0 {
		return
	}
	m, n := len(board), len(board[0])
	tmpBoard := make([][]int, m)
	for i := 0; i < m; i++ {
		tmpBoard[i] = make([]int, n)
	}
	for x := 0; x < m; x++ {
		for y := 0; y < n; y++ {
			s := getNumOfAliveCell(board, x-1, y) +
				getNumOfAliveCell(board, x+1, y) +
				getNumOfAliveCell(board, x, y-1) +
				getNumOfAliveCell(board, x, y+1) +
				getNumOfAliveCell(board, x-1, y-1) +
				getNumOfAliveCell(board, x-1, y+1) +
				getNumOfAliveCell(board, x+1, y-1) +
				getNumOfAliveCell(board, x+1, y+1)
			if s < 2 {
				tmpBoard[x][y] = 0
			} else if (s == 2 || s == 3) && board[x][y] == 1 {
				tmpBoard[x][y] = 1
			} else if s > 3 {
				tmpBoard[x][y] = 0
			} else if s == 3 && board[x][y] == 0 {
				tmpBoard[x][y] = 1
			}
		}
	}
	copy(board, tmpBoard)
}

func getNumOfAliveCell(board [][]int, x, y int) int {
	m, n := len(board), len(board[0])
	if x < 0 || y < 0 || x >= m || y >= n {
		return 0
	}
	return board[x][y]
}
