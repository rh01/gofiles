package main

func longestPalindrome(s string) int {
	m := make(map[byte]int)
	for _, v := range s {
		m[byte(v)]++
	}

	res := 0
	rem := false
	for _, v := range m {
		if v%2 == 0 {
			rem = true
		} else {
			res += (v >> 1) << 1
		}
	}
	if rem {
		return res + 1
	} else {
		return res
	}
}

func main() {

}
