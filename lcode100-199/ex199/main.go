package main

func surfaceArea(grid [][]int) int {
	n, area := len(grid), 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			level := grid[i][j]
			if level > 0 {
				area += (level << 2) + 2
				if i > 0 {
					area -= min(level, grid[i-1][j]) << 1
				}
				if j > 0 {
					area -= min(level, grid[i][j-1]) << 1
				}
			}
		}
	}
	return area
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}


func main() {

}
