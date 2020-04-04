func singleNumber(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	m := make(map[int]int, 0)
	for _, v := range nums {
		m[v]++
	}
	for i, v := range m {
		if v == 1 {
			return i
		}
	}
	return 0
}