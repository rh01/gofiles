func lengthOfLongestSubstring(s string) int {
	l, r, i, ret := 0, 0, 0, 0
	for ; r < len(s); r++ {
		for l= i ; l < r; l++ {
			if s[l] == s[r] {
				i = l + 1
				break
			}
		}

		if r - i + 1 > ret {
			ret = r - i + 1
		}
	}
	return ret
}

// second method
func lengthOfLongestSubstring(s string) int {

	// locations stand at first occurebce
	// 256 -> ascii code
	locations := [256]int{}
	left, maxLen := 0, 0
	
	// initilize locationn
	for i := range locations {
		locations[i] = -1
	}

	// iterator 
	for i := 0; i < len(s); i ++ {
		if locations[s[i]] >= left {
			left = locations[s[i]] +1
		} else if i - left + 1 > maxLen {
			maxLen = i - left + 1
		}

		locations[s[i]] = i
	}

	return maxLen
}

// method3 sliding window
// TODO: 2020-05-02