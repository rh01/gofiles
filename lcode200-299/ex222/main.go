package main

import "fmt"

func heapify(ar []int, i int, n int) {
	if i >= n {
		return
	}

	l, r, m := 2*i+1, 2*i+2, i
	if l < n && ar[l] < ar[m] {
		m = l
	}
	if r < n && ar[r] < ar[m] {
		m = r
	}
	if m != i {
		ar[m], ar[i] = ar[i], ar[m]
		heapify(ar, m, n)
	}
}

func buildHeap(ar []int, n int) {
	l := (n - 1) >> 1
	for ; l >= 0; l-- {
		heapify(ar, l, n)
	}
}

func sortArray(nums []int) []int {
	sz := len(nums)
	buildHeap(nums, sz)
	ret := make([]int, 0)
	for n := sz - 1; n >= 0; n-- {
		item := nums[0]
		nums[0], nums[n] = nums[n], nums[0]
		heapify(nums, 0, n)
		ret = append(ret, item)
	}
	return ret
}

func main() {
	ar := []int{3, 2, 1, 4, 5}
	fmt.Println(sortArray(ar))
}
