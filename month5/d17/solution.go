func maxProduct(A []int) int {
	sz := len(A)
    B := reverse(A)
	for i := 1; i < sz; i++ {
		A[i] =computeA(A, i)
		B[i] =computeA(B, i)
	}
	return max(max(A...), max(B...))
}

func computeA(A []int, i int) int {
	if A[i-1] == 0 {
		return A[i]
	}
	return A[i] * A[i-1]
}

func max(A ...int) int {
	ret := A[0]
	for _, v := range A {
		if v > ret {
			ret = v
		}
	}
	return ret
}

func reverse(A []int) []int {
	B := make([]int, len(A))
	copy(B, A)
	l, r := 0, len(B) - 1
	for l < r {
		B[l], B[r] = B[r], B[l]
		l++
		r--
	}
	return B
}