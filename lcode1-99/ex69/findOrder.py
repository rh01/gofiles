class Solution(object):
    def findOrder(self, numCourses, prerequisites):
        """
        :type numCourses: int
        :type prerequisites: List[List[int]]
        :rtype: List[int]
        """

        size = len(prerequisites)
        if size == 0: return []

        res = []
        queue = []

        ingrees = [0 for _ in range(numCourses)]
        adj = [set() for _ in range(numCourses)]

        for (first, second) in prerequisites:
            ingrees[first]+=1
            adj[second].add(first)

        for n in range(numCourses):
            if ingrees[n]==0: queue.append(n)
        count = 0
        while queue:
            cur = queue.pop(0)
            res.append(cur)
            count +=1
            for i in adj[cur]:
                ingrees[i]-=1
                if ingrees[i] == 0:
                    queue.append(i)
        if count == numCourses:
            return res
        else:
            return []
