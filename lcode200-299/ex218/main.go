package main

func maxDistance(grid [][]int) int {
	if grid == nil || len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	oceans, lands := [][2]int{}, [][2]int{}
	m, n := len(grid), len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				lands = append(lands, [2]int{i, j})
			}
		}
	}
	cnt := 0
	for len(lands) != 0 {
		tmp := make([][2]int, len(lands))
		copy(tmp, lands)
		lands = [][2]int{}
		for _, land := range tmp {
			x, y := land[0], land[1]
			if x+1 < m && grid[x+1][y] == 0 {
				grid[x+1][y] = 1
				lands = append(lands, [2]int{x + 1, y})
			}
			if x-1 >= 0 && grid[x-1][y] == 0 {
				grid[x-1][y] = 1
				lands = append(lands, [2]int{x - 1, y})
			}
			if y+1 < n && grid[x][y+1] == 0 {
				grid[x][y+1] = 1
				lands = append(lands, [2]int{x, y + 1})
			}
			if y-1 >= 0 && grid[x][y-1] == 0 {
				grid[x][y-1]
				lands = append(lands, [2]int{x, y - 1})
			}
		}
		if len(lands) != 0 {
			cnt++
		}
	}

	if cnt == 0 {
		return -1
	}

	return cnt
}

func main() {

}
