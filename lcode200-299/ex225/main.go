func maxDepthAfterSplit(seq string) []int {
	sbyte := []byte(seq)
	// cnt := 0
	sz := len(sbyte)
	ret := make([]int, sz)
	stack := []byte{}
	for i, b := range sbyte {
		if b == '(' {
			stack = append(stack, b)
			if len(stack) % 2 == 0 {
				ret[i] = 1
			} else {
				ret[i] = 0
			} 
		} else if b == ')' {
			if len(stack) % 2 == 0 {
				ret[i] = 1
			} else {
				ret[i] = 0
			} 
			stack = stack[:len(stack)-1]
		}
	}
	return ret
}
