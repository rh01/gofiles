func maxProfit(prices []int) int {
	mr, mc := -prices[0], 0
	for i := 1; i < len(prices); i++ {
		mr = min(mr, -prices[i])
		mc = max(mc, mr + prices[i])
	}
	return mc
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}