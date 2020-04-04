func longestPalindrome(s string) int {
	size := len(s)
	res := 0
	for i := 0; i < size; i++ {
		a := extend(s, i, i)
		b := extend(s, i, i+1)
		res = max(res, max(a, b))
	}
	return res
}

func extend(s string, l, r int) int {
	for l >=0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}
	return r - l - 1
}

func max(i0, i1 int) int {
	if i0 > i1 {
		return i0
	}
	return i1
}