package main

import "sort"

func main() {

}

func intersection(start1 []int, end1 []int, start2 []int, end2 []int) []float64 {
	x1, y1 := start1[0], start1[1]
	x2, y2 := end1[0], end1[1]

	x3, y3 := start2[0], start2[1]
	x4, y4 := end2[0], end2[1]

	var det func(int, int, int, int) int
	det = func(a, b, c, d int) int {
		return a*d - b*c
	}

	d := det(x1-x2, x4-x3, y1-y2, y4-y3)
	p := det(x4-x2, x4-x3, y4-y2, y4-y3)
	q := det(x1-x2, x4-x2, y1-y2, y4-y2)

	if d != 0 {
		lam, eta := p/d, q/d
		if !(0 <= lam && lam <= 1 && 0 <= eta && eta <= 1) {
			return []float64{}
		}
		return []float64{float64(lam*x1 + (1-lam)*x2), float64(lam*y1 + (1-lam)*y2)}
	}
	if p != 0 || q != 0 {
		return []float64{}
	}

	t1 := append(append([]int{}, start1...), end1...)
	t2 := append(append([]int{}, start2...), end2...)
	sort.Ints(t1)
	sort.Ints(t2)

	if t1[1] < t2[0] || t2[1] < t1[0] {
		return []float64{}
	}

	return max(t1[0], t2[0])
}

func max(a, b int) float64 {
	if a > b {
		return float64(a)
	}
	return float64(b)
}
