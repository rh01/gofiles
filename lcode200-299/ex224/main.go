package main

import "fmt"

func sortArray(nums []int) []int {
	var tmp = make([]int, 100000)
	for _, num := range nums {
		tmp[num+50000]++
	}
	res := make([]int, 0)
	var cnt int = 0
	for i := 0; i < 100000 && cnt < len(nums); i++ {
		j := tmp[i]
		for j > 0 {
			res = append(res, i-50000)
			cnt++
			j--
		}
	}
	return res
}

func main() {
	s := []int{5, 1, 1, 2, 0, 0}
	fmt.Println(sortArray(s))
}
