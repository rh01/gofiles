class Solution(object):
    def merge(self, intervals):
        """
        :type intervals: List[List[int]]
        :rtype: List[List[int]]
        """
        if intervals is None or len(intervals) == 0 or len(intervals[0]) == 0:
            return []
        
        intervals.sort(key=lambda a: a[0])
        res = []
        curmin, curmax = intervals[0][0], intervals[0][1]
        size = len(intervals)
        for i in range(size-1):
            x1, x2 = intervals[i]
            y1, y2 = intervals[i+1]
            
            if curmax >= y1:
                curmax = max(curmax, y2)
            else:
                res.append([curmin, curmax])
                curmin = y1
                curmax = y2
        res.append([curmin, curmax])
        
        return res
        
        
s = Solution()
intervals =  [[1,4],[2,3]]

print(s.merge(intervals))            