#-*-coding:utf-8-*-
class Solution(object):
    def generateMatrix(self, n):
        """
        :type matrix: List[List[int]]
        :rtype: List[int]
        """
        # if n is None or len(matrix) == 0:
        #     return []
        matrix = [[0 for _ in range(n)] for i in range(n)]
        i, j, di, dj = 0, 0, 0, 1
        count = 1
        if matrix != []:
            for _ in range(n*n):
                #r.append(matrix[i][j])
                matrix[i][j] = count
                if matrix[(i + di) % n][(j + dj) % n] != 0:
                    di, dj = dj, -di
                i += di
                j += dj
                count +=1
        return matrix
    
    
         

        
sol = Solution()   
print(sol.generateMatrix(3))
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
        

"""
给定一个正整数 n，生成一个包含 1 到 n2 所有元素，且元素按顺时针顺序螺旋排列的正方形矩阵。

示例:

输入: 3
输出:
[
 [ 1, 2, 3 ],
 [ 8, 9, 4 ],
 [ 7, 6, 5 ]
]
"""