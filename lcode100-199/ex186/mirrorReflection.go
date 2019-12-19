package ex186

func mirrorReflection(p int, q int) int {
	// 求解p和q的最大公约数
	PQGcd := gcd(p, q)

	// 求解最小公倍数
	PQGbs := p * q / PQGcd

	// 如果最小公倍数除以q的结果能被2整除，那么说明一定落在2处
	if res := PQGbs / q; res%2 == 0 {
		return 2
	}

	if res := PQGbs / p; res%2 == 0 {
		return 0
	}

	return 1
}

// 求解最大公约数
func gcd(a, b int) int {
	if b == 0 {
		return a
	}

	return gcd(b, a%b)
}
