func subarraysDivByK(A []int, K int) int {
	m := make(map[int][]int, 0)
	s := 0
	cnt := 0
	for i, v := range A {
		s = s + v 
		
		want := s - K 
		if _, exist := m[want]; exist {
			cnt += len(v)
		}

		if _, exist := m[s]; !exist {
			m[s] = []int{i}
		} else {
			m[s] = append(m[s], i)
		}
	}
	return cnt
}