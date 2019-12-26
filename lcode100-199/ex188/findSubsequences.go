package main

import "fmt"

func findSubsequences(nums []int) [][]int {
	// [4, 6, 7, 7]

	if nums == nil || len(nums) <= 1 {
		return [][]int{}
	}

	size := len(nums)
	res := make([][]int, 0)

	walk(nums, size, []int{}, 0, &res)

	return res
}

func walk(nums []int, size int, cur []int, start int, res *[][]int) {
	// 检查递归的退出条件
	if len(cur) >= 2 {
		*res = append(*res, cur)
	}

	for i := start; i < size; i++ {

		if i != start && nums[i] == nums[i-1] {
			continue
		}

		if (cur != nil && len(cur) != 0) && nums[i] <= cur[len(cur)-1] {
			continue
		}

		// cur = append(cur, nums[i])
		walk(nums, size, append(cur, nums[i]), i+1, res)
		// cur = cur[:len(cur)-1]
	}

	return
}

func main() {
	fmt.Println(findSubsequences([]int{4, 6, 7, 7}))
}
