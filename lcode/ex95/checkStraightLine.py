class Solution(object):
    def checkStraightLine(self, coordinates):
        """
        :type coordinates: List[List[int]]
        :rtype: bool
        """

        if len(coordinates) <= 2 or coordinates == None:
            return True
        startx, starty = coordinates[0]
        xs, ys = [],[]

        for item in coordinates[1:]:
            x, y = item
            xs.append(x - startx)
            ys.append(y - starty)
        
        if  any(xs)==False or  any(ys)==False:
            return True
        
        res = []
        for i in range(len(xs)):
            x = xs[i]
            y = ys[i]
            if y == 0:
                return False
            if x == 0:
                return False
            res.append(y*1.0/x)
        
        target = res[0]
        for item in range(1, len(res)):
            if res[item] != target:
                return False
        return True
             
            
             
sol = Solution()
print(sol.checkStraightLine(coordinates = [[1,1],[2,2],[3,4],[4,5],[5,6],[7,7]]))