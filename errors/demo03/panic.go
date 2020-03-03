package main

import "fmt"

func main() {
	var array = []int{1, 2, 3, 4}
	fmt.Printf("array[%d] = %d", 4, array[4])
}
// panic: runtime error: index out of range [4] with length 4
