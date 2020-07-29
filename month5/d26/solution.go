func findDuplicate(nums []int) int {
	if len(nums) == 0 {
		return  -1
	}

	slow, fast := 0,0
	for {
		// =看成快慢指针
		slow, fast = nums[slow], nums[nums[fast]]
		// 找到环
		if slow == fast {
			break
		}
	}
	fast = 0
	for slow != fast {
		slow, fast = nums[slow], nums[fast]
	}
	return slow
}