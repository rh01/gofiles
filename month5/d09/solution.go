func mySqrt(x int) int {
	// 牛顿法
	var res = x
	for res * res > x {
		res = (res + x / res)/2
	}
	return res
}