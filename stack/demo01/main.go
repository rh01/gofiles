// 单调栈
func nextGreaterElement(nums []int) []int {
	ans := make([]int, len(nums))
	s := []int{} // 模拟栈
	for i := len(nums) - 1; i >= 0; i-- { //倒着往栈里放
		for len(s) != 0 && s[len(s) - 1] <= nums[i] {
			s = s[:len(s) - 1]
		}
		if len(s) == 0 {
			ans[i] = -1
		} else {
			ans[i] = s[len(s)-1]
		}
		s = append(s, nums[i])
	}
	return s
}