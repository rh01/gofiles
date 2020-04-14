func trap(height []int) int {
	sz := len(height)
	l, r := 0, sz - 1
	lMax, rMax := 0, 0
	ret := 0
	for l < r {
		if height[l] <= height[r] {
			if height[l] > lMax {
				lMax = height[l]
			} else {
				ret += lMax - height[l]
			}
			l++
		}else {
			if height[r] > rMax {
				rMax = height[r]
			} else {
				ret += rMax - height[r]
			}
			r--
		}
	}
	return ret
}