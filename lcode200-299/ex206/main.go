func firstUniqChar(s string) byte {
	if s == "" {
		return ' '
	}
	m := make(map[byte]int, 0)
	for _, c := range s {
		m[c]++
	}
	for _, c := range s {
		if m[c] == 1 {
			return c
		}
	}
	return ' '
}