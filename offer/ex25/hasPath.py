# -*- coding:utf-8 -*-


# 请设计一个函数，用来判断在一个矩阵中是否存在一条包含某字符串所有字符的路径。
# 路径可以从矩阵中的任意一个格子开始，
# 每一步可以在矩阵中向左，向右，向上，向下移动一个格子。
# 如果一条路径经过了矩阵中的某一个格子，则该路径不能再进入该格子。 
# 
# 例如 
#      a b c e 
#      s f c s 
#      a d e e 
# 矩阵中包含一条字符串"bcced"的路径，
# 但是矩阵中不包含"abcb"路径，
# 因为字符串的第一个字符b占据了矩阵中的第一行第二个格子之后，
# 路径不能再次进入该格子。

# -*- coding:utf-8 -*-
class Solution:
    
    def hasPath(self, matrix, rows, cols, path):
        # write code here
        flag = [False for _ in range(len(matrix))]
        for i in range(rows):
            for j in range(cols):
                # 遍历
                if self.judge(matrix, rows, cols, flag, i, j, path, 0):
                    return True
        return False

    def judge(self, matrix, rows, cols, flag, i, j, path, level):
        index = i * cols + j
        # 判断越界条件,递归终止条件
        if i >= rows or j >= cols or i < 0 or j < 0 or \
                matrix[index] != path[level] or \
                flag[index] == True:
            return False

        if level == len(path) - 1:
            return True

        flag[index] = True

        if (self.judge(matrix, rows, cols, flag, i + 1, j, path, level + 1) or \
                self.judge(matrix, rows, cols, flag, i, j + 1, path, level + 1) or \
                self.judge(matrix, rows, cols, flag, i - 1, j, path, level + 1) or \
                self.judge(matrix, rows, cols, flag, i, j - 1, path, level + 1)):
            return True

        flag[index] = False
        return False
