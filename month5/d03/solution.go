func maxSubArray(nums []int) int {
	var dp = make([]int, len(nums))
	dp[0] = nums[0]
	maxSub := dp[0]
	size := len(nums)

	for  i := 1; i < size; i++{
		dp[i] =  nums[i] +  max(dp[i-1] + nums[i], 0)
		if dp[i] > maxSub {
			maxSub = dp[i]
		}
	}
	return maxSub
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}