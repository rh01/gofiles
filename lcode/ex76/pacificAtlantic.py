"""
给定一个 m x n 的非负整数矩阵来表示一片大陆上各个单元格的高度。“太平洋”处于大陆的左边界和上边界，而“大西洋”处于大陆的右边界和下边界。

规定水流只能按照上、下、左、右四个方向流动，且只能从高到低或者在同等高度上流动。

请找出那些水流既可以流动到“太平洋”，又能流动到“大西洋”的陆地单元的坐标。

 

提示：

输出坐标的顺序不重要
m 和 n 都小于150
 

示例：

 

给定下面的 5x5 矩阵:

  太平洋 ~   ~   ~   ~   ~ 
       ~  1   2   2   3  (5) *
       ~  3   2   3  (4) (4) *
       ~  2   4  (5)  3   1  *
       ~ (6) (7)  1   4   5  *
       ~ (5)  1   1   2   4  *
          *   *   *   *   * 大西洋

返回:

[[0, 4], [1, 3], [1, 4], [2, 2], [3, 0], [3, 1], [4, 0]] (上图中带括号的单元).

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/pacific-atlantic-water-flow
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
"""
class Solution(object):
    def pacificAtlantic(self, matrix):
        """
        :type matrix: List[List[int]]
        :rtype: List[List[int]]
        """
        if matrix is None or len(matrix) == 0:
            return []
        
        res = []
        rows = len(matrix)
        cols = len(matrix[0])

        # pacificAtlantic
        # 太平洋
        pacific = [[False for _ in range(cols)] for _ in range(rows)]
        atlantic = [[False for _ in range(cols)] for _ in range(rows)]
       
        # 从四个边进行遍历
        for i in range(rows):
            self.walk(matrix, rows, cols, i, 0, pacific, matrix[i][0]) # 第一列
            self.walk(matrix, rows, cols, i, cols-1,atlantic, matrix[i][cols-1])

        for i in range(cols):
            self.walk(matrix, rows, cols, 0, i, pacific,  matrix[0][i]) # 第一行
            self.walk(matrix, rows, cols, rows-1, i, atlantic, matrix[rows-1][i])

        res = []
        for i in range(rows):
            for j in range(cols):
                if pacific[i][j] and atlantic[i][j]:
                    res.append([i, j])
        return res

    def walk(self, matrix, rows, cols, row, col, visited, target):
        # 终止条件
        if row < 0 or row >= rows or col <0 or col>=cols or visited[row][col] or matrix[row][col] < target:
            return

        # 正常的逻辑
        visited[row][col] = True
        self.walk(matrix, rows, cols, row+1, col, visited, matrix[row][col])
        self.walk(matrix, rows, cols, row-1, col, visited, matrix[row][col])
        self.walk(matrix, rows, cols, row, col+1, visited, matrix[row][col])
        self.walk(matrix, rows, cols, row, col-1, visited, matrix[row][col])
        return 
        
