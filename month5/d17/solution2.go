func maxProduct(nums []int) int {
	// dp algorithms...
	sz := len(nums)
	maxDP := make([]int, sz)
	minDP := make([]int, sz)

	maxDP[0] = minDP[0] = nums[0]
	ans := nums[0]
	for i := 1; i < sz; i++ {
		maxDP[i] = max(minDP[i-1] * nums[i], nums[i], maxDP[i-1] * nums[i])
		minDP[i] = max(minDP[i-1] * nums[i], nums[i], maxDP[i-1] * nums[i])
		ans = max(ans, maxDP[i])
	}
	return ans
}

func max(A ...int) int  {
	ret := A[0]
	for _, v := range A {
		if v > ret {
			ret = v
		}
	}
	return ret
}