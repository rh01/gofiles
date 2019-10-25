#-*-coding:utf-8-*-
class Solution(object):
    def spiralOrder(self, matrix,):
        """
        :type matrix: List[List[int]]
        :rtype: List[int]
        """
        if matrix is None or len(matrix) == 0:
            return []
        res = []
        rows = len(matrix)
        cols = len(matrix[0])
        
        minvalue = min(rows, cols)
        for i in range((minvalue+1)//2 ):
            
            self.walk(matrix, res, i, i, rows-1-i, cols-1-i)
        return res
    
    def walk(self, matrix, res, r, c, rows, cols):
        if r == rows and c == cols:
            res.append(matrix[r][c])
            
        flag1 = flag2 = False
        if r == rows:
            flag1 = True
        
        if c == cols:
            flag2 = True
            
        for i in range(c, cols):
            res.append(matrix[r][i])
        
        for i in range(r, rows):
            res.append(matrix[i][cols])
        
        for i in range(cols, c, -1):
            res.append(matrix[rows][i])
            if flag1:
                break
            
        for i in range(rows, r, -1):
            res.append(matrix[i][c])
            if flag2:
                break
        
        return
         

        
sol = Solution()   
print(sol.spiralOrder(
[[1,2,3,4,5,6,7,8,9,10]




]))
"""
给定一个包含 m x n 个元素的矩阵（m 行, n 列），
请按照顺时针螺旋顺序，返回矩阵中的所有元素。

示例 1:

输入:
[
 [ 1, 2, 3 ],
 [ 4, 5, 6 ],
 [ 7, 8, 9 ]
]
输出: [1,2,3,6,9,8,7,4,5]
示例 2:

输入:
[
  [1, 2, 3, 4],
  [5, 6, 7, 8],
  [9,10,11,12]
]
输出: [1,2,3,4,8,12,11,10,9,5,6,7]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/spiral-matrix
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
"""