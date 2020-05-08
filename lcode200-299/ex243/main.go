package main

import "sort"

type SortBy [][]int

func (a SortBy) Len() int           { return len(a) }
func (a SortBy) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortBy) Less(i, j int) bool { return a[i][0] < a[j][0] }

func merge(intervals [][]int) [][]int {
	is := SortBy(intervals)
	sort.Sort(is)
	res := [][]int{}
	curmin, curmax := is[0][0], is[0][1]
	size := len(is)
	for i := 1; i < size; i++ {
		y1, y2 := is[i][0], is[i][1]

		if curmax >= y1 {
			curmax = max(curmax, y2)
		} else {
			res = append(res, []int{curmin, curmax})
			curmin = y1
			curmax = y2
		}
	}
	res = append(res, []int{curmin, curmax})
	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func main() {

}
