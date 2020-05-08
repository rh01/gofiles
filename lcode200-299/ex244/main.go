func canJump(nums []int) bool {
	l, n, i := len(nums) - 1, 1, 0
	
	for i = l - 1; i >= 0 ; i-- {
		if nums[i] >= n {
			n = 1
		} else {
			n++
		}
	}

	if i == 0 && n > 1 {
		return false
	}

	return true
}