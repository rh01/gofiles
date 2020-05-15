func subarraySum(nums []int, k int) int {
	hmap := make(map[int][]int, 0)
	// subsum := []int{}
	s := 0
	cnt := 0
	for i, v := range nums {
		s = s + v
		if v, exist := hmap[s]; !exist {
			hmap[s] = []int{i}
		} else {
			hmap[s] = append(hmap[s], i)
		}
		// subsum = append(subsum, s)
		want := s - k
		if v, ok := hmap[want]; ok {
			cnt += len(v)
		}
	}
	return cnt
}