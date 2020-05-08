package main

func maxArea(height []int) int {
	sz := len(height)
	l, r := 0, sz - 1
	res := 0
	// m := 0
	for l < r {
		h := min(height[l], height[r])
		volume := h * (r - l)
		if volume > res {
			res = volume
		}

		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return res
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func main() {
	
}