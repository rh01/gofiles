package main

import "fmt"

func majorityElement(nums []int) int {
	cnt := 0
	res := 0
	//res := 0
	for _, v := range nums {
		if cnt == 0 {
			res = v
		}
		if v == res {
			cnt++
		} else {
			cnt--
		}

	}
	return res
}

func main() {
	fmt.Println(majorityElement([]int{1, 2, 3, 3, 3, 3, 5, 4, 2}))
}
