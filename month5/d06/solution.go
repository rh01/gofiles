func mincostTickets(days []int, costs []int) int {
	size := len(days)
	dp := make([]int, days[size-1] + 1)
	// -1 表示当天旅行， 0表示不旅行
	for idx := range days {
		dp[days[idx]] = -1
	}

	for i := 1; i <= days[size - 1]; i ++ {
		// 若不旅行
		if dp[i] == 0 {
			dp[i] = dp[i - 1]
		} else {
			// 有三种情况
			dp[i] = min(
				dp[max(0, i - 1)] + costs[0], 
				dp[max(0, i - 7)] + costs[1],
				dp[max(0, i - 30)] + costs[2],
			)
		}
	}
	return dp[len(dp) - 1]
}

func min(a ...int) int {
	ret := a[0]
	for _, v := range a {
		if v < ret {
			ret = v
		}
	}
	return ret
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}