func jump(nums []int) int {
	start := 0
	size := len(nums)
	count := 0
	for start < size - 1 {
		step := nums[start] 
		if start + step > size {
			return count + 1
		}

		maxSkip := 0
		temp := start
		for i := cur + 1; i <= cur + steps; i++ {
			pos := nums[i] + i // 表示你能够达到的长度
			if pos > maxSkip {
				maxSkip = pos
				temp = i
			}
		}  
		start = temp
		count++
	}
	return count
}