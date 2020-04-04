package main

import "fmt"

func constructArr(a []int) []int {
	// 构建两个数组，一个是从左到右另一个从右到左
	if a == nil || len(a) == 0 {
		return []int{}
	}
	sz := len(a)
	l, r := make([]int, sz), make([]int, sz)
	l[0] = 1
	for i := 1; i < sz; i++ {
		l[i] = l[i-1] * a[i-1]
	}
	r[sz-1] = 1
	for i := sz - 2; i >= 0; i-- {
		r[i] = r[i+1] * a[i+1]
	}

	ret := make([]int, sz)
	for i := 0; i < sz; i++ {
		ret[i] = l[i] * r[i]
	}
	return ret
}

func main() {
	fmt.Println(constructArr([]int{1, 2, 3, 4, 5}))
}
