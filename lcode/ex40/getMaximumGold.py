class Solution(object):
    def getMaximumGold(self, grid):
        """
        :type grid: List[List[int]]
        :rtype: int
        """
        rows = len(grid)
        cols = len(grid[0])
        path = []
        maxnum = []
        for i in range(rows):
            for j in range(cols):
                if grid[i][j] != 0:
                    self.walk(grid, rows, cols, i, j, path, maxnum)
        return maxnum[0]

    def walk(self, grid, rows, cols, curx, cury, path, maxNum):
        if curx >= rows or curx < 0 or cury < 0 or cury >= cols or grid[curx][cury] == 0:
            return
        
        path.append(grid[curx][cury])
        temp = grid[curx][cury]
        grid[curx][cury] = 0

        self.walk(grid, rows, cols, curx+1, cury, res, maxNum)
        self.walk(grid, rows, cols, curx-1, cury, res, maxNum)
        self.walk(grid, rows, cols, curx, cury+1, res, maxNum)
        self.walk(grid, rows, cols, curx, cury-1, res, maxNum)
        if sum(path) > maxNum[0]:
            maxNum[0] = sum(path)
        grid[curx][cury]=temp
        return

