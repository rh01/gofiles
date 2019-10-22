class Solution(object):
    def canFinish(self, numCourses, prerequisites):
        """
        :type numCourses: int
        :type prerequisites: List[List[int]]
        :rtype: bool
        """
        # 计算节点
        size = len(prerequisites)
        if size == 0:
            return True

        # 构造节点入度列表
        incoming = [0 for _ in range(numCourses)]
        # 构建邻接表，即出度节点
        adj_nodes = [set() for _ in range(numCourses)]

        # 遍历prerequisite
        for (out, incom) in prerequisites:
            incoming[out]+=1
            adj_nodes[incom].add(out)
        
        # 将入度为0的节点入队列
        queue = []
        for node in range(numCourses):
            if incoming[node] == 0: queue.append(node)
        
        # 出队，然后将对应的节点入度减1并清除邻接节点
        counter = 0
        while queue:
            cur = queue.pop(0)
            counter +=1
            
            for i in adj_nodes[cur]:
                incoming[i]-=1
                
                if incoming[i] == 0:
                    queue.append(i)
        return counter == numCourses