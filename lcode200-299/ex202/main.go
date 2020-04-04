func maxSubArray(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	sz := len(nums)
	var dp = make([]int, sz)
	dp[0] = nums[0]
	ret := dp[0]
	for i := 1; i < sz; i++ {
		if dp[i-1] < 0 {
			dp[i] = nums[i]
		} else {
			dp[i] = nums[i] + dp[i-1]
		}
		if dp[i] > ret {
			ret =  dp[i]
		}
	}
	return ret
}
