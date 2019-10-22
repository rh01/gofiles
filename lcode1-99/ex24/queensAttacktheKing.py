class Solution(object):
    def queensAttacktheKing(self, queens, king):
        """
        :type queens: List[List[int]]
        :type king: List[int]
        :rtype: List[List[int]]
        """
        res = []

        # 八次循环
        
        x, y = king
        while True:
            move = [1, 0]
            move_x, move_y = move
            x += move_x
            y += move_y
            if not self.isValid((x, y)):
                break
            else:
                if [x, y] in queens:
                    res.append([x, y])
                    break
           


        x, y = king
        while True:
            move = [-1, 0]
            move_x, move_y = move
            x += move_x
            y += move_y
            if not self.isValid((x, y)):
                break
            else:
                if [x, y] in queens:
                    res.append([x, y])
                    break
            
        x, y = king
        while True:
            move = [0, 1]
            move_x, move_y = move
            x += move_x
            y += move_y
            if not self.isValid((x, y)):
                break
            else:
                if [x, y] in queens:
                    res.append([x, y])
                    break
                
        x, y = king   
        while True:
            move = [0, -1]
            move_x, move_y = move
            x +=move_x
            y +=move_y
            if not self.isValid((x, y)):
                break
            else:
                if [x, y] in queens:
                    res.append([x, y])
                    break
              
        x, y = king            
        while True:
            move = [1, 1]
            move_x, move_y = move
            x += move_x
            y += move_y
            if not self.isValid((x, y)):
                break
            else:
                if [x, y] in queens:
                    res.append([x, y])
                    break
                
            

        x, y = king
        while True:
            move = [-1, -1]
            move_x, move_y = move
            x += move_x
            y += move_y
            if not self.isValid((x, y)):
                break
            else:
                if [x, y] in queens:
                    res.append([x, y])
                    break
        
        x, y = king   
        while True:
            move = [-1, 1]
            move_x, move_y = move
            x += move_x
            y += move_y
            if not self.isValid((x, y)):
                break
            else:
                if [x, y] in queens:
                    res.append([x, y])
                    break
            

        x, y = king
        while True:
            move = [1, -1]
            move_x, move_y = move
            x += move_x
            y += move_y
            if not self.isValid((x, y)):
                break
            else:
                if [x, y] in queens:
                    res.append([x, y])
                    break
       
        return res

    def isValid(self, position):
        """
        :type position: List[int]
        :rtype bool
        """
        x, y = position
        if x < 0 or x >= 8 or y < 0 or y >= 8:
            return False
        else:
            return True
        
sol = Solution()
print(sol.queensAttacktheKing(queens = [[5,6],[7,7],[2,1],[0,7],[1,6],[5,1],[3,7],[0,3],[4,0],[1,2],[6,3],[5,0],[0,4],[2,2],[1,1],[6,4],[5,4],[0,0],[2,6],[4,5],[5,2],[1,4],[7,5],[2,3],[0,5],[4,2],[1,0],[2,7],[0,1],[4,6],[6,1],[0,6],[4,3],[1,7]], king = [3,4]))
print(sol.queensAttacktheKing(queens = [[0,0],[1,1],[2,2],[3,4],[3,5],[4,4],[4,5]], king = [3,3]))