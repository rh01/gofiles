package main

func missingNumber(nums []int) int {
	// missing := -1
	l, r := 0, len(nums)
	for l < r {
		mid := l + (r-l)>>1
		if mid <= nums[mid] {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l
}
