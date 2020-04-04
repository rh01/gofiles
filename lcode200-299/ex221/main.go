package main

import "fmt"

func sortArray(nums []int) []int {
	sz := len(nums)
	for i := 1; i < sz; i ++ {
		tmp := nums[i]
		var j int = i - 1
		for ; j >= 0 && nums[j] >= tmp ; j-- {
			nums[j+1] = nums[j]
		}
		nums[j+1] = tmp
	} 
	return nums
}

func main() {
	sArr := []int{2,3,1,4,-1,0,-2,-5,7}
	fmt.Println(sortArray(sArr))
}