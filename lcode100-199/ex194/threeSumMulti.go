package main

import "fmt"

func threeSumMulti(A []int, target int) int {
	l, r := 0, len(A)-1
	count := 0

	for l < r {
		j := l + 1
		cursum := A[l] + A[j] + A[r]
		fmt.Println(cursum)
		if cursum < target {
			j++
		} else if cursum == target {
			fmt.Printf("A[%v]=%v, A[%v]=%v, A[%v]=%v\n", l, A[l], j, A[j], r, A[r])
			a1, a2, a3 := 0, 0, 0
			for jj := j; jj < r-1 && A[jj] == A[jj+1]; jj++ {
				a1++
			}

			for jj := l; jj < r-1 && A[jj] == A[jj+1]; jj++ {
				a2++
			}

			for jj := r; jj > l+1 && A[jj] == A[jj-1]; jj-- {
				a3++
			}
			count += a1 * a2 * a3

		} else {
			r--
			// l = 0
		}
	}

	// j = l + 1

	return count
}

func main() {
	fmt.Println(threeSumMulti([]int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5}, 8))
}
