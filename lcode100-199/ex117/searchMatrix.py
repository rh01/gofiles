#coding:utf-8
class Solution(object):
    def searchMatrix(self, matrix, target):
        """
        :type matrix: List[List[int]]
        :type target: int
        :rtype: bool
        """
        if matrix is None or len(matrix) == 0:
            return False
        
        rows = len(matrix)
        cols = len(matrix[0])
        
        row = 0
        col = 0
        if matrix[0][0] == target:
            return True
        for i in range(1, rows):
            if matrix[i][0] == target:
                return True
            
            if matrix[i][0] > target:
                row = i -1
                break
            else:
                if i == rows-1:
                    row = rows - 1
                    break
        print(row)
        
        # for i in range(1, cols):
        #     if matrix[0][i] == target:
        #         return True
            
        #     if matrix[0][i] > target:
        #         col = i -1
        i = 0
        while i < cols:
            if matrix[row][i] == target:
                return True
            i += 1
        
        # i = 0
        # while i > rows:
        #     if matrix[i][col] == target:
        #         return True
        #     i += 1
        return False
                 
                
sol = Solution()      
matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,50]]
print(sol.searchMatrix(matrix, 30))


"""
编写一个高效的算法来判断 m x n 矩阵中，是否存在一个目标值。该矩阵具有如下特性：

每行中的整数从左到右按升序排列。
每行的第一个整数大于前一行的最后一个整数。
示例 1:

输入:
matrix = [
  [1,   3,  5,  7],
  [10, 11, 16, 20],
  [23, 30, 34, 50]
]
target = 3
输出: true
示例 2:

输入:
matrix = [
  [1,   3,  5,  7],
  [10, 11, 16, 20],
  [23, 30, 34, 50]
]
target = 13
输出: false

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/search-a-2d-matrix
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
"""