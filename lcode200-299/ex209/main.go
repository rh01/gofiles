func lengthOfLongestSubstring(s string) int {
	// q store index of first
	q := []int{}
	m := make(map[byte]int)
	maxLength := 0
	for i := 0; i < len(s); i++ {
		if v, ok := m[s[i]]; ok && v > 0 {
			for len(q) != 0 && s[q[0]] != s[i] {
				m[s[q[0]]]--
				q = q[1:]
			}
			m[s[q[0]]]--
			q = q[1:]
		} 
        m[s[i]]++
		q = append(q, i)
		if len(q) > maxLength {
			maxLength = len(q)
		}
	}
	return maxLength
}