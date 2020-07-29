
func kidsWithCandies(candies []int, extraCandies int) []bool {
	// 取出最大的candy
	if candies == nil || len(candies) == 0 {
		return []bool{}
	}
	res := make([]bool, len(candies))
	maxCandy := candies[0]
	for _, candy := range candies {
		if candy > maxCandy {
			maxCandy = candy
		}
	}

	for i, candy := range candies {
		if candy+extraCandies >= maxCandy {
			res[i] = true
		}
	}
	return res
}

