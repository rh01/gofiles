func singleNumbers(nums []int) []int {
	sum := 0
	// 异或和结果
	for _, v := range nums {
		sum = sum ^ v
	}
	// 得到sum的二进制的1的最低位
	flag := (-sum)&sum
	result := [2]int{0, 0}
	for _, num := range nums {
		if flag & num == 0 {
			result[0] ^=  num
		} else {
			result[1] ^= num
		}
	}
	return result[:]
}