package main

import "fmt"

// channel close
func main() {

	intChan := make(chan int, 3)

	intChan <- 10
	intChan <- 100
	close(intChan)

	// 这时候就不能插入, panic
	// panic: send on closed channel
	//intChan <- 1000

	// 读取没有问题
	fmt.Println(<-intChan)
}
