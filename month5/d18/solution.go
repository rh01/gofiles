func validPalindrome(s string) bool {

	size := len(s)
	l, r := 0, size-1
	for l <= r {
		lStr := s[l]
		rStr := s[r]

		if lStr != rStr  {
            return isValid(s, l, r-1) || isValid(s, l+1, r)
		}

		l++
		r--
	}
	return true
}

func isValid(s string, l ,r int) bool {
	for l <= r {
        lStr := s[l]
        rStr := s[r]
		if lStr != rStr {
			return false
		}
		l++
		r--
	}
	return true
}

作者：rh01
链接：https://leetcode-cn.com/problems/valid-palindrome-ii/solution/go-by-rh01-13/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。