
func sumNums(n int) int {
	_ = n > 0 && func() bool { n += sumNums(n - 1); return true }()
	return n
}
