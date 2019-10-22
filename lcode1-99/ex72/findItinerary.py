from collections import defaultdict
class Solution(object):
    def findItinerary(self, tickets):
        """
        :type tickets: List[List[str]]
        :rtype: List[str]
        """
        path = defaultdict(list)

        for _from, _to in tickets:
            path[_from].append(_to)
        
        for _from in path.keys():
            path[_from].sort()

        res = []

        def search(start):
            while path[start]:
                search(path[start].pop(0))
            res.append(start)
        
        search("JFK")
        return res
        


