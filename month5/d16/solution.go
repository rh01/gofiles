func findOrder(numCourses int, prerequisites [][]int) []int {
	// 拓扑排序
	sz := len(prerequisites)
	if sz == 0 {
		ret = []int{}
		for i := 0; i < numCourses; i++ {
			ret = append(ret, i)
		}
		return ret
	}

	// 初始化入度列表
	iNode := make([]int, numCourse)
	adjs := make([]map[int]bool, numCourse)
	for _, vertex := range prerequisites {
		// [0, 1] 若要修0，必须先修1
		to, from := vertex[0], vertex[1]
		iNode[to]++
		if _, exist := adjs[from][to]; !exist {
			adjs[from][to] = true
		}
	}

	// 先找到入度为0的节点，然后从入度节点开始
	queue := []int{}
	for i, node := range iNode {
		if node == 0 {
			queue = append(queue, i)
		}
	}

	res := []int{}
	cnt := 0
	for len(queue) != 0 {
		cur := queue[0]
		queue = queue[1:]
		res = append(res, cur)
		cnt++
		for k, _ := range adjs[cur] {
			iNode[k]--
			if iNode[k] == 0 {
				queue = append(queue, k)
			}
		}
	}
	if cnt == numCourse{
		return res
	}
	return []int{}
}