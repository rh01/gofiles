package main

import (
	"fmt"
)

// func main() {
//     n:=0
//     ans:=0

//     fmt.Scan(&n)
//     for i := 0; i < n; i++ {
//         for j := 0; j < n; j++ {
//             x:=0
//             fmt.Scan(&x)
//             ans = ans + x
//         }
//     }
//     fmt.Printf("%d\n",ans)
// }package main



func solve(length int, nums []int, sum int) [2]int {
	hmap := make(map[int]int, 0)
	for i := 0; i < length; i++ {
		cnum := nums[i]
		if _, exist := hmap[cnum]; !exist {
			hmap[cnum] = i
		}

		target := sum - cnum
		if val, exist := hmap[target]; exist && val != i {
			return [2]int{i, val}
		}
	}
	return [2]int{-1, -1}
}

func main() {
	n := 0
	//ans:=0
	fmt.Scan(&n)

	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
	}

	fmt.Println(nums)
}
